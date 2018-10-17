// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iter

import (
	"github.com/GoLangsam/tees/list"
)

// ===========================================================================

// Zipper returns an iterator over pairs of (x,y)'s
func Zipper(X, Y *list.List) func() (*list.Element, *list.Element) {
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
