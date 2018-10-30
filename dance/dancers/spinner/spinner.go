// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
package spinner provides some Dance- and Turn-functions which
keep a dancer busy dancing a dance on sth in search for goals, or turning back
*/
package spinner

import (
	"github.com/GoLangsam/tees"
	"github.com/GoLangsam/tees/dance/dancing"
	"github.com/GoLangsam/tees/dance/dancing/turn"
)

// ========================================================

// Tees are danced upon
type Tees = tees.Tees

// Spinner is the signature of a dance function
type Spinner func(*dancing.Dancing, *tees.Tees)

// Turner is the signature of a turn function
type Turner func(*turn.Dancing, *tees.Tees)
