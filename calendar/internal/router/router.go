package router

import (
	"calendar/internal/config"
	"calendar/internal/controller"
	"calendar/internal/repository"
	"calendar/internal/usecase"
	"fmt"
	"net/http"
)

func (s *Server) Run() {
	cfg := config.Get()
	loggedMux := s.Init()
	s.logger.Printf("server started at: localhost:%s", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), loggedMux); err != nil {
		s.logger.Fatalf("failed to listen and serve: %v", err)
	}
}

func (s *Server) Init() http.Handler {

	repo := repository.New(s.logger)
	service := usecase.New(repo, s.logger)
	handler := controller.New(service, s.logger)

	s.mux.HandleFunc("/create_event", handler.Create)
	s.mux.HandleFunc("/update_event", handler.Update)
	s.mux.HandleFunc("/delete_event", handler.Delete)
	s.mux.HandleFunc("/events_for_day", handler.GetForDay)
	s.mux.HandleFunc("/events_for_week", handler.GetForWeek)
	s.mux.HandleFunc("/events_for_month", handler.GetForMonth)

	return logMiddleware(s.mux)

}
