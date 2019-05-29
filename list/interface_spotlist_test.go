// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"github.com/GoLangsam/tees/list"
)

func Example_nova() {
	var _ Nova = list.New()
	var _ Node = list.New().Root()
}

// Nova is a spot which can create new elements and Remove existing ones
// and thus abstracts common behaviour of *list.List and *anda.Anda, Aton, Apep ...
type Nova interface {
	Spot
	Coll
}

// Node is an atomic spot.
type Node interface {
	Spot
	Atom
}
