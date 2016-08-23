package euler

func sumOfMultiplesOf3And5Below(target int) int {
	sum := 0
	for i := 1; i != target; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
