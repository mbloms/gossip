/*
*/

package tools

import (
	"fmt"
	"log"
	"os"
)

func AppendToFile( line string, file string ) {
	fmt.Printf( " > \"%v\" appended to %v.", line, file )
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile( file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644 )

	if err != nil {
		log.Fatal( err )
	}

	if _, err := f.Write( []byte( line + "\n" ) ); err != nil {
		log.Fatal( err )
	}

	if err := f.Close(); err != nil {
		log.Fatal( err )
	}

}
