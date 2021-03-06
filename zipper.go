// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

// ===========================================================================

// Zipp returns an iterator function returning successive pairs.
func Zipp(X, Y iterator) func() (*aTee, *aTee) {
	var x = X.Front()
	var y = Y.Front()

	return func() (*aTee, *aTee) {
		var currx = x
		var curry = y
		if x != nil {
			x = x.Next()
		}
		if y != nil {
			y = y.Next()
		}
		return currx, curry
	}
}
