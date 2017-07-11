package lc32

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestLC32(t *testing.T) { suite.Run(t, new(LC32Suite)) }

type LC32Suite struct {
	suite.Suite
}

func (s *LC32Suite) TestTest() {
	s.T().Log("")
}

func (s *LC32Suite) TestLCS() {
	s.Equal(1, lcs("ab", "a"))
	s.Equal(1, lcs("a", "ab"))

	s.Equal(2, lcs("eat", "sat"))
	s.Equal(2, lcs("sea", "sed"))

	s.Equal(3, lcs("1a2b3c4", "abc"))
	s.Equal(2, lcs("sea", "sed"))

}

func (s *LC32Suite) TestMinDistance() {
	s.Equal(2, minDistance("sea", "set"))
	s.Equal(0, minDistance("", ""))

	s.Equal(2, minDistance("ab", ""))

	s.Equal(2, minDistance("", "ab"))

}

func (s *LC32Suite) TestKillProcess() {
	var pid []int
	var ppid []int

	pid = []int{1, 3, 10, 5}
	ppid = []int{3, 0, 5, 3}
	s.T().Logf("%v", killProcess(pid, ppid, 5))

	s.T().Logf("%v", killProcess(pid, ppid, 0))

	pid = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ppid = []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	N := 100000
	pid = make([]int, N)
	for i := range pid {
		pid[i] = i + 1
	}
	ppid = make([]int, N)
	for j := range ppid {
		if j == 0 {
			ppid[j] = 0
		} else {
			ppid[j] = 1
		}
	}
	killProcess(pid, ppid, 1)
	s.T().Logf("%v", 1)
}

func (s *LC32Suite) Test223() {
	s.Equal(5, findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
}
