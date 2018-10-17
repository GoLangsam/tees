// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"github.com/GoLangsam/tees/list"
)

// ===========================================================================

// Board returns a new fully populated N*M matrix
func Board(N, M int) *list.List {
	var r = list.NewList("Ranks")
	var f = list.NewList("Files")

	for _, id := range IDs("R-", N) {
		r.AddBeam(id)
	}
	var rc = r.AwayLists()
	for _, id := range IDs("F-", M) {
		f.AddJunk(id, rc...)
	}

	r.Junk(f)
	return r
}
