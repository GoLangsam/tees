// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

// ===========================================================================

// Xross returns a new list: the cross product of X with Y
//	the returned list is a handle to the X-Dimension
//	the elements are Junk'ed
func Xross(X, Y xrosser) *Tees {

	cols := make([]*Tees, 0, X.Len())

	var xl = New(X.CVs())
	for x := X.Front(); x != nil; x = x.Next() {
		cols = append(cols, xl.AddBeam(x.CVs()))
	}

	var yl = New(Y.CVs())
	for y := Y.Front(); y != nil; y = y.Next() {
		yl.AddJunk(y.CVs(), cols...) // use AddOnes to keep elements light (Value = <nil>)
	}
	xl.Join(yl)

	return xl
}
