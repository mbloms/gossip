/*

Description:
Display a menu with options.
Return the choice.
*/

package menu

import (
	"fmt"
	"../homemade_necessary_tools"
)

func Display() string {
	fmt.Printf( "\n | -------------------------" )
	fmt.Printf( "\n | RUMOR\n | =====" )
	fmt.Printf( "\n | [1] View rumors" )
	fmt.Printf( "\n | [2] Add rumor" )
	fmt.Printf( "\n | [3] Add peer" )
	fmt.Printf( "\n | [4] List peer" )
	fmt.Printf( "\n | [q] Quit" )
	fmt.Printf( "\n | -------------------------" )
	choice := tools.ReadInput( "\n > " )
	return choice
}
