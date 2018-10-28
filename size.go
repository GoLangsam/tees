// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

// ===========================================================================

// TrixSize returns the sum of AwaySize of e's AwayList
func TrixSize(e *aTee) int {
	size := 0
	for f := e.AwayList().Front(); f != nil; f = f.Next() {
		size += AwaySize(f.AwayList()) // same as f.AwayList().Size()
	}
	return size
}

// AwaySize returns the sum of Len() of the AwayLists along i - same as i.Size()
func AwaySize(i iterator) int {
	size := 0
	for e := i.Front(); e != nil; e = e.Next() {
		size += e.AwayList().Len()
	}
	return size
}
