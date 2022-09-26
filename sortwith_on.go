// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, test")
	a := []int{0, 1, 1, 0, 1, 0, 0, 0, 1, 0} // o/p>> [0,0,0,1,1,1]
	l := len(a)
	f, r := 0, 0
	i := 0

	for i <= (len(a)/2)-1 {

		f = a[i]
		r = a[l-1]

		if f > r {

			a[i], a[l-1] = a[l-1], a[i]

		}
		l--
		if l == i+1 {
			i++
			l = len(a)

		}

	}
	fmt.Printf("\n val %v", a)

}
