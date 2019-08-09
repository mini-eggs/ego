// run

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test switch statements.

package main

import (
	"errors"
)

func assert(cond bool, msg string) {
	if !cond {
		print("assertion fail: ", msg, "\n")
		panic(1)
	}
}

func t1(t bool) maybe string {
	if t {
		return ok("triumph")
	} else {
		return err(errors.New("pair stmt did not pass"))
	}
}

func main() {
	pair t1(true) {
		(val string) {
			// should be called
			assert(true, val)
		}
		(e error) {
			// should not be called
			assert(false, e.Error())
		}
	}
	pair t1(false) {
		(val string) {
			// should not be called
			assert(false, val)
		}
		(e error) {
			// should be called
			assert(true, e.Error())
		}
	}
}
