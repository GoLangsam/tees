// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package walk

import (
	"github.com/GoLangsam/tees/list"
)

// Here is where You are - and wherever You walk to.
type Here = *list.Element

// Trail is a series of `Here` marks.
type Trail = []*list.Element
