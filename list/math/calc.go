// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package math provides operations on lists such as addition and multiplication
	- Append( l, lists...)		*list.List
	- Times( l, lists...)		*list.List
	- Xross( X, Y *List)		*list.List (as handle X-dimension of matrix)
which return a new list or matrix with composedValues
*/
package math

import (
	"github.com/GoLangsam/tees/list"
)

// ===========================================================================

// Xross returns a new list: the cross product of X with Y
//	the returned list is a handle to the X-Dimension
//	the elements are Junk'ed
// Note: such Xross may be explored by Step Jump Walk Tour ...
func Xross(X, Y *list.List) *list.List {

	cols := make([]*list.List, 0, X.Len())

	var xl = list.NewList(X.CVs())
	for x := X.Front(); x != nil; x = x.Next() {
		cols = append(cols, xl.AddBeam(x.CVs()))
	}

	var yl = list.NewList(Y.CVs())
	for y := Y.Front(); y != nil; y = y.Next() {
		yl.AddJunk(y.CVs(), cols...) // use AddOnes to keep elements light (Value = <nil>)
	}
	xl.Root().Junk(yl.Root())

	return xl
}
