// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

import (
	"github.com/GoLangsam/tees/list"
)

// New returns a new *list.List
func New(v interface{}, vals ...interface{}) *list.List {
	return list.NewList(v, vals...)
}
