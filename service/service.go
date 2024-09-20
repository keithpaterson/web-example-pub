package service

import (
	"fmt"
	"net/http"
	"webkins/service/bodkins"
	"webkins/service/logger"
	"webkins/service/utility/response"
)

type Handler interface {
	HandleRequest(writer response.Writer, r *http.Request)
}

type Service struct {
	logger.LogWrapper

	port int

	// handlers
	ui      Handler
	bodkins Handler
}

func NewService(port int) *Service {
	return &Service{
		LogWrapper: logger.NewLogWrapper("server"),
		port:       port,

		ui:      NewUIHandler(),
		bodkins: bodkins.NewHtmlHandler(),
	}
}

func (s *Service) Run() error {
	http.HandleFunc("/bodkins", s.handleBodkins)
	http.HandleFunc("/", s.handleUI)

	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}

func (s *Service) handleUI(w http.ResponseWriter, r *http.Request) {
	writer := response.NewWriter(w)
	s.ui.HandleRequest(writer, r)
}

func (s *Service) handleBodkins(w http.ResponseWriter, r *http.Request) {
	writer := response.NewWriter(w)
	s.bodkins.HandleRequest(writer, r)
}
