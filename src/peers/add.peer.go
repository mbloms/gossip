/*
Author: Filip Johansson

Descirption:
(This file might not be necessary.)
Append an ip to the list of peers.
*/

package peers

import 	"../homemade_necessary_tools"

func Add( ip string ) {
	tools.AppendToFile( ip, "peers/peers.txt" )
}
