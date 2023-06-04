package services

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

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
	originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	filePath := "http://localhost:3030/images/single/" + filename

	out, err := os.Create("public/single/" + filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	defer out.Close()

	f, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	defer f.Close()

	if _, err = io.Copy(out, f); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	file1, err := os.Open("public/single/" + filename)
	if err != nil {
		return err
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
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"filepath": filePath})
}
