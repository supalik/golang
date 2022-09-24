// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	
)


func romantoI(val string) int {

	switch val {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000

	}
	return 0
}

func main() {
	val := "XXVII"
	ret := 0
	for _, item := range val {
		ret = ret + romantoI(string(item))
	}
	fmt.Printf("final %v", ret)
	

}
