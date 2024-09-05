package ui

import (
	"errors"
	"net/http"
	"os"
	"webkins/service/logger"
	"webkins/service/utility/response"
)

var (
	ErrNilServer = errors.New("server was nil")
)

type HtmlHandler struct {
	logger.LogWrapper

	baseDir string
	handler http.Handler
}

func NewHtmlHandler() *HtmlHandler {
	return &HtmlHandler{
		LogWrapper: logger.NewLogWrapper("html"),
		baseDir:    lookupEnv("WEBKINS_UI_PATH", "html/"),
		handler:    nil,
	}
}
func (hh *HtmlHandler) HandleRequest(writer response.Writer, r *http.Request) {
	hh.Infow("HandleRequest", "BaseDir", hh.baseDir)
	if hh.handler == nil {
		if _, err := os.Stat(hh.baseDir); err != nil {
			hh.Infow("Fetch", "Stat BaseDir", hh.baseDir, "Error", err.Error())
			writer.WriteResponse(http.StatusInternalServerError)
			return
		}

		hh.handler = http.StripPrefix("", http.FileServer(http.Dir(hh.baseDir)))
	}
	hh.handler.ServeHTTP(writer.HttpResponseWriter(), r)
}

func lookupEnv(key string, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}
