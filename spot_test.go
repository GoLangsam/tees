// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees_test

import (
	"github.com/GoLangsam/tees"
)

func test(s tees.Spot) {}

func ExampleSpot() {
	var l = tees.New("Test", 1, "2", "drei", false)
	test(l)
	var e = l.PushBack(l)
	test(e)

}
