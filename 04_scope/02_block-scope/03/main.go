/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @github https://github.com/prakashpandey
 * @create date 2018-09-11 06:23:39
 * @modify date 2018-09-11 06:23:39
 * @desc [description]
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Closure demonstration")
	var x = 0
	var increment = func() {
		x++
	}

	fmt.Print("Increment x: ")
	increment()
	fmt.Println(x)

	fmt.Print("Increment x again: ")
	increment()
	fmt.Println(x)
}
