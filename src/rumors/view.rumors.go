/*
Author: Filip Johansson

*/

package rumors

import (
	"fmt"
	"strconv"
	"../homemade_necessary_tools"
)

func View() {
	rumors := tools.FileLinesAsSlice( "rumors/rumors.txt" )

	l := len( rumors )

	// Take number of rumors to be read e.g. all or 10.
	fmt.Printf( " > %v rumor(s) available. \n > How many rumors do you wish to view?\n > (type \"all\" for all rumors): ", l )
	choice := tools.ReadInput( "" )
	fmt.Printf( "| -------------------------\n" )
	n := 0
	if choice != "all" {
		m, err := strconv.Atoi( choice )
		if err != nil {
			fmt.Printf( "\n > %v",err )
			fmt.Printf( "\n > Enter an integer or \"all\"" )
			View() // NOTE: This solution is lazy. It is recursive and reads the file many times. It could be improved in many ways...
		}
		// If a bigger number than max is put in,
		// reduce it to the maximum available.
		if m >= l {
			n = l
		} else {
			n = m
		}
	} else {
		n = l
	}

	// Display the rumors from bottom up,
	// since the newest are appended at the end of the file.
	for i := l-1; i >= l-n; i-- {
		fmt.Printf( "\n\t %v \n", rumors[ i ] )
	}

}
