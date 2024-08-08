package timeline

import (
	"net/http"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	timelineUsecase usecase.Timeline
}

func NewRouter(u usecase.Timeline) http.Handler {
	r := chi.NewRouter()

	h := &handler{
		timelineUsecase: u,
	}
	r.Get("/public", h.Get)
	return r
}
