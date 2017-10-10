/*
Author: Filip Johansson

Description:
Read the peers file.
RETURN: <[]string>

*/

package tools

import "strings"

func FileLinesAsSlice( file string ) []string {
	fileText := ReadTextFile( file )
	fileLinesAsslice := strings.Split( fileText, "\n" ) 								// convert content to a 'string'
	fileLinesAsslice = NoEmptyStringsInSlice( fileLinesAsslice )	// Remove all empty lines / strings
	return fileLinesAsslice
}
