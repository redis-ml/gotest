package projecteular

import (
	"fmt"
	"strconv"
)

func pe55() {
	rt := make([]int, 10000)
	cache := make(map[string]int)
	count := 0
	for i := len(rt) - 1; i >= 0; i-- {
		if isLychrel(strconv.Itoa(i), i, 0, cache, rt) {

		} else {
			//fmt.Printf("sss: %d\n", i)
			count++
		}
	}
	fmt.Printf("result: %d\n", count)
}

func isLychrel(n string, num, iter int, cache map[string]int, rt []int) bool {
	if iter == 53 {
		rt[num] = 1
		return false
	}

	if id, ok := cache[n]; ok && rt[id] != 0 {
		if rt[id] == 1 {
			//fmt.Printf("cached non-result: n: %s, num: %d, iter: %d, id: %d\n", n, num, iter, id)
			return false
		}
		return true
	}
	r := reverseString(n)
	if id, ok := cache[r]; ok && rt[id] != 0 {
		if rt[id] == 1 {
			//fmt.Printf("cached non-result: n: %s, num: %d, iter: %d, id: %d\n", n, num, iter, id)
			return false
		}
		return true
	}

	cache[n] = num
	cache[r] = num
	sum := sumString(n, r)
	realSum := reverseString(sum)
	//fmt.Printf("n: %s, r: %s, sum: %s\n", n, r, realSum)
	if isPalindrom(sum) {
		cache[sum] = num
		rt[num] = 2
		return true
	}
	return isLychrel(realSum, num, iter+1, cache, rt)
}

func isPalindrom(in string) bool {
	for l, r := 0, len(in)-1; l < r; l++ {
		if in[l] != in[r] {
			return false
		}
		r--
	}
	return true
}

const dict = "0123456789"

var dictMap map[int]rune

func init() {
	dictMap := make(map[int]rune)
	for i, v := range dict {
		dictMap[i] = v
	}
}

func sumString(in, r string) string {
	rt := make([]rune, 0, len(in)+1)
	e := 0
	for i, v := range in {
		s := int(v-'0') + int(r[i]-'0') + e
		//fmt.Printf("[%v]\n", s)
		rt = append(rt, rune(s%10)+'0')
		e = s / 10
	}
	if e != 0 {
		rt = append(rt, '1')
	}
	return string(rt)
}

func reverseString(in string) string {
	r := make([]rune, len(in))
	for i, v := range in {
		r[len(r)-1-i] = v
	}
	return string(r)
}
