// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

import (
	"github.com/GoLangsam/tees/list"
)

// ========================================================

// New returns a new collection of Tees.
func New(vals ...interface{}) *Tees {
	return list.NewList(vals...)
}

// ========================================================

// aTee is it - an indivisible individuum, an atom.
type aTee = list.Element

// Tees is a collection of aTee's.
type Tees = list.List

// This is inside our Tees-Monad
type This = aTee

// ========================================================

// Dancing holds what it needs in order to Dance.
type Dancing = *list.Dancing

// A dancer loves dancing his Dance.
type dancer interface {
	Away() *list.Element
	ForEachNext(f func(*list.Element))
	ForEachPrev(f func(*list.Element))
}

// ========================================================

// iterator allows to iterate forward by starting with Front() and, if non-nil, repeating Next() until Next() returns nil
type iterator interface {
	Front() *list.Element
	Next() *list.Element
	ForEachNext(f func(*list.Element))
}

// trailer represents a finite iterator with known length
type trailer interface {
	iterator
	Len() int
}

// xrosser supports constructions such as Xross
type xrosser interface {
	trailer
	CVs() *list.ComposedValue
}

// calcer supports calculations such as Append and Times
type calcer interface {
	iterator
	CVs() *list.ComposedValue
	With(*list.List) *list.ComposedValue
}

// A line may be parallel to another line
type line interface {
	List() *list.List
	Root() *list.Element
}

// A edge may be an Edge
type edge interface {
	Away() *list.Element
	IsRoot() bool
}
