// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
listmany.go extends the (stolen and extended) list.go
with stuff, which is considered useful and helpful, such as:

	- l.Elements()		[]*Element

	- l.Values()		Values

	- l.ValuesPushBack( v... )
	- l.ValuesPushFront( v... )

*/

package list

// ===========================================================================

// Values aliases a slice of Element.Value
type Values []interface{}

// ===========================================================================

// Elements returns the elements of list l as a slice
func (l *List) Elements() []*Element {
	var data = make([]*Element, 0, l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		data = append(data, e)
	}
	return data
}

// Values returns all Element.Values as Values-slice
func (l *List) Values() Values {
	var data = make([]interface{}, 0, l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		data = append(data, e.Value)
	}
	return data
}

// ===========================================================================
// func (l *List) ...

// ValuesPushBack appends a slice of Values
func (l *List) ValuesPushBack(values ...interface{}) {
	valuesDo(l.PushBack, values...)
}

// ValuesPushFront prepends a slice of Values
func (l *List) ValuesPushFront(values ...interface{}) {
	valuesDo(l.PushFront, values...)
}

// valuesDo executes the given function on a slice of Values
func valuesDo(do func(v interface{}) *Element, values ...interface{}) {
	for i := range values {
		do(values[i])
	}
}
