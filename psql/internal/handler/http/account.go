package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"psql/internal/usecase"
)

type TransferRequest struct {
	FromID int `json:"from_id"`
	ToID   int `json:"to_id"`
	Amount int `json:"amount"`
}

type AccountHandler struct {
	uc *usecase.AccountUseCase
}

func NewAccountHandler(uc *usecase.AccountUseCase) *AccountHandler {
	return &AccountHandler{uc: uc}
}

func (h *AccountHandler) HandleTransfer(w http.ResponseWriter, r *http.Request) {
	var req TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := AccountHandlerValidate(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.uc.Transfer(r.Context(), req.FromID, req.ToID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func AccountHandlerValidate(req *TransferRequest) error {
	if req.ToID <= 0 {
		return errors.New("to_id must be greater than zero")
	}
	if req.FromID <= 0 {
		return errors.New("from_id must be greater than zero")
	}
	if req.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if req.ToID == req.FromID {
		return errors.New("to_id and from_id must be different")
	}
	return nil
}
