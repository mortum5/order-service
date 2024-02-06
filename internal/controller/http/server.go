package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strings"

	"github.com/mortum5/order-service/internal/service"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "github.com/mortum5/order-service/docs"
)

type Server struct {
	service service.OrderService
	srv     *http.Server
	errChan chan error
}

func New(orderService service.OrderService) Server {
	return Server{
		service: orderService,
		errChan: make(chan error, 1),
	}
}

func (s *Server) Run() {
	router := http.NewServeMux()
	router.HandleFunc("/orders/", s.getHandlers())
	router.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(s.service.Config.SwaggerURL),
	))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	s.srv = srv

	log.Println("server started")
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.errChan <- fmt.Errorf("server: working %v", err)
		}
	}()
}

func (s *Server) Stop() error {
	if err := s.srv.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("server: stoping %v", err)
	}
	return nil
}

func (s *Server) Error() chan error {
	return s.errChan
}

// GetOrder godoc
// @Summary      Show an order
// @Description  get order by id
// @Produce      json
// @Param        id   path      string  true  "Object ID"
// @Success      200  {object}  entity.Order
// @Router       /orders/{id} [get].
func (s *Server) getHandlers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/orders/")

		order, err := s.service.Get(id)
		if err != nil {
			slog.Info("server:", slog.Any("err", err.Error()))
			w.WriteHeader(404)
			return
		}

		json, _ := json.Marshal(order)

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
