package main

import (
	"errors"
)

func t1() maybe int {
	return struct{ Err error ; Val int }{Err: errors.New("error here")}
}

func main () {
	t1()
}
