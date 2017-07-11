package lc29

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestProb1(t *testing.T) {
	suite.Run(t, new(Prob1Suite))
}

type Prob1Suite struct {
	suite.Suite
}

func (s *Prob1Suite) TestMaxVacationDays() {
	var flights, days [][]int

	flights = [][]int{
		[]int{0, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 0},
	}
	days = [][]int{
		[]int{1, 3, 1},
		[]int{6, 0, 3},
		[]int{3, 3, 3},
	}
	s.Equal(12, maxVacationDays(flights, days))

	flights = [][]int{
		[]int{0, 0, 0},
		[]int{0, 0, 0},
		[]int{0, 0, 0},
	}
	days = [][]int{
		[]int{1, 1, 1},
		[]int{7, 7, 7},
		[]int{7, 7, 7},
	}
	s.Equal(3, maxVacationDays(flights, days))

	flights = [][]int{
		[]int{0, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 0},
	}
	days = [][]int{
		[]int{7, 0, 0},
		[]int{0, 7, 0},
		[]int{0, 0, 7},
	}
	s.Equal(21, maxVacationDays(flights, days))

	flights = [][]int{
		[]int{0, 1, 0},
		[]int{0, 0, 0},
		[]int{0, 0, 0},
	}
	days = [][]int{
		[]int{0, 0, 7},
		[]int{2, 0, 0},
		[]int{7, 7, 7},
	}
	s.Equal(7, maxVacationDays(flights, days))
}

func (s *Prob1Suite) TestCheckInclusion() {
	s.True(checkInclusion("ab", "ab"))
	s.True(checkInclusion("ab", "eidba000"))
	s.False(checkInclusion("ab", "eidbooa000"))
}

func (s *Prob1Suite) TestSubarraySum() {
	s.Equal(2, subarraySum([]int{1, 1, 1}, 2))
}

func (s *Prob1Suite) TestMatrixReshape() {
	var input, output [][]int

	input = [][]int{
		[]int{
			1, 2,
		},
		[]int{
			3, 4,
		},
	}
	output = matrixReshape(input, 1, 4)
	s.T().Logf("%v", output)

	input = [][]int{
		[]int{
			1, 2,
		},
		[]int{
			3, 4,
		},
	}
	output = matrixReshape(input, 2, 4)
	s.T().Logf("%v", output)

	input = [][]int{
		[]int{},
		[]int{},
	}
	output = matrixReshape(input, 2, 4)
	s.T().Logf("%v", output)
}
