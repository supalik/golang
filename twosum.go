// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	//"sync"
)

//func twoSum(nums []int, target int) []int {
//
//}

func main() {
	nums := []int{3, 1, 4, 5, 2}
	target := 6
	rem := 0
	var retarr []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			rem = target - nums[i]
			if rem == nums[j] {
				retarr = append(retarr, i)
				retarr = append(retarr, j)

			}
		}
	}
	fmt.Printf("ret arr %v", retarr)
	/***/
	var retarr1 []int
	var rem1 = 0
	mymap := make(map[int]int)
	for idx, num := range nums {
		rem1 = target - num
		if j, ok := mymap[rem1]; ok {
			retarr1 = append(retarr1, j)
			retarr1 = append(retarr1, idx)

		}
		mymap[num] = idx

	}
	fmt.Printf("\n ret arr %v", retarr1)

}
