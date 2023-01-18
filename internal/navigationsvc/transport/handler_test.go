package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/betalo-sweden/navigationsvc/internal/navigationsvc/domain"
	"github.com/betalo-sweden/navigationsvc/pkg/rest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestGetLocation(t *testing.T) {
	testCases := []struct {
		name                 string
		given                getLocationRequest
		expectedResponseCode int
		expectedResponse     interface{}
	}{
		{
			name: "HappyFlow",
			given: getLocationRequest{
				CoordinateX: "123.12",
				CoordinateY: "456.56",
				CoordinateZ: "789.89",
				Velocity:    "20.1",
			},
			expectedResponseCode: http.StatusOK,
			expectedResponse: getLocationResponse{
				Loc: 1,
			},
		},
		{
			name: "Invalid Parameter",
			given: getLocationRequest{
				CoordinateX: "",
				CoordinateY: "456.56",
				CoordinateZ: "789.89",
				Velocity:    "20.1",
			},
			expectedResponseCode: http.StatusBadRequest,
			expectedResponse: rest.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "invalid_coordinate_x",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(tc.given)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/location", &buf)
			if err != nil {
				t.Fail()
			}
			require.NoError(t, err)

			h := NewHandler(zap.NewNop(), &domain.NavigationServiceMock{
				GetLocationFunc: func(ctx context.Context, request domain.GetLocationRequest) float64 {
					assert.Equal(t, tc.given.CoordinateX, fmt.Sprint(request.CoordinateX))
					assert.Equal(t, tc.given.CoordinateY, fmt.Sprint(request.CoordinateY))
					assert.Equal(t, tc.given.CoordinateZ, fmt.Sprint(request.CoordinateZ))
					assert.Equal(t, tc.given.Velocity, fmt.Sprint(request.Velocity))
					return 1
				},
			})
			handler := http.HandlerFunc(h.GetLocation)

			rr := httptest.NewRecorder()

			//Act
			handler.ServeHTTP(rr, req)

			//Assert
			assert.Equal(t, tc.expectedResponseCode, rr.Code)

			// Assert error response
			if tc.expectedResponseCode != http.StatusOK {
				var actual rest.Error

				err := json.Unmarshal(rr.Body.Bytes(), &actual)
				assert.NoError(t, err)

				assert.Equal(t, tc.expectedResponse, actual)
				return
			}

			// Assert success response
			var actual getLocationResponse

			err = json.Unmarshal(rr.Body.Bytes(), &actual)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedResponse, actual)
		})
	}
}
