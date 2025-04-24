package http

import (
	"encoding/json"
	"net/http"
	"psql/internal/usecase"
)

type ReserveRequest struct {
	RoomID int  `json:"room_id"`
	Set    bool `json:"set"`
}

type RoomHandler struct {
	uc *usecase.RoomUseCase
}

func NewRoomHandler(uc *usecase.RoomUseCase) *RoomHandler {
	return &RoomHandler{uc: uc}
}

func (h *RoomHandler) Reserve(w http.ResponseWriter, r *http.Request) {
	var req ReserveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Set {
		err := h.uc.Reserve(r.Context(), req.RoomID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		err := h.uc.UnsetReserve(r.Context(), req.RoomID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
