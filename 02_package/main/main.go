/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @github https://github.com/prakashpandey
 * @create date 2018-09-07 04:27:02
 * @modify date 2018-09-07 04:27:02
 * @desc A program to learn about packages
 */

package main

import (
	"fmt"

	"github.com/prakashpandey/learngo/02_package/stringutil"
)

func main() {
	fmt.Printf("My name is %s \n", stringutil.Name)
	fmt.Printf("Name in reverse format is %s \n", stringutil.Reverse(stringutil.Name))
}
