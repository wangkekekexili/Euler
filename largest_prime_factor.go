package euler

func largestPrimeFactor(num int64) int64 {
	if num <= 1 {
		return 0
	}

	var currentLargestPrime int64 = 1
	for i := int64(2); i <= num && num > 1; i++ {
		if num%i == 0 {
			currentLargestPrime = i
			for num%i == 0 {
				num /= i
			}
		}
	}

	return currentLargestPrime
}
