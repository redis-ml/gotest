package lc32

import (
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	sorted := make([]int, len(nums))
	for i, num := range nums {
		sorted[i] = num
	}
	sort.Sort(sort.IntSlice(sorted))
	var left, right int
	for left = 0; left < len(nums); left++ {
		if sorted[left] != nums[left] {
			break
		}
	}
	if left == len(nums) {
		return 0
	}
	for right = len(nums) - 1; right > left; right-- {
		if sorted[right] != nums[right] {
			break
		}
	}
	return right - left + 1
}

func killProcess2(pid []int, ppid []int, kill int) []int {
	pidMap := make(map[int]int)
	for i, id := range pid {
		pidMap[id] = i
	}

	table := make([]int, len(pid))

	for i, id := range pid {
		if ppid[i] == kill {
			table[i] = -1
		}
		mark(table, pidMap, ppid, id, kill)
	}
	var ret []int
	for i, t := range table {
		if t < 0 {
			ret = append(ret, pid[i])
		}
	}
	return ret
}

func mark(table []int, pidMap map[int]int, ppid []int, check, kill int) int {
	if check == 0 {
		return 1
	}
	ord := pidMap[check]
	tmp := table[ord]
	if tmp != 0 {
		return tmp
	}
	if check == kill {
		table[ord] = -1
		return -1
	}
	table[ord] = mark(table, pidMap, ppid, ppid[ord], kill)
	return table[ord]
}

func killProcess(pid []int, ppid []int, kill int) []int {
	ppidMap := make(map[int][]int)
	for i, id := range ppid {
		ppidMap[id] = append(ppidMap[id], pid[i])
	}

	ret := make([]int, 0, len(pid))
	stack := make([]int, 0, len(pid))
	stack = append(stack, kill)
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		ret = append(ret, item)
		stack = append(stack[0:len(stack)-1], ppidMap[item]...)
	}
	return ret
}

func subKillProcess(ppidMap map[int][]int, kill int, ret *[]int) {
	*ret = append(*ret, kill)

	for _, newID := range ppidMap[kill] {
		subKillProcess(ppidMap, newID, ret)
	}
	return
}

func minDistance(word1 string, word2 string) int {
	return len(word1) + len(word2) - lcs(word1, word2)*2
}

func lcs(word1 string, word2 string) int {
	if len(word1) == 0 {
		return 0
	}
	if len(word2) == 0 {
		return 0
	}

	table := make([][]int, len(word1))
	for i := range table {
		table[i] = make([]int, len(word2))
	}

	for i := range word2 {
		found := false
		for j := 0; j <= i; j++ {
			if word2[j] == word1[0] {
				found = true
				break
			}
		}
		if found {
			table[0][i] = 1
		}
	}
	for i := range word1 {
		found := false
		for j := 0; j <= i; j++ {
			if word1[j] == word2[0] {
				found = true
				break
			}
		}
		if found {
			table[i][0] = 1
		}
	}
	for i := 1; i < len(word1); i++ {
		for j := 1; j < len(word2); j++ {
			if word2[j] == word1[i] {
				table[i][j] = table[i-1][j-1] + 1
				continue
			}
			tmp := table[i][j-1]
			if tmp < table[i-1][j] {
				tmp = table[i-1][j]
			}
			table[i][j] = tmp
		}
	}
	return table[len(word1)-1][len(word2)-1]
}
