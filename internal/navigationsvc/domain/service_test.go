package domain

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"
)

func TestService_GetLocation(t *testing.T) {
	type request func(r GetLocationRequest) GetLocationRequest

	testCases := []struct {
		name     string
		given    request
		expected float64
	}{
		{
			name: "Success",
			given: func(r GetLocationRequest) GetLocationRequest {
				return r
			},
			expected: 1389.57,
		},
		{
			name: "CoordinateXis0",
			given: func(r GetLocationRequest) GetLocationRequest {
				r.CoordinateX = 0.0
				return r
			},
			expected: 1266.45,
		},
		{
			name: "CoordinateYis0",
			given: func(r GetLocationRequest) GetLocationRequest {
				r.CoordinateY = 0.0
				return r
			},
			expected: 933.01,
		},
		{
			name: "CoordinateZis0",
			given: func(r GetLocationRequest) GetLocationRequest {
				r.CoordinateZ = 0.0
				return r
			},
			expected: 599.68,
		},
		{
			name: "VelocityIs0",
			given: func(r GetLocationRequest) GetLocationRequest {
				r.Velocity = 0.0
				return r
			},
			expected: 1369.57,
		},
		{
			name: "CoordinateXisNegative",
			given: func(r GetLocationRequest) GetLocationRequest {
				r.CoordinateX = -1
				return r
			},
			expected: 1265.45,
		},
		{
			name: "CoordinateYisNegative",
			given: func(r GetLocationRequest) GetLocationRequest {
				r.CoordinateY = -1
				return r
			},
			expected: 932.01,
		},
		{
			name: "CoordinateZisNegative",
			given: func(r GetLocationRequest) GetLocationRequest {
				r.CoordinateZ = -1
				return r
			},
			expected: 598.68,
		},
	}

	testService := NewService(zap.NewNop(), 1)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			request := testCase.given(GetLocationRequest{
				CoordinateX: 123.12,
				CoordinateY: 456.56,
				CoordinateZ: 789.89,
				Velocity:    20.0,
			})

			actual := testService.GetLocation(context.Background(), request)
			assert.Equal(t, testCase.expected, actual)
		})

	}
}
