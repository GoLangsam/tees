// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

// ===========================================================================
// boolean - unary

// IsEdge determines, if it makes sense to go further upward from x
func IsEdge(x Border) bool {
	//	if x == nil { return false }
	//	if x.IsNode() { return false }
	//	if x.IsSolo() { return true }
	return (x != nil && x.IsRoot() && x.Away() != nil && x.Away().IsNode())
}

// ===========================================================================
// boolean - binnary

// AreParallel reports if two lines are parallel with each other.
//	Note: nil is NOT considered parallel to nil
func AreParallel(x, y Line) bool {
	if x == nil { return false }
	if y == nil { return false }

	return (x.List() == y.List() || x.Root().AwayList() == y.Root().AwayList())
}
