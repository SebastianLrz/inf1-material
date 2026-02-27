package aufgabe4

// ElementProducts erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils das Produkt der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt,
// soll dieses Element ins Ergebnis übernommen werden.
func ElementProducts(l1, l2 []int) []int {
	MaxMult := 0
	out := []int{}

	if len(l1) < len(l2) {
		MaxMult = len(l1)
	} else {
		MaxMult = len(l2)
	}

	for i := 0; i < MaxMult; i++ {
		out = append(out, l1[i]*l2[i])
	}

	if len(l1) < len(l2) {
		out = append(out, l2[len(l1):]...)
	} else if len(l1) > len(l2) {
		out = append(out, l1[len(l2):]...)
	} else {
		return out
	}

	return out
}
