/*
*/

package peers

import (
	"fmt"
	"log"
	"os"
)

func Add( ip string ) {
	fmt.Printf( "\n > Adding new peer with ip %v", ip )
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile( "peers/peers.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644 )

	if err != nil {
		log.Fatal( err )
	}

	if _, err := f.Write( []byte( ip + "\n" ) ); err != nil {
		log.Fatal( err )
	}

	if err := f.Close(); err != nil {
		log.Fatal( err )
	}

	fmt.Printf( "\n > Peer added [√]" )

}
