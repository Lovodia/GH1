package mymath

func Average(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0.0
	for _, n := range numbers {
		sum += n
	}
	return sum / float64(len(numbers))
}
