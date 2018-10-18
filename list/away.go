// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
away.go extends the (stolen and extended) list.go
with stuff, which is considered useful and helpful, such as:

	- e.Junk (x *Element)	// cross-links Away's
	- l.Junk (x *List)	// cross-links Away's of Root() of two lists

	- l.Away()		*Element
	- e.Away()		*Element
	- l.AwayList()		*List
	- e.AwayList()		*List
	- l.AwayNext()		*Element
	- e.AwayNext()		*Element
	- l.AwayPrev()		*Element
	- e.AwayPrev()		*Element
	- l.Home()		*Element
	- e.Home()		*Element

	- l.Junks( *List )	bool
	- e.Junks( *Element )	bool

	- l.IsSolo()		bool
	- e.IsSolo()		bool

	- e.IsJunk()		bool

*/

package list

// ===========================================================================
// func (e *Element) ...
// Move

// Away returns the away element (an element of the orthogonal fiber-list)
func (l *List) Away() *Element {
	return l.root.Away()
}

// Away returns the away element (an element of the orthogonal fiber-list)
func (e *Element) Away() *Element {
	if e == nil {
		return nil
	}
	return e.away
}

// AwayList returns next of the away element (an element of the orthogonal fiber-list)
func (l *List) AwayList() *List {
	return l.root.AwayList()
}

// AwayList returns list of the away element (the orthogonal fiber-list)
func (e *Element) AwayList() *List {
	if e == nil || e.away == nil {
		return nil
	}
	return e.away.List()
}

// AwayNext returns next of the away element (an element of the orthogonal fiber-list)
func (l *List) AwayNext() *Element {
	return l.root.AwayNext()
}

// AwayNext returns next of the away element (an element of the orthogonal fiber-list)
func (e *Element) AwayNext() *Element {
	if e == nil || e.away == nil {
		return nil
	}
	return e.away.Next()
}

// AwayPrev returns prev of the away element (an element of the orthogonal fiber-list)
func (l *List) AwayPrev() *Element {
	return l.root.AwayPrev()
}

// AwayPrev returns prev of the away element (an element of the orthogonal fiber-list)
func (e *Element) AwayPrev() *Element {
	if e == nil || e.away == nil {
		return nil
	}
	return e.away.Prev()
}

// Home returns the away element of the root of the list
func (l *List) Home() *Element {
	if l == nil {
		return nil
	}
	return l.root.Away()
}

// Home returns the away element of the root of the list of e
func (e *Element) Home() *Element {
	if e == nil {
		return nil
	}
	if &e.list == nil {
		return nil
	}
	return e.list.Home()
}

// ===========================================================================

// SetAway sets the away (an element of the orthogonal fiber-list)
// If e == mark or e and mark are element of the same list, the element is not modified.
// NOTE: only any new List's list-Element points to itself!
/* Is this needed by the outer world?
func (e *Element) SetAway(away *Element) {
	if e == away || e.list == away.list {
		return
	}
	e.away = away
}
*/

// ===========================================================================
// Binary func(...)

// Junk links another lists Root() as a mutual Junc-tion of their Away's
func (l *List) Junk(x *List) {
	l.Root().Junk(x.Root())
}

// Junk links another element as a mutual Junc-tion of their Away's
func (e *Element) Junk(x *Element) {
	if e == nil {
		panic("e.Junk: e is nil!")
	}
	if x == nil {
		panic("e.Junk: x is nil!")
	}
	if e.away != nil {
		panic("Junk: away in e already set!")
	}
	if x.away != nil {
		panic("Junk: away in x already set!")
	}
	e.away = x
	x.away = e
}

// ===========================================================================
// Binary => bool

// Junks reports whether the roots of l and t are Junk'ed
func (l *List) Junks(t *List) bool {
	return (l.root.away == &t.root) && (t.root.away == &l.root)
	//	return l.Root().Junks(t.Root()) is less fast
}

// Junks reports whether the elements e and i are Junk'ed
func (e *Element) Junks(i *Element) bool {
	return (e.away == i) && (i.away == e)
}

// ===========================================================================
// => bool

// IsSolo reports whether the list l has no Away()
func (l *List) IsSolo() bool {
	//	if l.root == nil { panic("IsSolo: Root is nil") }
	return l.root.away == nil
	//	return l.Root().IsSolo() is less fast
}

// IsSolo reports whether the element e has no Away()
func (e *Element) IsSolo() bool {
	return e.away == nil
}

// IsJunk reports whether the lists root has it's Away() Junk'ed
func (l *List) IsJunk() bool {
	//	if l.root == nil { panic("IsJunk: Root is nil") }
	return (l.root.away != nil) && (&l.root == l.root.away.away)
	//	return l.Root().IsJunk() is less fast
}

// IsJunk reports whether the element e has it's Away() Junk'ed
func (e *Element) IsJunk() bool {
	return (e.away != nil) && (e == e.away.away)
}
