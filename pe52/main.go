package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	for n := 100; ; n++ {
		totalDigits, digits := analysis(n)
		tmp := n
		found := true
		for i := 2; i < 7; i++ {
			tmp += n
			newTotalDigits, newDigits := analysis(tmp)
			if newTotalDigits != totalDigits {
				// skip later data
				n = int(math.Pow10(totalDigits))
				found = false
				break
			}
			if !reflect.DeepEqual(newDigits, digits) {
				found = false
				break
			}
		}
		if found {
			fmt.Printf("got data: %d\n", n)
			break
		}
	}
}

func analysis(n int) (totalDigits int, digits [10]int) {
	for ; n != 0; n /= 10 {
		d := n % 10
		digits[d]++
		totalDigits++
	}
	return
}
