package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFirstLastDigit(words []string) []int {
	nums := make([]int, 0)
	
	for _, word := range words {
		var firstDigit, lastDigit string
		i, j := 0, len(word)-1
	
		for i < len(word) {
			if firstDigit == "" {
				if _, err := strconv.Atoi(string(word[i])); err == nil {
					firstDigit = string(word[i])
				}
			}
			if lastDigit == "" {
				if _, err := strconv.Atoi(string(word[j])); err == nil {
					lastDigit = string(word[j])
				}
			}
			if firstDigit != "" && lastDigit != "" {
				break
			}
			i++
			j--
		}
	
		num := firstDigit + lastDigit
		digit, err := strconv.Atoi(num)
		if err == nil {
			nums = append(nums, digit)
		}
	}

	return nums
}

func getWordsSlice(doc string) []string {
	words := make([]string, 0)

	file, err := os.Open(doc)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	return words
}

func main() {
	words := getWordsSlice("input.txt")

	nums := getFirstLastDigit(words)
	
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}