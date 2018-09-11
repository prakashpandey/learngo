/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @github https://github.com/prakashpandey
 * @create date 2018-09-11 06:53:23
 * @modify date 2018-09-11 06:53:23
 * @desc [description]
 */

package main

import (
	"fmt"
)

// The return annonymous function closes over 'var x'
// to form a closure
func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	var increment = wrapper()
	fmt.Printf("Increment is: %d \n", increment())
	fmt.Printf("Increment again: %d \n", increment())
}
