/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @github https://github.com/prakashpandey
 * @create date 2018-09-11 04:31:04
 * @modify date 2018-09-11 04:31:04
 * @desc [description]
 */

package main

import (
	"fmt"
)

func max(n int, m int) int {
	if n >= m {
		return n
	}
	return m
}

func main() {
	fmt.Println("Variable shadowing function")
	var n = 200
	var m = 400
	// max variable have shadowed max() function
	var max = max(n, m)
	fmt.Printf("Max of [%d, %d] is %d \n", n, m, max)
	// can not use max fun as it is showed by max var
	fmt.Printf("Value returned from max() function %d", max(n, m))
}
