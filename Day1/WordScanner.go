package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// func getFirstLastDigit(words []string) []int {
// 	// slice of nums to return, sum of which provide solution
// 	nums := make([]int, 0)

//     for _, word := range words {
// 		re1 := regexp.MustCompile(`\d`)
// 		matches1 := re1.FindAllString(word, -1)

// 		if matches1 != nil {
// 			number, err := strconv.Atoi(matches1[0] + matches1[len(matches1)-1])
// 			if err == nil {
// 				nums = append(nums, number)
// 			}
// 		}
// 	}

// 	return nums
// }

// func regex()  {
// 	word := "oneight"

// 	re1 := regexp.MustCompile(`one?`)
// 	re2 := regexp.MustCompile(`eight?`)

// 	exp := make([]regexp.Regexp, 0)
// 	exp = append(exp, *re1, *re2)

// 	matches := make([][]int, 0)

// 	for _, re := range exp {
// 		match := re.FindAllStringIndex(word, -1)
// 		matches = append(matches, match...)
// 	}

// 	first, last := []int{-1, -1}, []int{-1, -1}
// 	for _, val := range matches {
// 		if first[0] == -1 {
// 			first = val
// 			last = val
// 		}
// 		if val[0] < first[0] {
// 			first = val
// 		}
// 		if val[0] > last[0] {
// 			last = val
// 		}
// 	}

// 	fmt.Println(word[:first[1]])
// 	fmt.Println(word[last[0]:last[1]])
// }

func getFirstLastDigitOrString(words []string) []int {
	re1 := regexp.MustCompile(`one`)
	re2 := regexp.MustCompile(`two`)
	re3 := regexp.MustCompile(`three`)
	re4 := regexp.MustCompile(`four`)
	re5 := regexp.MustCompile(`five`)
	re6 := regexp.MustCompile(`six`)
	re7 := regexp.MustCompile(`seven`)
	re8 := regexp.MustCompile(`eight`)
	re9 := regexp.MustCompile(`nine`)
	re10 := regexp.MustCompile(`zero`)
	re11 := regexp.MustCompile(`\d`)
	digitString := []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	exprs := []*regexp.Regexp{
		re1, re2, re3, re4, re5, re6, re7, re8, re9, re10, re11,
	}
	nums := make([]int, 0)

	for _, word := range words {
		matches := make([][]int, 0)
		for _, expr := range exprs {
			match := expr.FindAllStringIndex(word, -1)
			matches = append(matches, match...)
		}
		first, last := []int{-1, -1}, []int{-1, -1} 
		for _, val := range matches {
			if first[0] == -1 {
				first = val
				last = val
			}
			if val[0] < first[0] {
				first = val
			}
			if val[0] > last[0] {
				last = val
			}
		}
		firstDigit, lastDigit := word[first[0]:first[1]], word[last[0]:last[1]]
		fmt.Println(word + ": " + word[first[0]:first[1]] + ", " + word[last[0]:last[1]])

		digit, err := strconv.Atoi(firstDigit + lastDigit)
		if err == nil {
			nums = append(nums, digit)
		}
		if err != nil {

			for i := 0; i < len(digitString); i++ {
				if firstDigit == digitString[i] {
					firstDigit = strconv.Itoa(i)
				}
				if  lastDigit == digitString[i] {
					lastDigit = strconv.Itoa(i)
				}
			}
			digit, err := strconv.Atoi(firstDigit + lastDigit)
			if err == nil {
				nums = append(nums, digit)
			}
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
	// regex()

	words := getWordsSlice("input.txt")

	// test := []string{
	// 	"8sdonlkn3lkns3",
	// 	"jnw33klndlknsa",
	// 	"73klndskndstwofour",
	// 	"nknlds3oneight",
	// } //sum = 222, 228

	// // nums := getFirstLastDigit(test)
	nums := getFirstLastDigitOrString(words)

	
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}