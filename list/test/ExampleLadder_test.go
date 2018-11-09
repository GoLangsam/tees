// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

func ExampleLadder() {
	var N = 7
	var l = Ladder(N)

	l.PrintAways("Ladder")
	l.AwayList().PrintAways("Ladder-Away")
	// Output:
	// Ladder: List=LLeft | Total=7
	// => : List=Right	R-1 | R-2 | R-3 | R-4 | R-5 | R-6 | R-7 | Total=7
	// => : List=Right	R-1 | R-2 | R-3 | R-4 | R-5 | R-6 | R-7 | Total=7
	// => : List=Right	R-1 | R-2 | R-3 | R-4 | R-5 | R-6 | R-7 | Total=7
	// => : List=Right	R-1 | R-2 | R-3 | R-4 | R-5 | R-6 | R-7 | Total=7
	// => : List=Right	R-1 | R-2 | R-3 | R-4 | R-5 | R-6 | R-7 | Total=7
	// => : List=Right	R-1 | R-2 | R-3 | R-4 | R-5 | R-6 | R-7 | Total=7
	// => : List=Right	R-1 | R-2 | R-3 | R-4 | R-5 | R-6 | R-7 | Total=7
	// Ladder-Away: List=Right | Total=7
	// => : List=LLeft	L-1 | L-2 | L-3 | L-4 | L-5 | L-6 | L-7 | Total=7
	// => : List=LLeft	L-1 | L-2 | L-3 | L-4 | L-5 | L-6 | L-7 | Total=7
	// => : List=LLeft	L-1 | L-2 | L-3 | L-4 | L-5 | L-6 | L-7 | Total=7
	// => : List=LLeft	L-1 | L-2 | L-3 | L-4 | L-5 | L-6 | L-7 | Total=7
	// => : List=LLeft	L-1 | L-2 | L-3 | L-4 | L-5 | L-6 | L-7 | Total=7
	// => : List=LLeft	L-1 | L-2 | L-3 | L-4 | L-5 | L-6 | L-7 | Total=7
	// => : List=LLeft	L-1 | L-2 | L-3 | L-4 | L-5 | L-6 | L-7 | Total=7
}
