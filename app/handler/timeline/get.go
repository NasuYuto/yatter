package timeline

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type GetRequest struct {
	id    int
	limit int
	flag  bool
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	var req GetRequest
	queryParams := r.URL.Query()

	req.flag, _ = strconv.ParseBool(queryParams.Get("only_media"))
	req.id, _ = strconv.Atoi(queryParams.Get("since_id"))
	req.limit, _ = strconv.Atoi(queryParams.Get("limit"))

	ctx := r.Context()
	dto, err := h.timelineUsecase.Get(ctx, req.id, req.limit, req.flag)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Timeline); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
