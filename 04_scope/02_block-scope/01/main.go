/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @github https://github.com/prakashpandey
 * @create date 2018-09-11 05:51:41
 * @modify date 2018-09-11 05:51:41
 * @desc [description]
 */

package main

import (
	"fmt"
)

// package level declaration
var name = "Prakash"

func main() {
	fmt.Printf("Name: %s \n", name)
	{
		// scope within this curley braces
		var lastName = "Pandey"
	}
	// can not access var lastName outside the curley braces
	fmt.Printf("Last Name: %s \n", lastName)
}
