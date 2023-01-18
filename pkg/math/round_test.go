package math

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundFloat(t *testing.T) {
	cases := []struct {
		given     float64
		precision uint
		expected  float64
	}{
		{
			given:     1397.565,
			precision: 2,
			expected:  1397.57,
		},
		{
			given:     1397.5700001,
			precision: 2,
			expected:  1397.57,
		},
		{
			given:     1397.5700001,
			precision: 3,
			expected:  1397.57,
		},
		{
			given:     1397,
			precision: 3,
			expected:  1397,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("RoundFloat %v_%v", tc.given, tc.precision), func(t *testing.T) {
			actual := RoundFloat(tc.given, tc.precision)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
