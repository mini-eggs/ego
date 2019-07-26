package main

import "fmt"

func getmaybe() maybe int {
	return ok(5)
}

func main() {
	fmt.Printf("val: %v\n", getmaybe())
}
