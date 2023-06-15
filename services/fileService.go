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
)

type Transaction struct {
	Type    string
	Date    string
	Product string
	Value   string
	Seller  string
}

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
	//fileExt := filepath.Ext(filename)

	hash := md5.New()
	if _, err := io.Copy(hash, bytes.NewReader([]byte(data))); err != nil {
		return "", err
	}

	hashInBytes := hash.Sum(nil)[:16]
	fileHash := hex.EncodeToString(hashInBytes)

	return fileHash + ".txt", nil
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
