package arrayslices

func Sum(s []int) int {
	sum := 0
	for _, number := range s {
		sum += number
	}
	return sum
}
