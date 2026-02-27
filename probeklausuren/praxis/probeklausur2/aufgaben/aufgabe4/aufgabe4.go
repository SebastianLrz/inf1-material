package aufgabe4

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// ElementSums erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils die Summe der beiden Elemente enthält.
//
// Annahmen für die Berechnung:
// Falls eine Liste kürzer ist als die andere, soll für die Berechnung der
// hinteren Werte ihr letztes Element verwendet werden.
// Für leere Listen soll für die Berechnung ggf. 0 verwendet werden.
func ElementSums(l1, l2 []int) []int {
	output := []int{}
	AddLen := 0
	var l1Short bool = false
	var l2Short bool = false

	if len(l1) <= len(l2) {
		AddLen = len(l1)
		l1Short = true
	} else {
		AddLen = len(l2)
		l2Short = true
	}

	for i := 0; i < AddLen; i++ {
		output = append(output, l1[i]+l2[i])
	}
	if l1Short {
		for i := AddLen; i < len(l2); i++ {
			output = append(output, l2[i]+l1[len(l1)-1])
		}
	}
	if l2Short {
		for i := AddLen; i < len(l1); i++ {
			output = append(output, l1[i]+l2[len(l2)-1])
		}
	}
	return output
}
