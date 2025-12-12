package primes

func IsPrime(n int) bool {

	for i := 2; i < n; i++ {

		if n%i == 0 {
			return false
		}
	}
	return n < 2
}

// liefert eine Liste mmit allen Primzahlen bis n
func Primes(n int) []int {

	result := []int{}

	for i := 2; i <= n; i++ {

		if IsPrime(i) {
			result = append(result, i)
		}

	}
	return result
}
