package game

import "fmt"

func ReadNumber() int {
	//TODO
	var n int

	fmt.Print("Rate eine Zahl von 1 bis 10: ")
	fmt.Scan(&n)

	return n
}
