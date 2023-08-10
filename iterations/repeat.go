package iterations

func Repeat(char string, number int) string {
	// takes input a char and number and repeats string number times
	var repeated string
	for i := 0; i < number; i++ {
		repeated += char
	}
	return repeated
}
