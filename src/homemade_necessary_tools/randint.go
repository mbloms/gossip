/*
Author: Filip Johansson
*/

package tools

import (
	"math/rand"
	"time"
)

func RandInt ( min int, max int ) int {
	rand.Seed( time.Now().UTC().UnixNano() )
	return min + rand.Intn( max-min )
}
