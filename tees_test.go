// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

// validate interfaces

var t Tees = New("test", 1)
var e *aTee = t.Front()
var _ = new(aTee)

var _, _ dancer = t, e
var _, _ iterator = t, e
var _, _ trailer = t, e
var _, _ xrosser = t, e

var _ calcer = t // e has no fitting With
var _ line = e   // t has no List()
var _ edge = e   // t has no IsRoot()
