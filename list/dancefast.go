// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
dancefast.go extends list.go with:

	- l.DanceFast( d *Dancing )
	- e.DanceFast( d *Dancing )
*/

package list

// ========================================================

// DanceFast l is where the dancing begins
func (l *List) DanceFast(d *Dancing) {
	var r, ri *Element
	var xa, ya, yr  *Element
	var a, p, n *Element

	r = &l.root
	{
		{
			{
				ri = r
				a = ri.away
				p, n = a.prev, a.next
				n.prev, p.next = p, n
				a.list.len--
				for y := ri.next; y != ri; y = y.next {
					ya = y.away
					yr = &ya.list.root
					for z := ya.next; z != ya; z = z.next {
						if z.list == nil || z != yr {
							a = z.away
							p, n = a.prev, a.next
							n.prev, p.next = p, n
							a.list.len--
						}
					}
				}
			}
		}
	}
	for x := r.next; x != r; x = x.next {
		d.OnGoal(x) // Push
		xa = x.away
		for i := xa.next; i != xa; i = i.next {
			if i.list == nil || i != &xa.list.root {
				ri = &i.away.list.root
				a = ri.away
				p, n = a.prev, a.next
				n.prev, p.next = p, n
				a.list.len--
				for y := ri.next; y != ri; y = y.next {
					ya = y.away
					yr = &ya.list.root
					for z := ya.next; z != ya; z = z.next {
						if z.list == nil || z != yr {
							a = z.away
							p, n = a.prev, a.next
							n.prev, p.next = p, n
							a.list.len--
						}
					}
				}
			}
		}
		d.Dance() // Dance d is where the dancing recurs to
		for i := xa.prev; i != xa; i = i.prev {
			if i.list == nil || i != &xa.list.root {
				ri = &i.away.list.root
				for y := ri.prev; y != ri; y = y.prev {
					ya = y.away
					yr = &ya.list.root
					for z := ya.prev; z != ya; z = z.prev {
						if z.list == nil || z != yr {
							a = z.away
							p, n = a.prev, a.next
							n.prev, p.next = a, a
							a.list.len++
						}
					}
				}
				a = ri.away
				p, n = a.prev, a.next
				n.prev, p.next = a, a
				a.list.len++
			}
		}
		d.OnFail() // Pop
	}
	{
		{
			{
				ri = r
				for y := ri.prev; y != ri; y = y.prev {
					ya = y.away
					yr = &ya.list.root
					for z := ya.prev; z != ya; z = z.prev {
						if z.list == nil || z != yr {
							a = z.away
							p, n = a.prev, a.next
							n.prev, p.next = a, a
							a.list.len++
						}
					}
				}
				a = ri.away
				p, n = a.prev, a.next
				n.prev, p.next = a, a
				a.list.len++
			}
		}
	}
}

// DanceFast e is where the dancing continues.
//  Note: You must not call this directly - implemented only for symmetry, see interface Spot.
func (e *Element) DanceFast(d *Dancing) {
	panic("You cannot call this directly - implemented only for symmetry, see interface Spot")
}
