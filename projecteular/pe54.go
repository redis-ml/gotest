package projecteular

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type cardType int

const (
	// HighCard is card type.
	HighCard cardType = iota
	OnePair
	TwoPairs
	ThreeOfKind
	Straight
	Flush
	FullHouse
	FourOfKind
	StraightFlush
)

const suiteList = "23456789TJQKA"

var pokerMap map[rune]int

func init() {
	pokerMap = make(map[rune]int)
	for i, r := range suiteList {
		pokerMap[r] = i
	}
}

func pe54(fn string) {
	fo, err := os.Open(fn)
	if err != nil {
		panic("failed to open file")
	}
	defer fo.Close()

	sum, lines := 0, 0
	scan := bufio.NewScanner(fo)
	for scan.Scan() {
		line := scan.Text()
		s := strings.Split(line, " ")
		p1 := reverse(judge(s[0:5]))
		p2 := reverse(judge(s[5:]))
		fmt.Printf("input: %v %v, %v\n", s, p1, p2)
		sum += comp(p1, p2)
		lines++
	}
	fmt.Printf("count: %d lines: %d\n", sum, lines)
}

func comp(l, r []int) (ret int) {
	for i, v := range l {
		if v > r[i] {
			return 1
		}
		if v < r[i] {
			return 0
		}
	}
	return 0
}

func reverse(in []int) []int {
	o := make([]int, len(in))
	for i, v := range in {
		o[len(in)-1-i] = v
	}
	return o
}

func judge(cards []string) []int {
	num := make(map[int]int)
	suite := make(map[int]int)
	for _, str := range cards {
		//fmt.Printf("out: %s\n", str)
		for i, v := range str {
			if i == 0 {
				//fmt.Printf("d: %v, v: %v ", v, pokerMap[v])
				num[pokerMap[v]]++
			} else {
				suite[int(v)]++
			}
		}
	}
	var t cardType
	if len(suite) == 1 {
		ks := keyset(num)
		sort.Sort(sort.IntSlice(ks))
		t = Flush
		if isStraight(ks) {
			t = StraightFlush
		}
		ks = append(ks, int(t))
		return ks
	}
	//fmt.Printf("suite len: %d, nums: %d\n", len(suite), len(num))
	switch len(num) {
	case 2:
		var large, small int
		t = FullHouse
		for k, v := range num {
			switch v {
			case 1:
				t = FourOfKind
				small = k
			case 2:
				small = k
			case 3:
				large = k
			case 4:
				t = FourOfKind
				large = k
			}
		}
		return []int{small, large, int(t)}
	case 3:
		var ks, pair []int
		var tok int
		t = ThreeOfKind
		for k, v := range num {
			switch v {
			case 1:
				ks = append(ks, k)
			case 2:
				t = TwoPairs
				pair = append(pair, k)
			case 3:
				t = ThreeOfKind
				tok = k
			}
		}
		switch t {
		case ThreeOfKind:
			sort.Sort(sort.IntSlice(ks))
			ks = append(ks, tok)
		case TwoPairs:
			sort.Sort(sort.IntSlice(pair))
			ks = append(ks, pair...)
		}
		return append(ks, int(t))
	case 4:
		var extra int
		var ks []int
		for k, v := range num {
			if v == 2 {
				extra = k
			} else {
				ks = append(ks, k)
			}
		}
		sort.Sort(sort.IntSlice(ks))
		ks = append(ks, extra, int(OnePair))
		return ks
	case 5:
		ks := keyset(num)
		sort.Sort(sort.IntSlice(ks))
		if isStraight(ks) {
			t = Straight
		} else {
			t = HighCard
		}
		ks = append(ks, int(t))
		return ks
	}

	return nil
}
func isStraight(ks []int) bool {
	prev := 0
	for i, v := range ks {
		if i > 0 && v != prev+1 {
			return false
		}
		prev = v
	}
	return true
}
func keyset(m map[int]int) (ret []int) {
	for k := range m {
		ret = append(ret, k)
	}
	return
}
