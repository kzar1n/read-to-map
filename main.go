package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFileAndCreateMap(filename string) ([]map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var headers []string
	var records []map[string]string

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ";")

		if headers == nil {
			headers = fields
		} else {
			record := make(map[string]string)
			for i, header := range headers {
				if i < len(fields) {
					record[header] = fields[i]
				}
			}
			records = append(records, record)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func main() {
	records, err := readFileAndCreateMap("file.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, record := range records {
		fmt.Println(record)
		fmt.Println("Name:", record["coluna3"])
	}
}
