package projecteular

import (
	"testing"

	"github.com/redisliu/gotest/utils"
	"github.com/stretchr/testify/suite"
)

func TestProjectEular(t *testing.T) { suite.Run(t, new(ProjectEularSuite)) }

type ProjectEularSuite struct {
	suite.Suite
}

func (s *ProjectEularSuite) Test57() {
	//s.T().Log("%v", verifyCylic([]int{1223, 2314, 1415}))
	pe57()
}
func (s *ProjectEularSuite) Test56() {
	N := 2
	l, r, adj := int64(1), int64(2), make(map[int64]map[int64]bool)
	s.T().Logf("%v", addEdge(l, r, adj, N))
	N = 3
	l, r = 2, 3
	s.T().Logf("%v", addEdge(l, r, adj, N))
	l, r = 2, 4
	s.T().Logf("%v", addEdge(l, r, adj, N))
	l, r = 1, 3
	s.T().Logf("%v", addEdge(l, r, adj, N))
	s.T().Logf("%v", adj)
	//pe56()
	s.T().Logf("isPrime: %d, %v", 115, utils.IsPrimeByCompute(115))
}

func (s *ProjectEularSuite) TestMain() {
	//pe54("p054_poker.txt")
	s.T().Log(sumString("47", "74"))
	pe55()
	/*
		cache := make(map[string]int)
		rt := make([]int, 10999)
		isLychrel("99409", 10677, 0, cache, rt)
	*/
}
