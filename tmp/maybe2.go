package main

import "fmt"

func getmaybe() maybe int {
	return struct {
		Val int
		Err error 
	}{
		Val :5,
		Err : nil,
	}
}

func main() {
	fmt.Printf("val: %v\n", getmaybe())
}
