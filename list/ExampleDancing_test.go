// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"github.com/GoLangsam/tees/list"
	"github.com/GoLangsam/tees/list/test"
)

func ExampleDancing() {
	var d = list.NewDancing()
//	var l = test.Board(100, 100)
//	var l = test.NQueensR(10)
	var l = test.SmallMatrix()
	l.Dance(d)
	// Output:
}
