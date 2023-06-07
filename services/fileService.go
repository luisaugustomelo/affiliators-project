package services

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Transaction struct {
	Type    string
	Date    string
	Product string
	Value   string
	Seller  string
}

func UploadSingleFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("file err : %s", err.Error()))
	}

	fileExt := filepath.Ext(file.Filename)
	fileReader, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	defer fileReader.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, fileReader); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	hashInBytes := hash.Sum(nil)[:16]
	fileHash := hex.EncodeToString(hashInBytes)
	filename := fileHash + fileExt
	filePath := "public/single/" + filename

	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		out, err := os.Create(filePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		defer out.Close()

		fileReader, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		if _, err = io.Copy(out, fileReader); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
	}

	file1, err := os.Open(filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	defer file1.Close()

	scanner := bufio.NewScanner(file1)
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

	if err := scanner.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"filepath": "/images/single/" + filename})
}
