package game

import "fmt"

func ReadNumber() int {
	//TODO
	var n int

	fmt.Print("Rate eine Zahl: ")
	fmt.Scan(&n)


	return n
}