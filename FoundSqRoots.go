package main

import (
	"fmt"
	"math"
)

func SqRoots() {
	var k1, k2, k3 float64
	fmt.Scanln(&k1, &k2, &k3)

	d := k2*k2 - 4.0*k1*k3

	x1 := ((-k2) + math.Pow(d, 0.5)) / (2.0 * k1)
	x2 := ((-k2) - math.Pow(d, 0.5)) / (2.0 * k1)

	if math.IsNaN(x1) && math.IsNaN(x2) {
		fmt.Println(0, 0)
	} else if x1 == x2 {
		fmt.Printf("%f\n", x1)
	} else {
		fmt.Printf("%f %f\n", min(x1, x2), max(x1, x2))
	}

}
