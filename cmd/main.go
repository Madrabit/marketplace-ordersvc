package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/madrabit/marketplace-ordersvc/internal/adapters/pricing"
	"github.com/madrabit/marketplace-ordersvc/internal/orders"
	"github.com/madrabit/marketplace-ordersvc/internal/web"
	"net/http"
)

func main() {
	repo := orders.NewRepository()
	pricer := pricing.NewPricingController("http://localhost:8081")
	service := orders.NewService(repo, pricer)
	controller := orders.NewController(service)
	server := web.NewServer()
	server.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Mount("/orders", controller.Routes())
		})
	})
	addr := ":8080" // слушать все интерфейсы
	fmt.Println("listening on", addr)
	err := http.ListenAndServe(addr, server.Router)
	if err != nil {
		fmt.Println("server cant start")
		return
	}
}
