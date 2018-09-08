/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @github https://github.com/prakashpandey
 * @create date 2018-09-08 05:50:39
 * @modify date 2018-09-08 05:50:39
 * @desc [description]
 */

package main

import (
	"fmt"
)

// Issue:
// If you run go run main.go here it will fail.
// Read link(https://stackoverflow.com/questions/21293000/go-build-works-fine-but-go-run-fails)

// How to run then?
// 1. go run main.go stars.go
// 2. go build and then ./executable_file_here

func main() {
	fmt.Printf("Our nearest star is %s \n", nameOfNearestStar)
}
