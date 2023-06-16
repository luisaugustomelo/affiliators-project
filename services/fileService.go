package services

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/luisaugustomelo/hubla-challenge/database/models"
)

func ProcessFile(encodedString string) (string, error) {
	hashedFilename, err := getHashedFilename(encodedString)
	if err != nil {
		return "", err
	}

	filePath := "public/single/" + hashedFilename
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

func ReadSales(filename string) ([]models.Sale, error) {
	filePath := "public/single/" + filename
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var sales []models.Sale
	for scanner.Scan() {
		line := scanner.Text()
		t := models.Sale{
			Type:    strings.TrimSpace(line[0:1]),
			Date:    strings.TrimSpace(line[1:26]),
			Product: strings.TrimSpace(line[26:56]),
			Value:   strings.TrimSpace(line[56:66]),
			Seller:  strings.TrimSpace(line[66:]),
		}
		sales = append(sales, t)
	}

	return sales, scanner.Err()
}
