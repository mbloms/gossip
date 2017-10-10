/*

Description:
Get a random IP from the list of IP addresses

*/

package peers

import (
	"fmt"
	"io/ioutil"
	"strings"
	"../homemade_necessary_tools"
)

func Read() []string {

	fmt.Printf( "\n > Reading peers." )
	ips, err := ioutil.ReadFile( "peers/peers.txt" ) // just pass the file name
	if err != nil {
		fmt.Print( err )
	}

	peers := strings.Split( string( ips ), "\n" ) 				// convert content to a 'string'
	peers = tools.NoEmptyStringsInSlice( peers )					// Remove all empty lines / strings
	//fmt.Printf( "\n %v, peers ) // print the content as a 'string'
	fmt.Printf( "\n > All peers read." )
	return peers



}
