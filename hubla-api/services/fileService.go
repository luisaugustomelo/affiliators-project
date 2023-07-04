package services

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/luisaugustomelo/hubla-challenge/database/models"
	"github.com/luisaugustomelo/hubla-challenge/interfaces"
	"gorm.io/gorm"
)

var productToProducer map[string]string

func ProcessFile(encodedString string) (string, error) {
	hashedFilename, err := getHashedFilename(encodedString)
	if err != nil {
		return "", err
	}

	filePath := "hubla-api/public/single/" + hashedFilename
	err = saveFile([]byte(encodedString), filePath)
	if err != nil {
		return "", err
	}

	return hashedFilename, nil
}

func saveFile(data []byte, filepath string) error {
	return ioutil.WriteFile(filepath, data, 0644)
}

func getHashedFilename(data string) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, bytes.NewReader([]byte(data))); err != nil {
		return "", err
	}

	hashInBytes := hash.Sum(nil)[:16]
	fileHash := hex.EncodeToString(hashInBytes)

	return fileHash + ".txt", nil
}

func mapProductToProducer(salesFiles []models.SalesFile) map[string]string {
	productToProducer := make(map[string]string)

	for _, sale := range salesFiles {
		if sale.SalesType == 1 {
			productToProducer[sale.Product] = sale.Seller
		}
	}

	return productToProducer
}

// função para calcular o saldo para vendedores
func calculateBalances(sales []models.SalesFile) map[string]float64 {
	productToProducer = mapProductToProducer(sales)
	var balances = make(map[string]float64)

	for _, t := range sales {
		// checar o tipo de transação
		switch t.SalesType {
		case 1, 4:
			// paga para o produto, o vendedor é o produtor
			balances[t.Seller] += t.Value
		case 2:
			// debita da conta do produtor
			producer := productToProducer[t.Product]
			balances[producer] += t.Value
		case 3:
			balances[t.Seller] -= t.Value

		}
	}
	return balances
}

func ReadSales(filename string) ([]models.SalesFile, []models.Balance, error) {
	filePath := "hubla-api/public/single/" + filename
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var sales []models.SalesFile
	for scanner.Scan() {
		line := scanner.Text()
		salesType, err := strconv.Atoi(line[0:1])
		if err != nil {
			return nil, nil, err
		}

		cents, err := strconv.Atoi(line[56:66])
		if err != nil {
			return nil, nil, err
		}

		value := float64(cents) / 100.0
		t := models.SalesFile{
			SalesType: uint(salesType),
			Date:      strings.TrimSpace(line[1:26]),
			Product:   strings.TrimSpace(line[26:56]),
			Value:     value,
			Seller:    strings.TrimSpace(line[66:]),
		}
		sales = append(sales, t)
	}

	balances := calculateBalances(sales)

	var finalBalances []models.Balance
	for index, balance := range balances {
		isProducer := false
		for _, producer := range productToProducer {
			if producer == index {
				isProducer = true
				break
			}
		}

		userBalance := models.Balance{
			Balance: balance,
			Role:    1,
			Name:    index,
			Hash:    strings.TrimSuffix(filename, ".txt"),
		}

		if isProducer {
			finalBalances = append(finalBalances, userBalance)
			continue
		}

		userBalance.Role = 2
		finalBalances = append(finalBalances, userBalance)
	}
	return sales, finalBalances, scanner.Err()
}

func CheckFileStatus(db interfaces.Datastore, id int) ([]models.Balance, error) {
	mq := &models.QueueProcessing{}
	if err := db.Where("id = ?", id).First(mq).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("process not found")
		}
		return nil, fmt.Errorf("failed to get process status")
	}

	if mq.Status == "success" {
		var balances []models.Balance
		if err := db.Where("hash = ?", strings.TrimSuffix(mq.Hash, ".txt")).Find(&balances).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("balances not found")
			}
			return nil, fmt.Errorf("failed to get balances")
		}
		return balances, nil
	}

	if mq.Status == "error" {
		return nil, fmt.Errorf("error to proccess file")
	}

	return nil, nil
}

func FileProducts(db interfaces.Datastore, id int) ([]models.SalesFile, error) {
	mq := &models.QueueProcessing{}
	if err := db.Where("id = ?", id).First(mq).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("process not found")
		}
		return nil, fmt.Errorf("failed to get process status")
	}

	var result []struct {
		SalesFile models.SalesFile
		SalesType models.SaleType
	}

	if mq.Status == "success" {
		var sales []models.SalesFile
		if err := db.Table("sales_files").Select("product, seller, value, description").Joins("sales_files LEFT JOIN sale_types ON sales_type = sale_types.id").
			Where("hash = ?", strings.TrimSuffix(mq.Hash, ".txt")).
			Find(&result).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("balances not found")
			}
			return nil, fmt.Errorf("failed to get balances")
		}
		return sales, nil
	}

	if mq.Status == "error" {
		return nil, fmt.Errorf("error to proccess file")
	}

	return nil, nil
}
