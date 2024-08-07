package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type GetRequest struct {
	Username string
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {

	var req GetRequest
	req.Username = chi.URLParam(r, "username")
	ctx := r.Context()
	dto, err := h.accountUsecase.Get(ctx, req.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
