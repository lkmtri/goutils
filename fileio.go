package goutils

import (
	"bufio"
	"log"
	"os"
)

func ReadLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	res := []string{}

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}
