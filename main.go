package main

import (
	"fmt"
	"strings"
)

func main() {
	// Soal 1  - Pencocokan String
	n := 4
	input := []string{"abcd", "acbd", "aaab", "acbd"}
	fmt.Printf("Jawaban Soal 1: %v\n", soal1FindMatchingStrings(n, input))

}

func soal1FindMatchingStrings(n int, stringsInput []string) interface{} {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if strings.ToLower(stringsInput[i]) == strings.ToLower(stringsInput[j]) {
				return []int{i + 1, j + 1}
			}
		}
	}
	return false
}
