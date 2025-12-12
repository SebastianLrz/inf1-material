package primes

func IsPrime(n int) bool {

	if n < 2 {
		return false
	}

	for i := 2; i < n; i++ {

		if n%i == 0 {
			return false
		}
	}
	return true
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
