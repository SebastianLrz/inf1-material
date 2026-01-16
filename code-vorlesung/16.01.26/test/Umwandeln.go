package Umwandeln

import "strconv"

func main(x, base, wantBase int) int {
	if base == 10 {
		println(dezUmw(x, wantBase))
	} else {
		dezUmw(UmwDez(x, base), 10)
	}

	return 0
}

func ListeInZahl(zahlListe []int) int {
	str := ""
	for _, digit := range zahlListe {
		str += strconv.Itoa(digit)
	}
	Zahl, _ := strconv.Atoi(str)
	return Zahl
}

func dezUmw(n, base int) int {
	result := []int{}

	for n != 0 {
		last_digit := n % base
		result = append([]int{last_digit}, result...) //direkt umgekehrt anhängen
		n /= base
	}
	return ListeInZahl(result)
}

func UmwDez(n, base int) int {
	return 0 //TODO: Wandelt Zahlen jeder Basis in Dez um, gibt (maybe mit ListeInZahl) Int aus
}
