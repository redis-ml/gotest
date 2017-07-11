package lc29

import (
	"fmt"
)

func maxVacationDays(flights [][]int, days [][]int) int {
	N := len(flights)
	if N == 0 {
		return 0
	}
	K := len(days[0])
	ret := 0

	table := make([][]int, K)
	for i := range table {
		table[i] = make([]int, N)
		for j := range table[i] {
			table[i][j] = -1
		}
	}
	for i := 0; i < N; i++ {
		if i != 0 && flights[0][i] == 0 {
			continue
		}
		table[0][i] = days[i][0]
		if table[0][i] > ret {
			ret = table[0][i]
		}
	}
	for week := 1; week < K; week++ {
		for city := 0; city < N; city++ {
			max := -1
			for i := 0; i < N; i++ {
				if (flights[i][city] == 0 && city != i) || table[week-1][i] == -1 {
					continue
				}
				if table[week-1][i] > max {
					max = table[week-1][i]
					table[week][city] = max + days[city][week]
					if table[week][city] > ret {
						ret = table[week][city]
					}
				}
			}
		}
	}
	return ret
}

func checkInclusion(s1 string, s2 string) bool {
	if s1 == "" {
		return true
	}
	inputLen := len(s1)
	input := 0
	chMap := make(map[int]int)
	for _, ch := range s1 {
		input = input ^ int(ch)
		chMap[int(ch)]++
	}

	hash := 0
	for i, ch := range s2 {
		code := int(ch)
		if i >= len(s1) {
			if hash == input {
				if checkSubResult(chMap, input, inputLen, i, s2) {
					return true
				}
			}
			hash = hash ^ int(s2[i-len(s1)])
		}
		hash = hash ^ code
	}
	// Check the last case.
	if hash == input {
		if checkSubResult(chMap, input, inputLen, len(s2), s2) {
			return true
		}
	}
	return false

}

func checkSubResult(chMap map[int]int, input, inputLen, i int, s2 string) bool {
	hasError := false
	tmpMap := make(map[int]int)
	for j := i - inputLen; j < i; j++ {
		tmp := int(s2[j])
		if _, ok := chMap[tmp]; !ok {
			hasError = true
			break
		}
		tmpMap[tmp]++
	}
	if !hasError && len(tmpMap) == len(chMap) {
		hasError := false
		for k, v := range tmpMap {
			if expected, ok := chMap[k]; !ok || expected != v {
				hasError = true
				break
			}
		}
		if !hasError {
			return true
		}
	}
	return false
}

func subarraySum(nums []int, k int) int {
	num := 0
	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				num++
			}
		}
	}
	return num
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	origR := len(nums)
	if origR == 0 {
		return nums
	}
	origC := len(nums[0])

	if origR*origC != r*c {
		return nums
	}

	curR := 0
	curC := 0

	ret := make([][]int, r)
	for i := range ret {
		ret[i] = make([]int, c)
		for j := range ret[i] {
			ret[i][j] = nums[curR][curC]
			curC++
			if curC >= origC {
				curC = 0
				curR++
			}
		}
	}
	return ret
}

func prob1() {
	fmt.Print("abc")
}
