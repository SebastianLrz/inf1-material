package aufgabe4

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion MaxElements.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein!
*/

// MaxElements erwartet zwei int-Listen und liefert eine Liste, die an jeder Position
// jeweils das größere der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt, gilt dieses Element als das größere.
func MaxElements(l1, l2 []int) []int {
	if len(l1) == 0 {
		return l2
	}
	if len(l2) == 0 {
		return l1
	}

	// List MaxElements aus l1 und l2
	// if pos nur in einer l, append es.  | if lenl1=0 return l2
	greater := l1[0] // und andersrum als base
	if l2[0] > greater {
		greater = l2[0]
	}
	return append([]int{greater}, MaxElements(l1[1:], l2[1:])...)
}
