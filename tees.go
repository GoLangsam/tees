// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

import (
	"github.com/GoLangsam/tees/list"
)

// ========================================================

// New returns a new collection of Tees.
func New(v interface{}, vals ...interface{}) Tees {
	return list.NewList(v, vals...)
}

// ========================================================

// This is it - an indivisible individuum, an atom.
type This = *list.Element

// Tees is a collection of This'ses.
type Tees = *list.List

// ========================================================

// Dancing holds what This needs in order to Dance.
type Dancing = *list.Dancing

// A Dancer loves dancing his Dance.
type Dancer interface {
	Away() *list.Element
	ForEachNext(f func(*list.Element))
	ForEachPrev(f func(*list.Element))
}

// ========================================================

// Iterator allows to iterate forward by starting with Front() and, if non-nil, repeating Next() until Next() returns nil
type Iterator interface {
	Front() *list.Element
	Next() *list.Element
	ForEachNext(f func(*list.Element))
}

// Xrosser supports constructions such as Xross
type Xrosser interface {
	Iterator
	CVs() *list.ComposedValue
	Len() int
}

// Calcer supports calculations such as Append and Times
type Calcer interface {
	Iterator
	CVs() *list.ComposedValue
	With(*list.List) *list.ComposedValue
}

// A Line may be parallel to another Line
type Line interface {
	List() *list.List
	Root() *list.Element
}

// A Border is directly next to any Edge
type Border interface {
	Away() *list.Element
	IsRoot() bool
}

