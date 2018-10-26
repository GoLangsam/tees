// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

import (
	"github.com/GoLangsam/tees/list"
)

// ===========================================================================

// Zipp returns an iterator function returning successive pairs.
func Zipp(X, Y CanIter) func() (*list.Element, *list.Element) {
	var x = X.Front()
	var y = Y.Front()

	return func() (*list.Element, *list.Element) {
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
