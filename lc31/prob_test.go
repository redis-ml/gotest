package lc31

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestLC31(t *testing.T) { suite.Run(t, new(LC31Suite)) }

type LC31Suite struct {
	suite.Suite
}

func (st *LC31Suite) TestDistributeCandies() {
	st.Equal(3, distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	st.Equal(2, distributeCandies([]int{1, 1, 2, 3}))
}

func (st *LC31Suite) TestMinDistance() {
	height := 5
	width := 7
	tree := []int{2, 2}
	squirrel := []int{4, 4}
	nuts := [][]int{
		[]int{3, 0},
		[]int{2, 5},
	}
	st.Equal(12, minDistance(height, width, tree, squirrel, nuts))

	height = 1
	width = 3
	tree = []int{0, 1}
	squirrel = []int{0, 0}
	nuts = [][]int{
		[]int{0, 2},
	}
	st.Equal(3, minDistance(height, width, tree, squirrel, nuts))

}

func (st *LC31Suite) TestMinDistancePoint() {
	st.Equal(
		6,
		minDistanceForPoint(
			[]int{2, 2},
			[]int{0, 2},
			[]int{0, 0, 1, 0, 1, 1, 0, 0, 0},
			3, 3,
		))
}

func (st *LC31Suite) TestDistance() {
	st.Equal(2, distance([]int{2, 2}, []int{1, 1}))
}

func (st *LC31Suite) TestTest() {
	st.T().Log("hehe")
}

func (st *LC31Suite) TestIsSubTree() {
	var s, t *TreeNode

	s = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	t = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	st.True(isSubtree(s, t))

	s = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 0,
				},
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	t = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	st.False(isSubtree(s, t))
}
