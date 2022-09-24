// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

func longestprefix(val []string) string {
	longest := ""
	sort.Strings(val)
	first := string(val[0])
	last := string(val[len(val)-1])
	end := false
	for i := 0; i < len(first); i++ {

		if !end && string(first[i]) == string(last[i]) {
			longest += string(last[i])
		} else {
			end = true
		}

	}

	return longest

}

func main() {
	val := []string{"flower", "flow", "flight"}
	ret := longestprefix(val)
	fmt.Printf("\nfinal %v", ret)

}
