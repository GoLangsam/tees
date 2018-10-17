// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"github.com/GoLangsam/tees/list"
	"github.com/GoLangsam/tees/list/iter"
)

// ===========================================================================

// Stick returns a list of lenght N where each element is junk'ed with itself
func Stick(N int) *list.List {
	var list = list.NewList("Stick")
	list.Junk(list)

	for _, id := range IDs("S-", N) {
		list.PushBack(id)
	}

	var next = iter.Zipper(list, list)
	for x, y := next(); x != nil && y != nil; x, y = next() {
		x.Junk(y)
	}
	return list
}

// Ladder returns a list of lenght N where each element is junk'ed
// with the corresponding element of a second list of same lenght N
func Ladder(N int) *list.List {
	var left = list.NewList("LLeft")
	var rigt = list.NewList("Right")
	left.Junk(rigt)

	for _, id := range IDs("L-", N) {
		left.PushBack(id)
	}
	for _, id := range IDs("R-", N) {
		rigt.PushBack(id)
	}

	var next = iter.Zipper(left, rigt)
	for x, y := next(); x != nil && y != nil; x, y = next() {
		x.Junk(y)
	}
	return left
}
