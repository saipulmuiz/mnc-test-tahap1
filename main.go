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

	// Soal 2  - Menghitung ulang kembalian kasir
	total := 700649
	paid := 800000
	fmt.Printf("Jawaban Soal 2: %v\n", soal2CalculateChange(total, paid))

}

func soal1FindMatchingStrings(n int, stringsInput []string) interface{} {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if strings.EqualFold(stringsInput[i], stringsInput[j]) {
				return []int{i + 1, j + 1}
			}
		}
	}
	return false
}

func soal2CalculateChange(total, paid int) interface{} {
	if paid < total {
		return false
	}

	change := (paid - total) / 100 * 100
	denominations := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	resultText := "Pecahan uang: \n"

	for _, denom := range denominations {
		if change >= denom {
			qty := change / denom
			pecahan := ""
			if denom < 1000 {
				pecahan = "koin"
			} else {
				pecahan = "lembar"
			}
			resultText += fmt.Sprintf(" %v %s %v\n", qty, pecahan, denom)
			change %= denom
		}
	}

	return resultText
}
