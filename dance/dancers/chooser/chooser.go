// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
choose.go defines the signature of a choosing function

*/
package chooser

import (
	"github.com/GoLangsam/tees"
)

// ========================================================

// Tees are chosen from
type Tees = tees.Tees

// Chooser is the signature of a choosing function
type Chooser func(*tees.Tees) (*tees.Tees, bool)
