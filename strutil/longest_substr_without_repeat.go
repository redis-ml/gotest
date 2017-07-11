package strutil

// LongestSubstrWithoutRepeat calculates the length of the longtest substring of given str,
// which doesn't have any repeated chars.
func LongestSubstrWithoutRepeat(str string) int {
	posTable := make([]int, 128)
	for i := range posTable {
		posTable[i] = -1
	}

	maxStart := 0
	maxLen := 0
	curLen := 0
	curStart := 0

	for i, ch := range str {
		prevPos := posTable[int(ch)]
		if prevPos >= 0 {
			// exists
			curStart = prevPos + 1
			for j := prevPos - 1; j >= maxStart; j-- {
				posTable[int(str[j])] = -1
			}
		} else {
			curLen++
			posTable[int(ch)] = i
			if curLen > maxLen {
				maxStart = curStart
				maxLen = curLen
			}
		}

	}
	return maxLen
}
