package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	s := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	vals := make([]int, 0)

	for _, str := range s {
		var firstDigit, lastDigit string

		for _, char := range str {
			if _, err := strconv.Atoi(string(char)); err == nil {
				firstDigit = string(char)
				break
			}
		}

		for i := len(str) - 1; i >= 0; i-- {
			if _, err := strconv.Atoi(string(str[i])); err == nil {
				lastDigit = string(str[i])
				break
			}
		}

		num := firstDigit + lastDigit
		digit, err := strconv.Atoi(string(num)); 
		if err == nil {
			vals = append(vals, digit)
		}
	}

	sum := 0
	for _, val := range vals {
		sum += val
	}
	fmt.Println(sum)
}
