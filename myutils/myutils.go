package myutils
type Number interface {
	int | float64 | float32
}
func Sum[T Number](numbers ...T) T {
	var totalSum T

	for _, val := range numbers {
		totalSum += val
	}

	return totalSum
}