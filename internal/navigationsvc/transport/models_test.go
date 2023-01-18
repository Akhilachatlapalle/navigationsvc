package transport

import (
	"testing"

	"github.com/betalo-sweden/navigationsvc/internal/navigationsvc/domain"

	"github.com/stretchr/testify/assert"
)

func TestGetLocationRequestToDomainFailure(t *testing.T) {
	type request func(r getLocationRequest) getLocationRequest

	testCases := []struct {
		name        string
		given       request
		expectedErr error
	}{
		{
			name: "InvalidCoordinateX",
			given: func(r getLocationRequest) getLocationRequest {
				r.CoordinateX = ""
				return r
			},
			expectedErr: errInvalidCoordinateX,
		},
		{
			name: "InvalidCoordinateY",
			given: func(r getLocationRequest) getLocationRequest {
				r.CoordinateY = ""
				return r
			},
			expectedErr: errInvalidCoordinateY,
		},
		{
			name: "InvalidCoordinateZ",
			given: func(r getLocationRequest) getLocationRequest {
				r.CoordinateZ = ""
				return r
			},
			expectedErr: errInvalidCoordinateZ,
		},
		{
			name: "InvalidVelocity",
			given: func(r getLocationRequest) getLocationRequest {
				r.Velocity = ""
				return r
			},
			expectedErr: errInvalidVelocity,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			request := testCase.given(getLocationRequest{
				CoordinateX: "123.12",
				CoordinateY: "456.56",
				CoordinateZ: "789.89",
				Velocity:    "20.0",
			})

			_, err := request.toDomain()
			assert.Error(t, err)
			assert.EqualError(t, testCase.expectedErr, err.Error())
		})
	}
}

func TestGetLocationRequestToDomainSuccess(t *testing.T) {
	type request func(r getLocationRequest) getLocationRequest

	testCases := []struct {
		name     string
		given    request
		expected domain.GetLocationRequest
	}{
		{
			name: "HappyFlow",
			given: func(r getLocationRequest) getLocationRequest {
				return r
			},
			expected: domain.GetLocationRequest{
				CoordinateX: 1.1,
				CoordinateY: 2.2,
				CoordinateZ: 3.3,
				Velocity:    4.4,
			},
		},
		{
			name: "CoordinateXIsNegative",
			given: func(r getLocationRequest) getLocationRequest {
				r.CoordinateX = "-1.1"
				return r
			},
			expected: domain.GetLocationRequest{
				CoordinateX: -1.1,
				CoordinateY: 2.2,
				CoordinateZ: 3.3,
				Velocity:    4.4,
			},
		},
		{
			name: "CoordinateXIsZero",
			given: func(r getLocationRequest) getLocationRequest {
				r.CoordinateX = "0"
				return r
			},
			expected: domain.GetLocationRequest{
				CoordinateX: 0,
				CoordinateY: 2.2,
				CoordinateZ: 3.3,
				Velocity:    4.4,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			request := testCase.given(getLocationRequest{
				CoordinateX: "1.1",
				CoordinateY: "2.2",
				CoordinateZ: "3.3",
				Velocity:    "4.4",
			})

			actual, err := request.toDomain()
			assert.NoError(t, err)
			assert.Equal(t, &testCase.expected, actual)
		})
	}
}
