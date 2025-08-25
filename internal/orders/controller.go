package orders

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Controller struct {
	svc Svc
}

type Svc interface {
	CreateOrder(ctx context.Context, req CreatRequest) (Response, error)
}

func NewController(svc Svc) *Controller {
	return &Controller{svc: svc}
}

func (c *Controller) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", c.CreateOrder)
	return r
}

func (c *Controller) CreateOrder(w http.ResponseWriter, r *http.Request) {
	//TODO написать контроллер и проверить его через постман
	var createOrder CreatRequest
	err := json.NewDecoder(r.Body).Decode(&createOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	order, err := c.svc.CreateOrder(r.Context(), createOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
