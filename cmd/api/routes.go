package api

import (
	healthcheck "sarath/url_shortner/cmd/api/services"
	"sarath/url_shortner/cmd/api/services/shortner"

	"github.com/gorilla/mux"
)

func (app *Application) Routes() *mux.Router{
	mux := mux.NewRouter()

  health_check_handler := healthcheck.New(app.Logger)
  shortner_handler := shortner.New(app.Logger, app.Db, app.Cache)

  mux.HandleFunc("/health", health_check_handler.HandleHealthCheck)
  mux.HandleFunc("/shorturl", shortner_handler.ShortenURL)
  mux.HandleFunc("/get/{id}", shortner_handler.FindURL)

  return mux
}
