package transport

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/betalo-sweden/navigationsvc/internal/navigationsvc/domain"
	"github.com/betalo-sweden/navigationsvc/pkg/rest"
)

type Handler struct {
	Service domain.NavigationService
	Logger  *zap.Logger
}

func NewHandler(l *zap.Logger, s domain.NavigationService) Handler {
	return Handler{
		Service: s,
		Logger:  l,
	}
}

func (h *Handler) GetLocation(w http.ResponseWriter, r *http.Request) {
	var req getLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, h.Logger, err)
		return
	}

	reqDomain, err := req.toDomain()
	if err != nil {
		writeError(w, h.Logger, err)
		return
	}

	location := h.Service.GetLocation(r.Context(), *reqDomain)
	rest.WriteJSON(w, http.StatusOK, getLocationResponse{Loc: location})
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	rest.WriteJSON(w, http.StatusOK, struct{}{})
}
