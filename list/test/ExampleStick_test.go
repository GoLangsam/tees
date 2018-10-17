// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test_test

import (
	"github.com/GoLangsam/tees/list/test"
)

func ExampleStick() {
	var N = 7
	var s = test.Stick(N)

	s.PrintAways("Stick")
	s.AwayList().PrintAways("Stick-Away")
	// Output:
	// Stick: List=Stick | Total=7
	// => : List=Stick	S-1 | S-2 | S-3 | S-4 | S-5 | S-6 | S-7 | Total=7
	// => : List=Stick	S-1 | S-2 | S-3 | S-4 | S-5 | S-6 | S-7 | Total=7
	// => : List=Stick	S-1 | S-2 | S-3 | S-4 | S-5 | S-6 | S-7 | Total=7
	// => : List=Stick	S-1 | S-2 | S-3 | S-4 | S-5 | S-6 | S-7 | Total=7
	// => : List=Stick	S-1 | S-2 | S-3 | S-4 | S-5 | S-6 | S-7 | Total=7
	// => : List=Stick	S-1 | S-2 | S-3 | S-4 | S-5 | S-6 | S-7 | Total=7
	// => : List=Stick	S-1 | S-2 | S-3 | S-4 | S-5 | S-6 | S-7 | Total=7
	// Stick-Away: List=Stick | Total=7
	// => : List=Stick	S-1 | S-2 | S-3 | S-4 | S-5 | S-6 | S-7 | Total=7
	// => : List=Stick	S-1 | S-2 | S-3 | S-4 | S-5 | S-6 | S-7 | Total=7
	// => : List=Stick	S-1 | S-2 | S-3 | S-4 | S-5 | S-6 | S-7 | Total=7
	// => : List=Stick	S-1 | S-2 | S-3 | S-4 | S-5 | S-6 | S-7 | Total=7
	// => : List=Stick	S-1 | S-2 | S-3 | S-4 | S-5 | S-6 | S-7 | Total=7
	// => : List=Stick	S-1 | S-2 | S-3 | S-4 | S-5 | S-6 | S-7 | Total=7
	// => : List=Stick	S-1 | S-2 | S-3 | S-4 | S-5 | S-6 | S-7 | Total=7
}
