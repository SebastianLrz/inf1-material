package aufgabe2

import "strings"

/* AUFGABENSTELLUNG: Vervollständigen Sie unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// FilterDigits erwartet einen String s und liefert einen String zurück,
// der aus s entsteht, indem alle Ziffern entfernt werden.
// Alle anderen Zeichen sollen unverändert bleiben.
func FilterDigits(s string) string {

	s = strings.ReplaceAll(s, "1", "")
	s = strings.ReplaceAll(s, "2", "")
	s = strings.ReplaceAll(s, "3", "")
	s = strings.ReplaceAll(s, "4", "")
	s = strings.ReplaceAll(s, "5", "")
	s = strings.ReplaceAll(s, "6", "")
	s = strings.ReplaceAll(s, "7", "")
	s = strings.ReplaceAll(s, "8", "")
	s = strings.ReplaceAll(s, "9", "")
	s = strings.ReplaceAll(s, "0", "")

	return s
}
