// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

import (
	"github.com/GoLangsam/tees/list"
)

// Spot abstracts the common behaviour of *list.Element & *list.List
type Spot interface {
	Away
	Beam
	Dust
	Sama
}

// Away abstracts the 'orthogonal' behaviour common to *list.Element & *list.List
type Away interface {
	AwayBeam
	AwayDust
}

// AwayBeam abstracts the 'orthogonal' beam-behaviour common to *list.Element & *list.List
type AwayBeam interface {
	AwayList() *list.List
	AwayLists() []*list.List

	AwayNext() *list.Element
	AwayPrev() *list.Element

	Home() *list.Element

	PrintAways(args ...interface{})
	Size() int
}

// AwayDust abstracts the 'orthogonal' dust-behaviour common to *list.Element & *list.List
type AwayDust interface {
	Away() *list.Element
	IsJunk() bool
	IsSolo() bool
}

// Beam abstracts the 'lengthy' behaviour common to *list.Element & *list.List
type Beam interface {
	CanIter // Front Next
	CanReti // Back  Prev
	HasLen  // Len
	HasRoot // Root
}

// Dust abstracts the 'pointy' behaviour common to *list.Element & *list.List
type Dust interface {
	CVs() *list.ComposedValue

	IsAtom() bool
	IsComposed() bool

	PrintAtomValues(args ...interface{})
	PrintValue(args ...interface{})
	//	Values() list.Values
}

// Sama abstracts the 'dancing' behaviour common to *list.Element & *list.List
type Sama interface {
	Dance(d *list.Dancing)
	DanceFast(d *list.Dancing)
	DanceSlow(d *list.Dancing)

	UnLeaf(d *list.Dancing)
	ReLeaf(d *list.Dancing)

	UnLink()
	ReLink()
}
