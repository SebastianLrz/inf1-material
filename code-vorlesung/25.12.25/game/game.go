package game

import (
	"fmt"
	"math/rand"
)

func Run() {

	x := rand.Intn(10) + 1
	//fmt.Println(x)

	for n := 0; n < 3; n++ {

		guess := ReadNumber()

		if NumberGood(guess, x) {

			fmt.Println("Richtige Antwort, Sie haben gewonnen!")
			return

		} else {

			if HigherLower(x, guess) {
				fmt.Println("Deine Zahl war zu niedrig!")
			} else {
				fmt.Println("Deine Zahl war zu hoch!")
			}

		}
	}
	fmt.Println("Game Over, Sie haben verloren.")

}
