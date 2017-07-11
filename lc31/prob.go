package lc31

import (
	"fmt"
)

func distributeCandies(candies []int) int {
	m := make(map[int]int)
	for _, c := range candies {
		m[c]++
	}
	if len(m) >= len(candies)/2 {
		return len(candies) / 2
	}
	return len(m)
}

func minDistance(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	distanceSum := 0
	distances := make([]int, len(nuts))
	for i, nut := range nuts {
		distances[i] = distance(nut, tree) * 2
		distanceSum += distances[i]
	}

	// First, calulate the first nut to get.
	minDist := -1
	for i, nut := range nuts {
		dist := distance(squirrel, nut) +
			distance(nut, tree) +
			distanceSum - distances[i]
		if minDist == -1 || minDist > dist {
			minDist = dist
		}
	}
	return minDist
}

func distance(s, d []int) int {
	h := s[0] - d[0]
	if h < 0 {
		h = -h
	}
	w := s[1] - d[1]
	if w < 0 {
		return h - w
	}
	return h + w
}

func minDistance2(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	m := make([]int, height*width)
	for _, nut := range nuts {
		m[toIndex(nut[0], nut[1], width)] = 1
	}
	m[toIndex(tree[0], tree[1], width)] = 1

	distanceSum := 0
	distances := make([]int, len(nuts))
	for i, nut := range nuts {
		distances[i] = distance(nut, tree) * 2
		distanceSum += distances[i]
	}

	// First, calulate the first nut to get.
	minDist := -1
	for i, nut := range nuts {
		dist := minDistanceForPoint(squirrel, nut, m, height, width) +
			minDistanceForPoint(nut, tree, m, height, width) +
			distanceSum - distances[i]
		if minDist == -1 || minDist > dist {
			minDist = dist
		}
	}
	return minDist
}

func minDistanceForPoint(src, dst []int, m []int, height, width int) int {
	mark := make([]int, len(m))
	dstIndex := toIndex(dst[0], dst[1], width)

	q := make([][]int, 0, 1000)
	q = append(q, []int{src[0], src[1], 0})
	for len(q) > 0 {
		item := q[0]
		q = q[1:]
		// Up
		if item[0] > 0 {
			idx := toIndex(item[0]-1, item[1], width)
			if m[idx] == 0 && mark[idx] == 0 {
				q = append(q, []int{item[0] - 1, item[1], item[2] + 1})
				mark[idx] = 2
			} else if idx == dstIndex {
				return item[2] + 1
			}
		}
		// Down
		if item[0] < height-1 {
			idx := toIndex(item[0]+1, item[1], width)
			if m[idx] == 0 && mark[idx] == 0 {
				q = append(q, []int{item[0] + 1, item[1], item[2] + 1})
				mark[idx] = 2
			} else if idx == dstIndex {
				return item[2] + 1
			}
		}
		// right
		if item[1] > 0 {
			idx := toIndex(item[0], item[1]-1, width)
			if m[idx] == 0 && mark[idx] == 0 {
				q = append(q, []int{item[0], item[1] - 1, item[2] + 1})
				mark[idx] = 2
			} else if idx == dstIndex {
				return item[2] + 1
			}
		}
		// Left
		if item[1] < width-1 {
			idx := toIndex(item[0], item[1]+1, width)
			if m[idx] == 0 && mark[idx] == 0 {
				q = append(q, []int{item[0], item[1] + 1, item[2] + 1})
				mark[idx] = 2
			} else if idx == dstIndex {
				return item[2] + 1
			}
		}
	}
	panic("bay")
	return 0
}

func toIndex(row, col, width int) int {
	return row*width + col
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	return isSubtreeWithEarlyExit(s, t, false)
}

func isSubtreeWithEarlyExit(s *TreeNode, t *TreeNode, withEarlyExit bool) bool {
	if s == nil {
		if t == nil {
			return true
		}
		return false
	}
	if t == nil {
		return false
	}
	if s.Val == t.Val {
		if isSubtreeWithEarlyExit(s.Left, t.Left, true) &&
			isSubtreeWithEarlyExit(s.Right, t.Right, true) {
			return true
		}
	}
	if withEarlyExit {
		return false
	}
	return isSubtreeWithEarlyExit(s.Left, t, false) ||
		isSubtreeWithEarlyExit(s.Right, t, false)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func run() {
	fmt.Printf("%s\n", "test")
}
