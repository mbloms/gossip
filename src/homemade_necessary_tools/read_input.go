/*
Author: Filip Johansson

Description:
Read input from command line.
RETURN: <string>
*/

package tools

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func ReadInput( msg string ) string {
	fmt.Printf( msg )
	reader := bufio.NewReader( os.Stdin )
	input, _ := reader.ReadString( '\n' )
	input = strings.Replace( input, "\n", "", -1 )
	return input
}
