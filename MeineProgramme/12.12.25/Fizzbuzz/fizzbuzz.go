package main

import "fmt"

func main() {
	//fizzbuzz_v1()
	//fizzbuzz_v3()

	var a int
	fmt.Print("Wähle x!")
	fmt.Scan(&a)

	var b int
	fmt.Print("Wähle y!")
	fmt.Scan(&b)

	var m int
	fmt.Print("Bis wohin soll gezählt werden?")
	fmt.Scan(&m)

	fizzbuzz_allgemein(a, b, m)
}

func fizzbuzz_v1() {

	for i := 0; i < 45; i++ {

		if i%5 == 0 && i%7 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("fizz")
		} else if i%7 == 0 {
			fmt.Println("bus")
		} else {
			fmt.Println(i)

		}
	}
}

func fizzbuzz_v3() {
	for i := 0; i < 42; i++ {

		if i%5 == 0 && i%7 == 0 {
			fmt.Println("fizzbuzz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("fizz")
			continue
		}

		if i%7 == 0 {
			fmt.Println("buzz")
			continue
		}

	}
}

func fizzbuzz_v4() {
	for i := 0; i < 42; i++ {

		printedsmth := false

		if i%5 == 0 {
			fmt.Print("fizz")
			printedsmth = true
			continue
		}

		if i%7 == 0 {
			fmt.Print("buzz")
			printedsmth = true
			continue
		}

		if !printedsmth {
			fmt.Print(i)
		}
		fmt.Println()

	}
}

func fizzbuzz_allgemein(a, b, m int) {

	for i := 1; i < m; i++ {

		if i%a == 0 && i%b == 0 {
			fmt.Println("fizzbuzz")
			continue
		}

		if i%a == 0 {
			fmt.Println("fizz")
			continue
		}

		if i%b == 0 {
			fmt.Println("buzz")
			continue
		}
		fmt.Println(i)
	}
}
