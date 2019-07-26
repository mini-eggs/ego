package main

import "fmt"

// func getmaybe() maybe int {
// 	return ok(5)
// }
// ->
func getmaybe() struct {
	Val int
	Err error
} {
	return struct {
		Val int
		Err error
	}{
		Val: 5,
		Err: nil,
	}
}

func main() {
	fmt.Printf("val: %v\n", getmaybe())
}

// type: &{int {{{0xc000372000 10 6}}}}
// name: &{Val {{{0xc000372000 10 2}}}}
// pos: ../tmp/struct.go:10:2
// type: &{error {{{0xc000372000 11 6}}}}
// name: &{Err {{{0xc000372000 11 2}}}}
// pos: ../tmp/struct.go:11:2
// type: &{int {{{0xc000372000 18 6}}}}
// name: &{Val {{{0xc000372000 18 2}}}}
// pos: ../tmp/struct.go:18:2
// type: &{error {{{0xc000372000 19 6}}}}
// name: &{Err {{{0xc000372000 19 2}}}}
// pos: ../tmp/struct.go:19:2
// type: &{int {{{0xc000372000 27 6}}}}
// name: &{Val {{{0xc000372000 27 2}}}}
// pos: ../tmp/struct.go:27:2
// type: &{error {{{0xc000372000 28 6}}}}
// name: &{Err {{{0xc000372000 28 2}}}}
// pos: ../tmp/struct.go:28:2
