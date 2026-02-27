package aufgabe3

import "math"

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * RANDBEDINGUNG: Die Funktion muss rekursiv sein.
 * ERREICHBARE PUNKTE: 10
 */

// CountSquares erwartet eine Liste von Zahlen.
// CountSquares liefert die Anzahl der QuadratzahlenZahlen in der Liste.

func CountSquares(list []int) int {
	if len(list) == 0 {
		return 0
	}

	if math.Sqrt(float64(list[0])) == math.Round(math.Sqrt(float64(list[0]))) {
		return CountSquares(list[1:]) + 1
	} else {
		return CountSquares(list[1:])
	}
}
