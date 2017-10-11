/*
Author: Filip Johansson

*/

package peers

import (
	"fmt"
	"../homemade_necessary_tools"
)

func List() {
	peers := tools.FileLinesAsSlice( "peers/peers.txt" )
	l := len( peers )

	// Take number of rumors to be read e.g. all or 10.
	fmt.Printf( " > %v peers available. \n > List (y/n)?: ", l )
	choice := tools.ReadInput( "" )
	fmt.Printf( " | -------------------------\n" )
	if choice == "y" {
		// Display the rumors from bottom up,
		// since the newest are appended at the end of the file.
		for _, peer := range peers {
			fmt.Printf( "\n\t %v \n", peer )
		}
	}


}
