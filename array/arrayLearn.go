package array

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}
func SumAll(numbers ...[]int) (sum []int) {

	for _, value := range numbers {
		sum = append(sum, Sum(value))
	}
	return
}
func SumAllTails(numbers ...[]int) (sums []int) {
	for _, value := range numbers {
		tail := value[1:]
		sums = append(sums, Sum(tail))
	}
	return
}
