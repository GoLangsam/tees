// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"github.com/GoLangsam/do/id"
	"github.com/GoLangsam/tees"
	"github.com/GoLangsam/tees/list"
)

// ===========================================================================

// Stick returns a list of length N where each element is junk'ed with itself
func Stick(N int) *list.List {
	var list = tees.New("Stick")
	list.Join(list)

	for _, id := range id.S("S-", N) {
		list.PushBack(id)
	}

	var next = tees.Zipp(list, list)
	for x, y := next(); x != nil && y != nil; x, y = next() {
		x.Junk(y)
	}
	return list
}

// Ladder returns a list of length N where each element is junk'ed
// with the corresponding element of a second list of same length N
func Ladder(N int) *list.List {
	var left = tees.New("LLeft")
	var rigt = tees.New("Right")
	left.Join(rigt)

	for _, id := range id.S("L-", N) {
		left.PushBack(id)
	}
	for _, id := range id.S("R-", N) {
		rigt.PushBack(id)
	}

	var next = tees.Zipp(left, rigt)
	for x, y := next(); x != nil && y != nil; x, y = next() {
		x.Junk(y)
	}
	return left
}
