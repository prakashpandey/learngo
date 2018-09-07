/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @github https://github.com/prakashpandey
 * @create date 2018-09-07 01:49:17
 * @modify date 2018-09-07 01:49:17
 * @desc [description]
 */

package main

import "fmt"

func main() {
	fmt.Println("Decimal \t Binary \t Octal \t Hexadecimal \t\t Unicode")
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t\t %b \t\t %#o \t\t %#x \t\t %q \n", i, i, i, i, i)
	}
}
