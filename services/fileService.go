package services

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type Transaction struct {
	Type    string
	Date    string
	Product string
	Value   string
	Seller  string
}

func ProcessFile(file *multipart.FileHeader) (string, error) {
	hashedFilename, err := getHashedFilename(file)
	if err != nil {
		return "", err
	}

	filePath := "public/single/" + hashedFilename
	err = saveFile(file, filePath)
	if err != nil {
		return "", err
	}

	return hashedFilename, nil
}

func getHashedFilename(file *multipart.FileHeader) (string, error) {
	fileExt := filepath.Ext(file.Filename)
	fileReader, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileReader.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, fileReader); err != nil {
		return "", err
	}

	hashInBytes := hash.Sum(nil)[:16]
	fileHash := hex.EncodeToString(hashInBytes)

	return fileHash + fileExt, nil
}

func saveFile(file *multipart.FileHeader, filePath string) error {
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return nil
	}

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	fileReader, err := file.Open()
	if err != nil {
		return err
	}

	_, err = io.Copy(out, fileReader)
	return err
}

func ReadTransactions(filename string) ([]Transaction, error) {
	filePath := "public/single/" + filename
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var transactions []Transaction
	for scanner.Scan() {
		line := scanner.Text()
		t := Transaction{
			Type:    strings.TrimSpace(line[0:1]),
			Date:    strings.TrimSpace(line[1:26]),
			Product: strings.TrimSpace(line[26:56]),
			Value:   strings.TrimSpace(line[56:66]),
			Seller:  strings.TrimSpace(line[66:]),
		}
		transactions = append(transactions, t)
	}

	return transactions, scanner.Err()
}
