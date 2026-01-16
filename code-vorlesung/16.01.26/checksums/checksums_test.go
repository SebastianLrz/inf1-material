package checksums

import "fmt"

func Sum(numbers []int) int {
	result := 0

	for i, n := range numbers {
		if i%2 == 0 {
			result += n
		} else {
			result += 3 * n
		}

		// result += (n + 2*n*(i+1)%2)   macht alles auf einmal anstatt der if schleife (schneller)
	}
	return result
}

// 978-0-345-39180-3

//EAN erwartet Liste von Ziffern
//liefert prüfziffer dazu

func EAN(digits []int) int {
	return 10 - (Sum(digits)%10)%10 //2. modulo10 ist falls SumDigits 0 ist

}

func ExampleEAN() {
	fmt.Println(EAN([]int{9, 7, 8, 0, 3, 4, 5, 3, 9, 1, 8, 0}))

	// Output:
	// 3
}
