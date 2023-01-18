package domain

import (
	"context"

	"github.com/akhilachatlapalle/navigationsvc/pkg/math"

	"go.uber.org/zap"
)

//go:generate moq -out service_mock.go . NavigationService
type NavigationService interface {
	GetLocation(ctx context.Context, request GetLocationRequest) float64
}

type Service struct {
	Logger *zap.Logger

	SectorID int
}

func (s *Service) GetLocation(ctx context.Context, request GetLocationRequest) float64 {
	fSectorID := float64(s.SectorID)
	loc := (request.CoordinateX * fSectorID) + (request.CoordinateY * fSectorID) + (request.CoordinateZ * fSectorID) + request.Velocity
	return math.RoundFloat(loc, 2)
}

func NewService(l *zap.Logger, s int) NavigationService {
	return &Service{
		Logger:   l,
		SectorID: s,
	}
}
