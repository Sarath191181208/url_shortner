package shortner

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"sarath/url_shortner/internal/cache"
	"sarath/url_shortner/internal/data"
	"sarath/url_shortner/internal/json"
	"sarath/url_shortner/internal/json/logger"
	"sarath/url_shortner/internal/json/validator"
	"sarath/url_shortner/internal/response"

	"github.com/gorilla/mux"
)

type Handler struct {
	Logger logger.ApplicationLogger
	db     *data.Models
	cache  cache.Cache
}

func New(logger logger.ApplicationLogger, db *data.Models, cache cache.Cache) *Handler {
	return &Handler{
		Logger: logger,
		db:     db,
		cache:  cache,
	}
}

func (app *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Url    string `json:"url"`
		Newurl string `json:"new_url"`
	}
	writer := response.New(app.Logger)

	// read the json
	err := json.ReadJsonFromReq(&input, w, r)
	if err != nil {
		app.Logger.Log(fmt.Sprint(err))
		writer.ErrResponse(err, w)
		return
	}

	// validate the json
	v := validator.New()
	v.Check(len(input.Newurl) > 3, "new_url", "The shortened url name must be greater than 8")
	v.Check(len(input.Newurl) < 32, "new_url", "The shortened url name must be less than 32")
	v.Check(v.Matches(input.Newurl, validator.CharsDigitsRX), "new_url", "Only characters and digits are allowed")

	if !v.Valid() {
		writer.ValidationErrorResponse(v, w)
		return
	}

	// Try inserting into the db
	url := data.Url{
		OriginalUrl:  input.Url,
		ShortenedUrl: input.Newurl,
	}
	app.Logger.Log(fmt.Sprint("Original: ", url.OriginalUrl, " Shortened: ", url.ShortenedUrl))
	err = app.db.Urls.Insert(&url)
	if err != nil {
		writer.ErrResponse(err, w)
		app.Logger.Log(err.Error())
		return
	}

	// return the created response
	writer.CreatedResponse(json.Envelope{
		"data": url,
	}, w)
}

func (app *Handler) FindURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	writer := response.New(app.Logger)

	if id == "" {
		writer.ErrResponse(errors.New("path can't be empty"), w)
		return
	}

	v := validator.New()
	v.Check(len(id) > 4 && len(id) < 32, id, "invalid id length")
	v.Check(v.Matches(id, validator.CharsDigitsRX), id, "invalid id string")

	if !v.Valid() {
		writer.ValidationErrorResponse(v, w)
		return
	}

	url := &data.Url{
		ShortenedUrl: id,
	}

	// check in cache
	res, err := app.cache.Get(id)
	if err != nil {
		app.Logger.Log(fmt.Sprint("Cache get failed due to errror: ", err.Error()))
	}
	// send the response
	if err == nil && res != "" {
		app.Logger.Log("Fetched from cache")
		json.ReadJSONString(&url, strings.NewReader(res))
		writer.WriteJSONResponse(json.Envelope{"data": url}, w)
		return
	}

	app.Logger.Log("Fetched the url from the db")
	err = app.db.Urls.Find(url)
	if err != nil {
		writer.ErrResponse(err, w)
		return
	}

	// save in cache
	jsonStr, err := json.ToJSONString(url)
	if err != nil {
		app.Logger.Log(fmt.Sprint("Can't convert the response to json due to the following error: ", err.Error()))
	} else {
		app.Logger.Log("Saved the url into the cache")
		app.cache.Set(id, jsonStr, time.Minute*5)
	}

	writer.WriteJSONResponse(json.Envelope{
		"data": url,
	}, w)
}
