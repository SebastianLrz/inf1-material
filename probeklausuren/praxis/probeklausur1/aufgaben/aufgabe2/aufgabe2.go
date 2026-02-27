package aufgabe2

import "slices"

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	if slices.Index(list, last) == -1 || slices.Index(list, first) == -1 || slices.Index(list, last) < slices.Index(list, first) {
		return []string{}
	} else {
		return slices.Concat(slices.Clone(list[:slices.Index(list, first)]), slices.Clone(list[slices.Index(list, last)+1:]))
	}
}
