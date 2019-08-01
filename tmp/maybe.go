package main

import (
	"fmt"
	"errors"
)

func withInt() maybe int {
	return ok(5)
}

func withErr() maybe int {
	return err(errors.New("this is an error"))
}

func fact(top int) maybe int {
	if top < 0 {
		return err(errors.New("can't be negative"))
	} else if top == 1 {
		return ok(1)
	} else { 
		// pair fact( top - 1 ) {
		// 	(res int) {
		// 		return ok(top * res)
		// 	}
		// 	(e error) {
		// 		return err(e)
		// 	}
		// }
		{
			_item := fact(top - 1)
			if _item.Err == nil {
				res := _item.Val
				return ok(top * res)
			} else {
				e := _item.Err
				return err(e)
			}
		}
	}
}

func name() maybe string {
	return ok("Evan J")
}

func main (){
	fmt.Printf("val: %v\n", withInt())
	fmt.Printf("val: %v\n", withErr())
	fmt.Printf("val: %v\n", fact(5))
	fmt.Printf("name: %v\n", name())
}
