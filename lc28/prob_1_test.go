package main

import (
	"fmt"
	"testing"
)

func TestCheckRecord(t *testing.T) {
	trueTestCases := []string{
		"PALL",
		"PALL",
	}
	for _, c := range trueTestCases {
		t.Run(c, func(t *testing.T) {
			if !checkRecord(c) {
				t.FailNow()
			}
		})
	}

	falseTestCases := []string{
		"PALAL",
		"PALLL",
	}
	for _, c := range falseTestCases {
		t.Run(c, func(t *testing.T) {
			if checkRecord(c) {
				t.FailNow()
			}
		})
	}

}

func TestCheckRecordNum(t *testing.T) {
	trueTestCases := map[int]int{
		0: 0,
		1: 3,
		2: 8,
		3: 18,
	}
	for c, d := range trueTestCases {
		t.Run(fmt.Sprintf("%d", c), func(t *testing.T) {
			o := checkRecordNum(c)
			if o != d {
				t.Errorf("expected: %d, obtained: %d", d, o)
				t.FailNow()
			}
		})
	}

}

func TestFindTilt(t *testing.T) {
	left := &TreeNode{
		Val: 2,
	}
	right := &TreeNode{
		Val: 3,
	}
	root := &TreeNode{
		Val:   1,
		Left:  left,
		Right: right,
	}
	if findTilt(nil) != 0 {
		t.Fatalf("failed nil")
		t.FailNow()
	}

	if findTilt(left) != 0 {
		t.Fatalf("failed nil")
		t.FailNow()
	}

	if findTilt(right) != 0 {
		t.Fatalf("failed right")
		t.FailNow()
	}

	if findTilt(root) != 1 {
		t.Fatalf("failed root")
		t.FailNow()
	}
}

func TestArrayPairSum(t *testing.T) {
	if arrayPairSum([]int{1, 4, 2, 3}) != 4 {
		t.Fatalf("failed nil")
		t.FailNow()
	}
}

func TestLongestLine(t *testing.T) {
	var M [][]int
	var o int

	M = [][]int{
		[]int{0, 1, 1, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 0, 0, 1},
	}
	o = longestLine(M)
	if o != 3 {
		t.Fatalf("failed nil: %d", o)
		t.FailNow()
	}

	M = [][]int{
		[]int{0, 0, 1, 1},
		[]int{0, 1, 1, 0},
		[]int{0, 1, 0, 1},
	}
	o = longestLine(M)
	if o != 3 {
		t.Fatalf("failed nil: %d", o)
		t.FailNow()
	}

	M = [][]int{
		[]int{1, 1, 1, 1},
		[]int{0, 1, 1, 0},
		[]int{0, 1, 0, 1},
	}
	o = longestLine(M)
	if o != 4 {
		t.Fatalf("failed nil: %d", o)
		t.FailNow()
	}

}
