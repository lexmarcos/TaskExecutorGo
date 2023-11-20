package utils

import (
	"os"
)

func WriteToFile(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content + "\n")
	return err
}

func ReadFromFile(filePath string) (string, error) {

	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
