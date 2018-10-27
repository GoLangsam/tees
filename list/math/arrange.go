// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	"sort"

	"github.com/GoLangsam/tees/list"
)

// Arrange sorts l according to it's elements TrixSizes
func Arrange(l *list.List) *list.List {

	var keys sort.IntSlice
	_, sizes := TrixSizes(l)
	for size, _ := range sizes {
		keys = append(keys, size)
	}
	keys.Sort()

	var mark *list.Element
	for _, size := range keys {
		for _, e := range sizes[size] {
			mark = e.MoveToPrevOf(mark)
		}
	}
	return l
}
