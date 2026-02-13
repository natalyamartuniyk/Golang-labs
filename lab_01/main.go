//package main
//
//import (
//	"fmt"
//	"math"
//	"math/rand"
//)
//
//func main() {
//
//	var y float64
//	x := float64(rand.Int31n(100))
//
//	switch {
//	case x < 50:
//		y = math.Pow(x, 2) + 13 - 5*x
//	default:
//		y = 8 - math.Pow(x, 2) - math.Pow(x, 3)
//	}
//
//	fmt.Println("x=", x)
//	fmt.Println("y=", y)
//
//}

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	var y float64
	x := float64(rand.Int31n(100))

	if x < 50 {
		y = math.Pow(x, 2) + 13 - 5*x
	} else {
		y = 8 - math.Pow(x, 2) - math.Pow(x, 3)

	}

	fmt.Println("x=", x)
	fmt.Println("y=", y)

}
