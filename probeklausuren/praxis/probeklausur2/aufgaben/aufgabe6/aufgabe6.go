package aufgabe6

import (
	"slices"
)

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// DuplicateSinglets erwartet eine int-Liste list.
// Die Funktion liefert eine int-Liste, bei der alle Elemente,
// die in list nur einmal vorkommen, verdoppelt sind,
// also zwei Mal hintereinander stehen.
// Elemente, die schon in list mehrfach vorkommen, sollen wie sie sind
// ins Ergebnis übertragen werden.
func DuplicateSinglets(list []int) []int {
	output := slices.Clone(list)

	for i := 0; i < len(list); i++ {
		x := slices.Clone(list)

		x = slices.Delete(x, i, i+1)

		if slices.Contains(x, list[i]) == false {
			output = slices.Insert(output, i, list[i])
		}
		x = []int{}
	}

	return output
}
