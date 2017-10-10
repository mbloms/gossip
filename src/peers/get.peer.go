/*
Author: Filip Johansson

Description:
Get a random peer from the list of peers.
*/

package peers

import 	"../homemade_necessary_tools"

func GetOne( peerSlice []string ) string {
	// Get length of list with peers
	l := len( peerSlice )
	index := tools.RandInt( 0, l )
	peer := peerSlice[ index ]

	return peer

}
