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

	words := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	nums := make([]int, 0)

	for _, word := range words {
		var firstDigit, lastDigit string

		for _, char := range word {
			if _, err := strconv.Atoi(string(char)); err == nil {
				firstDigit = string(char)
				break
			}
		}

		for i := len(word) - 1; i >= 0; i-- {
			if _, err := strconv.Atoi(string(word[i])); err == nil {
				lastDigit = string(word[i])
				break
			}
		}

		num := firstDigit + lastDigit
		digit, err := strconv.Atoi(string(num)); 
		if err == nil {
			nums = append(nums, digit)
		}
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}
