/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @github https://github.com/prakashpandey
 * @create date 2018-09-11 08:25:41
 * @modify date 2018-09-11 08:25:41
 * @desc follow this medium article for more information
 * @link https://medium.com/@pcp/blank-identifiers-in-googles-golang-3f6a10e7483a
 */

package main

import (
	"fmt"
)

func fireUpRocket() (string, bool, int) {
	var message = "Howdy from rocket! We are good and going at light speed(at least it seems like)"
	var error = true
	var timeToReachProximaCentuary = 60 // years
	return message, error, timeToReachProximaCentuary
}

func main() {
	// '_' here is blank identifier.
	// Here we don't care about the 'error' message
	// In real system you should use the error message as well
	// Why we need
	message, ok, _ := fireUpRocket()
	if ok {
		fmt.Printf("Rocket launch successfully.\n%s \n", message)
	} else {
		fmt.Printf("Rocket launch failed \n")
	}
}
