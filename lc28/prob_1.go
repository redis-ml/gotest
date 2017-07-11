package main

import (
	"fmt"
	"sort"
)

type record struct {
	horizontal, vertical, diagonal int
}

func longestLine(M [][]int) int {
	m := len(M)
	if m == 0 {
		return 0
	}
	n := len(M[0])
	table := make([][]record, m)
	for i := range table {
		table[i] = make([]record, n)
	}

	table[0][0] = record{
		horizontal: M[0][0],
		vertical:   M[0][0],
		diagonal:   M[0][0],
	}

	for i := 1; i < n; i++ {
		if M[0][i] == 1 {
			table[0][i].horizontal = table[0][i-1].horizontal + 1
			table[0][i].vertical = 1
			table[0][i].diagonal = 1
		}
	}
	for j := 1; j < m; j++ {
		if M[j][0] == 1 {
			table[j][0].vertical = table[j-1][0].vertical + 1
			table[j][0].horizontal = 1
			table[j][0].diagonal = 1
		}
	}
	max := 0
	for j := 1; j < m; j++ {
		for i := 1; i < n; i++ {
			if M[j][i] == 1 {
				table[j][i].diagonal = table[j-1][i-1].diagonal + 1
				table[j][i].horizontal = table[j][i-1].horizontal + 1
				table[j][i].vertical = table[j-1][i].vertical + 1
			}
		}
	}
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			if max < table[j][i].diagonal {
				max = table[j][i].diagonal
			}
			if max < table[j][i].horizontal {
				max = table[j][i].horizontal
			}
			if max < table[j][i].vertical {
				max = table[j][i].vertical
			}
		}
	}

	// Reuse table
	table[0][n-1] = record{
		diagonal: M[0][n-1],
	}

	for j := 1; j < m; j++ {
		table[j][n-1].diagonal = M[j][n-1]
	}
	for j := 1; j < m; j++ {
		for i := n - 2; i >= 0; i-- {
			if M[j][i] == 1 {
				table[j][i].diagonal = table[j-1][i+1].diagonal + 1
				if max < table[j][i].diagonal {
					max = table[j][i].diagonal
				}
			} else {
				table[j][i].diagonal = 0
			}
		}
	}
	fmt.Printf("%v", table)
	return max
}

func arrayPairSum(nums []int) int {
	sort.Sort(sort.IntSlice(nums))
	var sum int
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	_, ret := traverseNode(root)
	return ret
}

func traverseNode(root *TreeNode) (sum, tiltSum int) {
	if root == nil {
		return 0, 0
	}
	leftSum, leftTilt := traverseNode(root.Left)
	rightSum, rightTilt := traverseNode(root.Right)
	tmp := leftSum - rightSum
	if tmp < 0 {
		tmp = -tmp
	}
	return leftSum + rightSum + root.Val, tmp + leftTilt + rightTilt
}

func checkRecordNum(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 3
	}

	m := int64(1000000007)

	num := prepareTable(n, m)

	// cases: 0 A, and A* and *A
	sum := (num[n-1] + (num[n-2] << 1)) % m

	for i := 0; i < n-2; i++ {
		sum = (sum + (num[i] * num[n-2-i])) % m
	}
	return int(sum)
}

func prepareTable(n int, m int64) []int64 {
	t1 := make([]int64, n)
	t2 := make([]int64, n)
	t3 := make([]int64, n)
	t4 := make([]int64, n)
	num := make([]int64, n)
	t1[0], t3[0] = 1, 1
	num[0] = t1[0] + t3[0]
	t1[1], t2[1], t3[1], t4[1] = 1, 1, 1, 1
	num[1] = 4

	for i := 2; i < n; i++ {
		bc := t2[i-1]%m + t3[i-1]%m
		ad := t1[i-1]%m + t4[i-1]%m
		t1[i] = bc
		t2[i] = bc
		t3[i] = ad
		t4[i] = ad
		num[i] = ((bc + ad) << 1) % m
	}
	return num
}

func checkRecord(s string) bool {
	lastLPos := -1
	numOfA := 0
	for i, ch := range s {
		switch ch {
		case 'A':
			lastLPos = -1
			numOfA++
			if numOfA > 1 {
				return false
			}
		case 'L':
			if lastLPos >= 0 {
				if i-lastLPos > 1 {
					return false
				}
			} else {
				lastLPos = i
			}
		default:
			lastLPos = -1
		}
	}
	return true
}

func main() {

}
