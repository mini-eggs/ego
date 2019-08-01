package main

import (
	"errors"
	"fmt"
)

func getmaybe() maybe int {
	return ok(5)
}

func getnil() maybe int {
	// return err("this form is also acceptable")
	return err(errors.New("this is an error"))
}

func main() {
	match getmaybe() {
		(val int) {
			fmt.Printf("value: %d\n", val)
		}
		(e error) {
			// not reached
		}
	}
	match getnil() {
		(val int) {
			// not reached
		}
		(e error) {
			fmt.Println(e.Error())
		}
	}
}
