package duration

type Duration uint

// Seconds erwartet ein int, das als Sekundenzah interpretiert wird,
func Seconds(m int) Duration {
	return Duration(m)
}

// Minutes erwartet ein int, das als minuten interpretert wird,
// liefert eine Duration für diese MinutenAnzahl
func Minutes(m int) Duration {
	return Duration(m * 60)
}

// liefert die Duration als Sekunden
func (s Duration) ToSeconds() int {
	return int(s)
}

// liefert die Duration als Minuten
func (s Duration) ToMinutes() int {
	return int(s / 60)
}

// SeocondsToMinutes erwatet ine SekundenAnzahl, liefert eine MinutenZahl
func SecondsToMinutes(s int) int {
	return int(Seconds(s).ToMinutes())
}

func Example() {

	// Output:
}
