package digits

import (
	"fmt"
	"slices"
)

//Bin Digits erwartet eine Zahl n und liefert ene Liste von Ziffern

func BinDigits(n int) []int {
	result := []int{}

	for n != 0 {
		last_digit := n % 2
		result = append(result, last_digit)
		//result = append([]int{last_digit}, result...) //direkt umgekehrt anhängen
		n /= 2 //n = n/2
	}
	slices.Reverse(result)

	return result
}

// erwartet Zahl und Basis, liefert Liste der Ziffern aus Zahl in neuer Basis
func Digits(n, base int) []int {
	result := []int{}

	for n != 0 {
		last_digit := n % base
		result = append(result, last_digit)
		//result = append([]int{last_digit}, result...) //direkt umgekehrt anhängen
		n /= base //n = n/2
	}
	slices.Reverse(result)

	return result
}

// Sum erwawrtet eine Liste von Zahlen und berechnet deren Summe
func Sum(numbers []int) int {
	result := 0
	for _, n := range numbers {
		result += n
	}
	return result
}

// ParityBit erwartet Zahl n und liefert
// 1: wenn anzahl der Einsen in Binärdarstellung von n ungerade ist
// 0: wenn anzahl der Einsen in Binärdarstellung von n gerade ist
func ParityBit(n int) int {
	//return Sum(Digits(n, 2)) % 2
	return DigitSum(n, 2)
}

// DigitSum berechnet die Quersumme von n zur gegebenen basis
func DigitSum(n, base int) int {
	return Sum(Digits(n, base))
}

func ExampleBinDigits() {
	fmt.Println(Digits(42, 2))
	fmt.Println(Digits(42, 16))
	fmt.Println(Digits(42, 10))
	fmt.Println(Digits(42, 8))

	fmt.Println(ParityBit(32))

	// Output:
	// [1 0 1 0 1 0]
	// [2 10]
	// [4 2]
	// [5 2]
	// 1
}
