/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @github https://github.com/prakashpandey
 * @create date 2018-09-11 06:08:59
 * @modify date 2018-09-11 06:08:59
 * @desc [description]
 */

package main

import (
	"fmt"
)

var x = 0

func increment() int {
	x++
	return x
}

func main() {
	fmt.Printf("Increment x %d \n", increment())
	fmt.Printf("Increment x %d \n", increment())
}
