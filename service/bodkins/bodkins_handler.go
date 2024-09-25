package bodkins

import (
	"fmt"
	"net/http"
	"sync"
	"webkins/service/logger"
	"webkins/service/utility/response"
	"webkins/service/utility/rw"
)

type HtmlHandler struct {
	logger.LogWrapper

	bodkins []Bodkin
	nextID  int
	mtx     sync.Mutex
}

func NewHtmlHandler() *HtmlHandler {
	return &HtmlHandler{
		LogWrapper: logger.NewLogWrapper("bodkins"),

		bodkins: make([]Bodkin, 0),
	}
}

func (hh *HtmlHandler) HandleRequest(writer response.Writer, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		hh.list(writer, r)
	case http.MethodPost:
		hh.create(writer, r)
	default:
		writer.WriteErrorResponse(http.StatusBadRequest, response.SvcErrorInvalidMethod)
	}
}

func (hh *HtmlHandler) list(writer response.Writer, _ *http.Request) {
	hh.mtx.Lock()
	defer hh.mtx.Unlock()

	if err := writer.WriteJsonResponse(http.StatusOK, hh.bodkins); err != nil {
		hh.Errorw("List", "response-write-error", fmt.Errorf("failed to write response body: %w", err))
	}
}

func (hh *HtmlHandler) create(writer response.Writer, req *http.Request) {
	var data Bodkin
	err := rw.UnmarshalJson(req.Body, &data)
	if err != nil {
		hh.Errorw("Create", "body-error", fmt.Errorf("failed to parse request body: %w", err))
		writer.WriteErrorResponse(http.StatusBadRequest, response.SvcErrorReadRequestFailed)
	}

	hh.mtx.Lock()
	defer hh.mtx.Unlock()

	// ignore id in create and set it to the "next" one
	data.Id = hh.nextID
	hh.nextID++

	hh.bodkins = append(hh.bodkins, data)
	if err := writer.WriteJsonResponse(http.StatusOK, data); err != nil {
		hh.Errorw("Create", "response-write-error", fmt.Errorf("failed to write response body: %w", err))
	}
}
