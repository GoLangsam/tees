// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"github.com/GoLangsam/tees/list"
)

func Example_spot() {
	var _ Spot = list.New()
	var _ Spot = list.New().Root()
}

// Spot abstracts the common behaviour of *list.Element & *list.List
type Spot interface {
	node
	Away
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
