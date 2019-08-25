ABOUT: 

Ego stand for Evan's Go. My name is Evan. Hello! Ego is a (tiny) collection of
exploratory Go extensions. Current additions include:

1. Maybe type.
2. builtin `ok` function.
3. builtin `err` function.
4. `pair` statement.

All current additions are centered around an "either" monad. For those familiar
with Rust's `Result` type the following should be familiar. Here's a minimal 
example:

| package main
| 
| import (
| 	"errors"
| 	"fmt"
| )
| 
| func fact(top int) maybe int {
| 	if top < 0 {
| 		return err(errors.New("must be greated than one"))
| 	}
| 
| 	if top == 0 {
| 		return ok(1)
| 	}
| 
| 	pair fact(top - 1) {
| 		(val int) {
| 			val = val * top
| 			if val == 0 {
| 				return err(errors.New("overflow"))
| 			}
| 			return ok(val)
| 		}
| 		(e error) {
| 			return err(e) // Isn't actually reached in this example.
| 		}
| 	}
| }
| 
| func main () {
| 	pair fact(5) {
| 		(val int) {
| 			fmt.Printf("value: %d\n", val)
| 		}
| 		(e error) {
| 			fmt.Println(e.Error())
| 		}
| 	}
| }

BUILDING: 

Clone repository. `cd` into `src` directory. Run `./make.bash` and binaries will
be ouput to the `bin` directory. If you intend to use Ego it's helpful to add
the `bin` directory to your path.
