package main

import (
	"fmt"
	"strings"
	"unicode"
)

func reverseString(str string) string {
	var strWord []rune
	var strNum []rune

	for _, char := range str {
		if unicode.IsLetter(char) {
			strWord = append(strWord, char)
		} else if unicode.IsDigit(char) {
			strNum = append(strNum, char)
		}
	}

	for i, j := 0, len(strWord)-1; i < j; i, j = i+1, j-1 {
		strWord[i], strWord[j] = strWord[j], strWord[i]
	}

	return string(strWord) + string(strNum)
}

func longest(str string) string {
	words := strings.Fields(str)

	var word string
	var length int

	for _, v := range words {
		if len(v) > length {
			word = v
			length = len(v)
		}
	}

	return word
}

func wordCount(INPUT, QUERY []string) []int {
	var tmp = make(map[string]int)

	for _, word := range QUERY {
		tmp[word] = 0
	}

	for _, word := range INPUT {
		if _, exists := tmp[word]; exists {
			tmp[word]++
		}
	}

	result := make([]int, len(QUERY))

	for i, word := range QUERY {
		result[i] = tmp[word]
	}

	return result
}

func main() {
	// Soal Nomor 1
	var s = "NEGIE1"
	fmt.Println(reverseString(s))

	// Soal Nomor 2
	const sentence = "Saya sangat senang mengerjakan soal algoritma"
	fmt.Println(longest(sentence))

	// Soal Nomor 3
	INPUT := []string{"xc", "dz", "bbb", "dz"}
	QUERY := []string{"bbb", "ac", "dz"}
	fmt.Println(wordCount(INPUT, QUERY))
}
