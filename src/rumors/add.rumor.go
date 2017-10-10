/*
Author: Filip Johansson

Descirption:
(This file might not be necessary.)
Append a rumor to the list of rumors.
*/

package rumors

import (
	"../homemade_necessary_tools"
	"fmt"
)

func Add( rumor string ) {
	// Check for redundancy
	// If the rumor already exists it should not
	// be added to the list of rumors
	rumors := tools.FileLinesAsSlice( "rumors/rumors.txt" )
	redundant := false

	for _, r := range rumors {
		if rumor == r {
			redundant = true
		}
	}

	if !redundant {
			tools.AppendToFile( rumor, "rumors/rumors.txt" )
	} else {
		fmt.Printf( " > Rumor already exists." )
	}

}
