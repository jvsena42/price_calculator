package filemanager

import (
	"bufio"
	"encoding/json"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, error := os.Open(path)

	if error != nil {
		return nil, error
	}

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err := scanner.Err()

	if err != nil {
		file.Close()
		return nil, err
	}

	file.Close()
	return lines, nil
}

func WriteJson(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	errorEncoder := encoder.Encode(data)

	if err != nil {
		file.Close()
		return errorEncoder
	}

	file.Close()
	return nil
}
