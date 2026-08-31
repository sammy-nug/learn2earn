package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		nb = 2
	}

	for {
		isPrime := true

		if nb > 3 {
			if nb%2 == 0 {
				isPrime = false
			} else {
				for i := 3; i*i <= nb; i += 2 {
					if nb%i == 0 {
						isPrime = false
						break
					}
				}
			}
		}

		if isPrime {
			return nb
		}

		nb++
	}
}
