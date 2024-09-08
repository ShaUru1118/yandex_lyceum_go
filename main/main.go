package main

import "math"

func CalculateSeriesSum(n int) float64 {
	s := 0.0
	for i := 1; i <= n; i++ {
		s += 1 / float64(i)
	}
	return s
}

func FindIntersection(k1, b1, k2, b2 float64) (float64, float64) {
	if k1 == k2 {
		return math.NaN(), math.NaN()
	}
	x := (b2 - b1) / (k1 - k2)
	y := k1*x + b1

	return x, y
}

func CalculateDigitalRoot(n int) int {
	sumOfDigits := func(x int) int {
		sum := 0
		for x > 0 {
			sum += x % 10
			x /= 10
		}
		return sum
	}

	for n >= 10 {
		n = sumOfDigits(n)
	}
	return n
}
