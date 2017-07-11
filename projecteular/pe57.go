package projecteular

import (
	"fmt"
)

func PE57() {
	pe57()
}
func pe57() {
	startBase := []int{45, 32, 26, 23, 21, 19}
	endBase := []int{141, 100, 82, 71, 64, 59}
	flBase := []func(int) int{
		Triangle, Square, Pentagonal, Hexagonal, Heptagonal, Octagonal,
	}
	dflBase := []func(int) int{
		DeltaTriangle, DeltaSquare, DeltaPentagonal, DeltaHexagonal, DeltaHeptagonal, DeltaOctagonal,
	}
	N := 6
	l := make([]int, N)
	start := make([]int, N)
	end := make([]int, N)
	fl, dfl := make([]func(int) int, N), make([]func(int) int, N)
	for i := range l {
		start[N-1-i] = startBase[i]
		end[N-1-i] = endBase[i]
		fl[N-1-i] = flBase[i]
		dfl[N-1-i] = dflBase[i]
	}
	ret := trav(0, l, start, end, fl, dfl)
	fmt.Printf("result: %v found: %v\n", l, ret)
}

func trav(n int, l []int, start, end []int, fl, dfl []func(int) int) bool {
	if n == len(l) {
		return verifyCylic(l)
	}

	for i := start[n]; i < end[n]; i++ {
		if i == start[n] {
			l[n] = fl[n](i)
		} else {
			l[n] += dfl[n](i)
		}
		if containsZero(l[n]) {
			continue
		}
		if trav(n+1, l, start, end, fl, dfl) {
			return true
		}
	}
	return false
}

func verifyCylic(l []int) bool {
	head := make(map[int]int)
	tail := make(map[int]int)
	for _, d := range l {
		head[d/100]++
		tail[d%100]++
	}
	if len(head) != len(tail) {
		return false
	}
	for k, v := range head {
		if tail[k] != v {
			return false
		}
	}
	// exact check
	headArr := make([]map[int]bool, 100)
	tailArr := make([]map[int]bool, 100)
	visited := make([]bool, len(l))
	for i := range headArr {
		headArr[i] = make(map[int]bool)
		tailArr[i] = make(map[int]bool)
	}
	for i, d := range l {
		headArr[d/100][i] = true
		tailArr[d%100][i] = true
	}

	return verifyCylicUtil(0, l[0], 0, visited, headArr, tailArr, l)
}

func verifyCylicUtil(n, num, id int, visited []bool, head, tail []map[int]bool, l []int) bool {
	if n == len(l) {
		return num == l[0]
	}
	if visited[id] {
		return false
	}
	visited[id] = true
	defer func() { visited[id] = false }()

	t := num % 100
	m := head[t]
	for nID := range m {
		if verifyCylicUtil(n+1, l[nID], nID, visited, head, tail, l) {
			return true
		}
	}
	return false

}

func containsZero(n int) bool {
	return n/1000 == 0 || (n%100)/10 == 0
}

func Triangle(n int) int        { return n * (n + 1) / 2 }
func DeltaTriangle(n int) int   { return n }
func Square(n int) int          { return n * n }
func DeltaSquare(n int) int     { return n*2 - 1 }
func Pentagonal(n int) int      { return n * (3*n - 1) / 2 }
func DeltaPentagonal(n int) int { return n*3 - 2 }
func Hexagonal(n int) int       { return n * (2*n - 1) }
func DeltaHexagonal(n int) int  { return n*4 - 3 }
func Heptagonal(n int) int      { return n * (5*n - 3) / 2 }
func DeltaHeptagonal(n int) int { return n*5 - 4 }
func Octagonal(n int) int       { return n * (3*n - 2) }
func DeltaOctagonal(n int) int  { return n*6 - 5 }
