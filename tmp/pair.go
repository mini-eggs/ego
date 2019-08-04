package main

import (
	"fmt"
)

func t1() maybe int {
	return ok(5)
}

func main() {
	pair t1() {
		(val int) {
			// prints `value: 5`
			fmt.Printf("value: %d\n", val)
		}
		(e error) {
			// not reached in this example
		}
	}
}
