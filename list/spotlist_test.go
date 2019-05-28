// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"github.com/GoLangsam/tees/list"
)

var _ Nova = list.New()
var _ Node = list.New().Root()

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

// Coll combines all methods unique to any list, and not shared with Element
type Coll interface {
	Init() *list.List
//	Clear() *list.List

	Remove(*list.Element) interface{}

	IsEmpty() bool

	Elements() []*list.Element
	Values() list.Values

	Print(args ...interface{})

	InsertAfter(v interface{}, mark *list.Element) *list.Element
	InsertBefore(v interface{}, mark *list.Element) *list.Element

	MoveAfter(e, mark *list.Element)
	MoveBefore(e, mark *list.Element)
	MoveToBack(e *list.Element)
	MoveToFront(e *list.Element)

	PushFront(v interface{}) *list.Element
	PushBack(v interface{}) *list.Element
	PushBackList(other *list.List)
	PushFrontList(other *list.List)

	ValuesPushFront(values ...interface{})
	ValuesPushBack(values ...interface{})
}

// Atom combines all methods unique to any element, and not shared with List
type Atom interface {
	IsNode() bool
	IsRoot() bool

	List() *list.List
}
