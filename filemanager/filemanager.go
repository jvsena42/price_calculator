package filemanager

import (
	"bufio"
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
