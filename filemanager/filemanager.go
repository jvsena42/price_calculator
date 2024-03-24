package filemanager

import (
	"bufio"
	"encoding/json"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputPath string, outPutPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outPutPath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, error := os.Open(fm.InputFilePath)

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

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

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
