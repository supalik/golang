// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func isValid(s string) bool {
	hight := 0
	for _, val := range s {
		fmt.Printf("\n val>> %v", string(val))
		if string(val) == "{" || string(val) == "(" || string(val) == "[" {
			hight++
		}
		if string(val) == "}" || string(val) == ")" || string(val) == "]" {
			if hight > 0 {
				hight--
			}
		}
	}
	if hight == 0 {
		return true
	}
	return false

}

func main() {
	val := "()[]{}"
	ret := isValid(val)
	fmt.Printf("\nfinal %v", ret)

}
