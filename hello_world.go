package main

import "fmt"

func Hello(aman string) string {
	return "Hello " + aman
}
func main() {
	fmt.Print(Hello("aman"))
}
