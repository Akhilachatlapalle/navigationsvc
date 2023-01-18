package transport

import (
	"errors"
	"net/http"

	"github.com/akhilachatlapalle/navigationsvc/pkg/rest"

	"github.com/davecgh/go-spew/spew"

	"go.uber.org/zap"
)

var (
	errInvalidCoordinateX = errors.New("invalid x coordinate")
	errInvalidCoordinateY = errors.New("invalid y coordinate")
	errInvalidCoordinateZ = errors.New("invalid z coordinate")
	errInvalidVelocity    = errors.New("invalid velocity")
)

func writeError(w http.ResponseWriter, logger *zap.Logger, e error) {
	err := newError(e, logger)
	rest.WriteJSON(w, err.StatusCode, err)
}

func newError(err error, logger *zap.Logger) rest.Error {
	switch err {
	case errInvalidCoordinateX:
		return rest.NewError(http.StatusBadRequest, "invalid_coordinate_x")
	case errInvalidCoordinateY:
		return rest.NewError(http.StatusBadRequest, "invalid_coordinate_y")
	case errInvalidCoordinateZ:
		return rest.NewError(http.StatusBadRequest, "invalid_coordinate_z")
	case errInvalidVelocity:
		return rest.NewError(http.StatusBadRequest, "invalid_velocity")
	}

	logger.Error("Unexpected error encountered",
		zap.String("error_dump", spew.Sdump(err)),
	)
	return rest.NewError(http.StatusInternalServerError, spew.Sdump(err))
}
