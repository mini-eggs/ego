package main

import "fmt"

func getmaybe() maybe int {
	return struct { Val int; Err error }{5, nil}
}

func main() {
	fmt.Printf("val: %v\n", getmaybe())
}
