package arrayslices

func Sum(s []int) int {
	sum := 0
	for _, number := range s {
		sum += number
	}
	return sum
}

// Write a func sumAll which will take a varying number of slices, returning a new slice containing the totals for each slice passed in.
func sumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums

}
