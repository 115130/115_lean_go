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
