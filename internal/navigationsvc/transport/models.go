package transport

import (
	"strconv"

	"github.com/betalo-sweden/navigationsvc/internal/navigationsvc/domain"
)

type getLocationRequest struct {
	CoordinateX string `json:"x"`
	CoordinateY string `json:"y"`
	CoordinateZ string `json:"z"`
	Velocity    string `json:"vel"`
}

type getLocationResponse struct {
	Loc float64 `json:"loc"`
}

func (r getLocationRequest) toDomain() (*domain.GetLocationRequest, error) {
	var request domain.GetLocationRequest

	x, err := strconv.ParseFloat(r.CoordinateX, 64)
	if err != nil {
		return nil, errInvalidCoordinateX
	}
	y, err := strconv.ParseFloat(r.CoordinateY, 64)
	if err != nil {
		return nil, errInvalidCoordinateY
	}
	z, err := strconv.ParseFloat(r.CoordinateZ, 64)
	if err != nil {
		return nil, errInvalidCoordinateZ
	}
	vel, err := strconv.ParseFloat(r.Velocity, 64)
	if err != nil {
		return nil, errInvalidVelocity
	}

	request.CoordinateX = x
	request.CoordinateY = y
	request.CoordinateZ = z
	request.Velocity = vel

	return &request, nil
}
