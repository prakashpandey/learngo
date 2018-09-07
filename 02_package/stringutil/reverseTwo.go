/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @github https://github.com/prakashpandey
 * @create date 2018-09-07 04:51:12
 * @modify date 2018-09-07 04:51:12
 * @desc [description]
 */

package stringutil

func reverseSentance(w string) string {
	var s string
	for i := len(w) - 1; i >= 0; i-- {
		s += string(w[i])
	}
	return s
}
