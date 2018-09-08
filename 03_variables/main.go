/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @github https://github.com/prakashpandey
 * @create date 2018-09-08 05:29:35
 * @modify date 2018-09-08 05:29:35
 * @desc [description]
 */

package main

import (
	"fmt"
)

func main() {
	// We can declare variables in many way in golang
	// Following are some of them

	// Empty or zero declaration
	var firstName string
	firstName = "Prakash"
	// declaration and initialization
	var middleName string = "C"
	// declartion and initialization shorthand(best way)
	lastName := "Pandey"
	fmt.Printf("%s %s %s \n", firstName, middleName, lastName)
}
