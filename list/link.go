// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
link.go extends list.go with:

	- l.UnLeaf( d *Dancing )
	- l.ReLeaf( d *Dancing )

	- e.UnLeaf( d *Dancing )
	- e.ReLeaf( d *Dancing )

	- l.UnLink()
	- l.ReLink()

	- e.UnLink()
	- e.ReLink()

Note: For good performance, the functions are implemented using the internal data structure.
*/

package list

// ========================================================

// UnLeaf removes it temporarily.
func (l *List) UnLeaf(d *Dancing) {
	for e := l.root.next; e != &l.root; e = e.next {
		// e.UnLeaf(d) - inlined for Performance
		e.next.prev, e.prev.next = e.prev, e.next
		if e.list != nil {
			e.list.len--
		}
		if d.OnLeaf != nil {
			d.OnLeaf()
		}
	}
}

// ReLeaf restores the temporarily Unleaf'ed/removed one.
func (l *List) ReLeaf(d *Dancing) {
	for e := l.root.prev; e != &l.root; e = e.prev {
		// e.ReLeaf(d) - inlined for Performance
		e.next.prev, e.prev.next = e, e
		if e.list != nil {
			e.list.len++
		}
	}
}

// UnLeaf removes it temporarily.
func (e *Element) UnLeaf(d *Dancing) {
	// e.UnLink() - inlined for Performance
	e.next.prev, e.prev.next = e.prev, e.next
	e.list.len--
	if d.OnLeaf != nil {
		d.OnLeaf()
	}
}

// ReLeaf restores the temporarily Unleaf'ed/removed one.
func (e *Element) ReLeaf(d *Dancing) {
	// e.ReLink() - inlined for Performance
	e.next.prev, e.prev.next = e, e
	e.list.len++
}

// ========================================================

// UnLink l temporarily from it's away's
func (l *List) UnLink() {
	for e := l.root.next; e != &l.root; e = e.next {
		// e.UnLink() - inlined for Performance
		e.away.next.prev = e.away.prev
		e.away.prev.next = e.away.next
		e.away.list.len--
	}
}

// ReLink temporarily UnLink'ed l back to it's away's
func (l *List) ReLink() {
	for e := l.root.prev; e != &l.root; e = e.prev {
		// e.ReLink() - inlined for Performance
		e.away.next.prev = e.away
		e.away.prev.next = e.away
		e.away.list.len++
	}
}

// UnLink e temporarily from it's list
func (e *Element) UnLink() { e.next.prev, e.prev.next = e.prev, e.next; e.list.len-- }

// ReLink temporarily UnLink'ed e back into it's list
func (e *Element) ReLink() { e.next.prev, e.prev.next = e, e; e.list.len++ }
