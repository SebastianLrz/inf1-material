package game

//NumberGood liefert true falls die geratene Zahl richtig ist

func NumberGood(guess, x int) bool {

	if guess == x {

		return true

	} else {

		return false
	}

}
