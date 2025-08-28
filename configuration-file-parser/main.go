package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseConfig() (map[string]string, error) {
	config := make(map[string]string, 0)

	file, err := os.Open("config.txt")
	if err != nil {
		return config, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "=")
		if len(tokens) < 2 {
			continue
		}
		config[tokens[0]] = tokens[1]

		if err = scanner.Err(); err != nil {
			return config, err
		}
	}

	return config, nil
}

func main() {

	config, err := parseConfig()
	if err != nil {
		fmt.Println("Failed due to: ", err)
	} else {
		fmt.Println(config)
	}

}
