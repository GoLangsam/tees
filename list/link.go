// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
link.go extends the (stolen and extended) list.go
with stuff, which is considered useful and helpful, such as:

	- l.UnLeaf( d *Dancing )
	- l.ReLeaf( d *Dancing )

	- e.UnLeaf( d *Dancing )
	- e.ReLeaf( d *Dancing )

	- l.UnLink()
	- l.ReLink()

	- e.UnLink()
	- e.ReLink()

Note: For good performance, the functions are implemented directly
*/

package list

// ========================================================

// Unleaf removes it temporarily.
func (l *List) UnLeaf(d *Dancing) {
	for e := l.root.next; e != &l.root; e = e.next {
		e.next.prev = e.prev
		e.prev.next = e.next
		e.list.len--
		d.OnLeaf(e)
	}
}

// Releaf restores the temporarily Unleaf'ed/removed one.
func (l *List) ReLeaf(d *Dancing) {
	for e := l.root.prev; e != &l.root; e = e.prev {
		e.next.prev = e
		e.prev.next = e
		e.list.len++
	}
}

// Unleaf removes it temporarily.
func (e *Element) UnLeaf(d *Dancing) {
	e.next.prev, e.prev.next = e.prev, e.next
	e.list.len--
	d.OnLeaf(e)
}

// Releaf restores the temporarily Unleaf'ed/removed one.
func (e *Element) ReLeaf(d *Dancing) {
	e.next.prev, e.prev.next = e, e
	e.list.len++
}

// ========================================================

// UnLink l temporarily from it's away's
func (l *List) UnLink() {
	for e := l.root.next; e != &l.root; e = e.next {
		e.away.next.prev = e.away.prev
		e.away.prev.next = e.away.next
		e.away.list.len--
	}
}

// ReLink temporarily UnLink'ed l back to it's away's
func (l *List) ReLink() {
	for e := l.root.prev; e != &l.root; e = e.prev {
		e.away.next.prev = e.away
		e.away.prev.next = e.away
		e.away.list.len++
	}
}

// UnLink e temporarily from it's list
func (e *Element) UnLink() { e.next.prev, e.prev.next = e.prev, e.next; e.list.len-- }

// ReLink temporarily UnLink'ed e back into it's list
func (e *Element) ReLink() { e.next.prev, e.prev.next = e, e; e.list.len++ }
