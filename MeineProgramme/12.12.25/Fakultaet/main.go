package main

import "fmt"

func GetInput() int {
	fmt.Println("Gib Zahl zwischen 1 und 10")
	var n int
	fmt.Scan(&n)

	if n <= 10 && n > 0 {
		return n
	}

	fmt.Println("Ungültig, versuche es nochmal")
	return GetInput()
}

func main() {
	n := GetInput()
	fmt.Printf("%v ist eine Gute Zahl!", n)
}
