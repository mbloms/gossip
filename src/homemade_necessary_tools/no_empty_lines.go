/*
Author: Filip Johansson

Description:
Remove empty strings from slice with strings.
*/

package tools

import (

)

func NoEmptyStringsInSlice ( s []string ) []string {
	var betterSlice []string
	for _, str := range s {
		if str != "" {
			betterSlice = append( betterSlice, str )
		}
	}
	return betterSlice
}
