package constants

const (
	PI = 3.14
	DB_URL = "localhost:5432"
)

func sum[T int](numbers ...T) T {
	var totalSum T

	for _, val := range numbers {
		totalSum += val
	}
	return totalSum
}