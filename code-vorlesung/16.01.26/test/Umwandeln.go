package Umwandeln

func main(x, base, wantBase int) int {
	if base == 10 {
		dezUmw(x, wantBase)
		return 1
	}

	return 0
}

func ListeInZahl([]int) int {
	return 0
}

func dezUmw(n, base int) []int {
	result := []int{}

	for n != 0 {
		last_digit := n % base
		result = append([]int{last_digit}, result...) //direkt umgekehrt anhängen
		n /= base                                     //n = n/2
	}

	return result
}
