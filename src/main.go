/*
Authors: Blomstarnd, Mikael & Johansson, Filip. Super cool students at KTH Stockholm.
NOTE: Its cool that we write that we are cool.
TODO:
* Get cooler
* Get shit done.
Github:  http://github.com/
*/

package main

import (
	"fmt"
	"./peers"
	"./rumors"
	"./menu"
	"./homemade_necessary_tools"
)

func main() {

	menuLoop:
	for {
		choice := menu.Display()
		switch choice {
		case "1":
			// View rumors
			rumors.View()
		case "2":
			// Add rumor
			rumor := tools.ReadInput( " > Enter rumor you wish to add: ")
			rumors.Add( rumor )
		case "3":
			// Add peer
			ip := tools.ReadInput( " > Add IP address of the peer you wish to add: " )
			peers.Add( ip )
		case "4":
			// List peers
			peers.List()
		case "q":
			break menuLoop
		default:
			fmt.Printf( " > Choose one of the given options." )
		}
	}

	//peer := peers.GetOne( p )

}
