package prime_numbers

func PrimeNumbers(a, b int) []int {
	var result []int

	// Проверка на простое число
	for i := a; i <= b; i++ {
		// Простые числа начинаются с 2
		if i < 2 {
			continue
		}

		isPrime := true
		// Проверка на деление на числа от 2 до i-1
		for j := 2; j < i; j++ {
			// Если i делится на j без остатка, то i не простое число
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			result = append(result, i)
		}
	}

	return result
}
