/*
Auhtor: Filip Johansson

Description:
Read a textfile into a variable.
RETURN: <string>

*/

package tools

import (
	"fmt"
	"io/ioutil"
)

func ReadTextFile( file string ) string {
	text, err := ioutil.ReadFile( file )
	if err != nil {
		fmt.Print( err )
	}
	return string( text )
}
