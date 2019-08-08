package main

import (
	"errors"
	"fmt"
)

func t1(val bool) maybe int {
	if val {
		return ok(5)
	} else { 
		return err(errors.New("this is an error"))
	}
}


func main() {
	pair t1(true) {
		(val int) {
			fmt.Printf("value: %d\n", val)
		}
		(e error) {
			// not reached
			fmt.Println(e.Error())
		}
	}
	pair t1(false) {
		(val int) {
			// not reached
			fmt.Printf("value: %d\n", val)
		}
		(e error) {
			fmt.Println(e.Error())
		}
	}
}
