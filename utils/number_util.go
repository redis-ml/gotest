package utils

import (
	"math"
)

func ComputePrimeArray(M int) []bool {
	arr := make([]bool, M)
	maxL := int64(len(arr)*2 + 1)
	max := int64(math.Sqrt(float64(maxL)))
	for i := int64(3); i <= max; i += 2 {
		p := InternalPrimePos(i)
		if arr[p] {
			continue
		}

		for t := i * i; t <= maxL; t += i * 2 {
			arr[InternalPrimePos(t)] = true
		}
	}
	return arr
}

func IsPrime(n int64, arr []bool) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	p := InternalPrimePos(n)
	if p < len(arr) {
		return !arr[p]
	}
	return IsPrimeByCompute(n)
}

func IsPrimeByCompute(n int64) bool {
	spec := []int64{7, 11, 13, 17, 19, 23, 29}
	if n%2 == 0 || n%3 == 0 || n%5 == 0 {
		return false
	}
	sqrt := int64(math.Sqrt(float64(n)))
	for base := int64(0); base <= sqrt; base += 30 {
		for _, d := range spec {
			if n%(base+d) == 0 {
				return false
			}
		}
	}
	return true
}

func InternalPrimePos(n int64) int {
	return (int(n) - 3) / 2
}
