package main

import (
	"errors"
	"fmt"
)

func t1() maybe int {
	// return ok(5)
	return err(errors.New("this is not a triumph"))
}

func main() {
	// pair t1() {
	// 	(val int) {
	// 		fmt.Printf("value: %d\n", val)
	// 	}
	// 	(e error) {
	// 		fmt.Println(e.Error())
	// 	}
	// }
	if tmp := t1() ; tmp.Err == nil {
		val := tmp.Val
		fmt.Printf("value: %d\n", val)
	} else {
		e := tmp.Err
		fmt.Println(e.Error())
	}
}

