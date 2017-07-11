package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubstrWithoutRepeat(t *testing.T) {
	cases := map[string]int{
		"abc":     3,
		"abca":    3,
		"abcc":    3,
		"abcabcd": 4,
		"":        0,
	}
	for key, v := range cases {
		t.Run(key, func(t *testing.T) {
			assert.Equal(t, v, LongestSubstrWithoutRepeat(key))
		})
	}
}
