package healthcheck

import (
	"net/http"

	"sarath/url_shortner/internal/json"
	"sarath/url_shortner/internal/json/logger"
)

type Handler struct{
  Logger logger.ApplicationLogger
}

func New(logger logger.ApplicationLogger) *Handler{
  return &Handler{
    Logger: logger,
  }
}

func (app *Handler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
  app.Logger.Log("Got a health check request")
	data := json.Envelope{
    "status" : "OK",
  }
	json.WriteJsonToResponseWriter(data, w)
}
