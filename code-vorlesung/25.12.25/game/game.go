package game

import "fmt"

func Run() {

	for n := 0; n < 3; n++ {

		guess := ReadNumber()

		if NumberGood(guess) {

			fmt.Println("Richtige Antwort, Sie haben gewonnen!")
			return

		}
		fmt.Println("Falsch, versuchen sie es erneut.")

	}
	fmt.Println("Game Over, Sie haben verloren.")

}
