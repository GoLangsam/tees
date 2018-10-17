// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test_test

import (
	"github.com/GoLangsam/tees/list/test"
)

func ExampleBoard() {
	var N = 4
	var M = 5
	var board = test.Board(N, M)
	board.PrintAways("Board")
	board.AwayList().PrintAways("Board-Away")
	// Output:
	// Board: List=Ranks | Total=4
	// => : List=R-1	R-1|F-1 | R-1|F-2 | R-1|F-3 | R-1|F-4 | R-1|F-5 | Total=5
	// => : List=R-2	R-2|F-1 | R-2|F-2 | R-2|F-3 | R-2|F-4 | R-2|F-5 | Total=5
	// => : List=R-3	R-3|F-1 | R-3|F-2 | R-3|F-3 | R-3|F-4 | R-3|F-5 | Total=5
	// => : List=R-4	R-4|F-1 | R-4|F-2 | R-4|F-3 | R-4|F-4 | R-4|F-5 | Total=5
	// Board-Away: List=Files | Total=5
	// => : List=F-1	F-1|R-1 | F-1|R-2 | F-1|R-3 | F-1|R-4 | Total=4
	// => : List=F-2	F-2|R-1 | F-2|R-2 | F-2|R-3 | F-2|R-4 | Total=4
	// => : List=F-3	F-3|R-1 | F-3|R-2 | F-3|R-3 | F-3|R-4 | Total=4
	// => : List=F-4	F-4|R-1 | F-4|R-2 | F-4|R-3 | F-4|R-4 | Total=4
	// => : List=F-5	F-5|R-1 | F-5|R-2 | F-5|R-3 | F-5|R-4 | Total=4
}
