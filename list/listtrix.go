// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
listtrix.go extends the (stolen and extended) list.go
with stuff, which is considered useful and helpfull, such as:

	- l.AddBeam
	- l.AddList

	- l.AddJunk
	- l.AddOnes

	- l.Size

*/
package list

// ===========================================================================

// AddBeam returns a new list with root-value v and new elements with values dots,
// the root of which is Junk'ed to a new PushBack-Element (with same value v) of l
func (l *List) AddBeam(v interface{}, vals ...interface{}) *List {

	var list = NewList(v, vals...)
	return l.AddList(list)
}

// AddList adds a new element with value v to list l
// Junk'es it with the list's root
// and returns list (not l) - useful for calculated lists
func (l *List) AddList(list *List) *List {

	var head = l.PushBack(l.With(list))
	list.root.Junk(head) //	:SameAs head.Junk( list.Root() )

	return list
}

// ===========================================================================

// AddJunk adds Junk'ed pairs of new elements to list l across respective lists
// The new elements carry pointers to both Roots() as ComposedValue.
// Thus, they carry their coordinates, so to say.
func (l *List) AddJunk(v interface{}, lists ...*List) *List {

	var list = NewList(v)

	for _, head := range lists {
		vert := head.PushBack(head.With(list))
		hori := list.PushBack(list.With(head))
		vert.Junk(hori)
	}
	return l.AddList(list)
}

// AddOnes adds Junk'ed pairs of new elements to list l across respective lists
// The new elements carry no Values. They are light, so to say.
func (l *List) AddOnes(v interface{}, lists ...*List) *List {

	var list = NewList(v)

	for _, head := range lists {
		vert := head.PushBack(nil)
		hori := list.PushBack(nil)
		vert.Junk(hori)
	}
	return l.AddList(list)
}

// Size returns the sum of Len() of the AwayLists along e.List
func (e *Element) Size() int {
	return e.List().Size()
}

// Size returns the sum of Len() of the AwayLists along l
func (l *List) Size() int {
	var c int
	for e := l.Front(); e != nil; e = e.Next() {
		c += e.AwayList().Len()
	}
	return c
}
