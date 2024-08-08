package statuses

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type GetRequest struct {
	id int
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {

	var req GetRequest
	req.id, _ = strconv.Atoi(chi.URLParam(r, "id"))
	ctx := r.Context()
	dto, err := h.statusUsecase.FindStatusByID(ctx, req.id)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
