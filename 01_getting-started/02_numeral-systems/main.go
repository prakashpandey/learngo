/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @github https://github.com/prakashpandey
 * @create date 2018-09-07 01:24:28
 * @modify date 2018-09-07 01:24:28
 * @desc Various types of string formating
 */

package main

import "fmt"

func main() {
	// %d: decimal
	// %b: binary
	fmt.Printf("%d - %b \n", 42, 42)
	// %x : hexadecimal
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	// hex value in caps
	fmt.Printf("%d - %b - %X \n", 42, 42, 42)
	// #x at prefix zero
	fmt.Printf("%d - %b - %#X \n", 42, 42, 42)

	// %o octal
	fmt.Printf("%d - %b - %o - %#x \n", 42, 42, 42, 42)

	// %#o octal with prefix 0
	fmt.Printf("%d - %b - %#o - %#x \n", 42, 42, 42, 42)

	// %q utf-8
	fmt.Printf("%d - %b - %#o - %#x - %q \n", 42, 42, 42, 42, 42)

}
