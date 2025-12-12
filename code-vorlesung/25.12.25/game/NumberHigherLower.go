package game

func HigherLower(x, guess int) bool {
	if guess < x {
		return true
	} else {
		return false
	}
}
