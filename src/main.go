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
)

func main() {

	fmt.Printf( "\n > Starting rumor" )

	//peers.Add( "192.168.1.1" )
	p := peers.Read()
	fmt.Printf( "\n PEERS: %v", p )
	peer := peers.GetOne( p )
	fmt.Printf( "\n > %v", peer )
	fmt.Printf( "\n > Stopping rumor" )

}
