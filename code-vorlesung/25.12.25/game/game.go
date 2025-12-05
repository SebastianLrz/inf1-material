package game

import "fmt"

func Run() {

	guess := ReadNumber()

	if NumberGood(guess) {

		fmt.Println("Richtige Antwort")

	} else {

		fmt.Println("Zu viele falsche Antworten, Game Over")

	}
}
