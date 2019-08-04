package main

import (
	"fmt"
)

func t1() maybe int {
	return ok(5)
}

func main() {
	{
		tmp := t1()
		if tmp.Err == nil {
			val := tmp.Val
			fmt.Printf("value: %d\n", val)
		} else {
			e := tmp.Err
			fmt.Println(e.Error())
		}
	}
}

