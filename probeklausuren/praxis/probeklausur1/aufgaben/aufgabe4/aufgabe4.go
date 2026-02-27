package aufgabe4

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion MaxElements.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein!
*/

// MaxElements erwartet zwei int-Listen und liefert eine Liste, die an jeder Position
// jeweils das größere der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt, gilt dieses Element als das größere.
func Angleichen(l1, l2 []int) []int {
	if len(l1) < len(l2) {
		l0 := 0
		difference := len(l2) - len(l1)
		for i := 0; i <= difference; i++ {
			l1 = append(l1, l0)
		}
	} else {
		return l1
	}

}

func MaxElements(l1, l2 []int) []int {

	l1 = Angleichen(l1, l2)
	l2 = Angleichen(l2, l1)

	if len(l1) == 0 || len(l2) == 0 {
		return []int{}
	}

	if l1[0] > l2[0] {
		out := append(MaxElements(l1[1:], l2[1:]), l1[0])
		return out
	} else {
		out := append(MaxElements(l1[1:], l2[1:]), l2[0])
		return out
	}
}
