// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// Dancing holds the dance functions.
type Dancing struct {
	Dance func()

	OnGoal func(e *Element)
	OnFail func() *Element
	OnLeaf func()
}

// NewDancing returns a new dancing.
func NewDancing() *Dancing { // cannot pass CallBack upon New, as it is a method of himself
	return new(Dancing).init()
}

func (d *Dancing) init() *Dancing {

	d.Dance = func() { return }

	d.OnGoal = func(e *Element) { return }
	d.OnFail = func() *Element { return nil }
	d.OnLeaf = func() { return }

	return d
}
