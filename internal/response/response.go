package response

import (
	"net/http"
	"sarath/url_shortner/internal/json"
	"sarath/url_shortner/internal/json/logger"
	"sarath/url_shortner/internal/json/validator"
)
 
type ResponseWriter struct{
  logger logger.ApplicationLogger
} 

func New(logger logger.ApplicationLogger) *ResponseWriter{
  return &ResponseWriter{
    logger: logger,
  }
}

func (writer *ResponseWriter) ErrResponse(error error, w http.ResponseWriter){
  _error := json.Envelope{"error" : error.Error()}
  w.WriteHeader(http.StatusInternalServerError)
  err := json.WriteJsonToResponseWriter(_error, w)
  if err != nil{
    w.WriteHeader(http.StatusInternalServerError)
  }
}


func (writer *ResponseWriter) ValidationErrorResponse(validator *validator.Validator, w http.ResponseWriter){
  errs := json.Envelope{
    "errors" : validator.Errors,
  }
  w.WriteHeader(http.StatusBadRequest)
  writer.WriteJSONResponse(errs, w)
}

func (writer *ResponseWriter) WriteJSONResponse(data json.Envelope, w http.ResponseWriter){
  err := json.WriteJsonToResponseWriter(data , w) 
  if err != nil{
    writer.ErrResponse(err, w)
  }
}

func (writer *ResponseWriter) CreatedResponse(data json.Envelope, w http.ResponseWriter){
  w.WriteHeader(http.StatusCreated)
  err := json.WriteJsonToResponseWriter(data, w)
  if err != nil{
    writer.ErrResponse(err, w)
  }
}


func (writer *ResponseWriter) NotFoundResponse(w http.ResponseWriter){
  w.WriteHeader(http.StatusNotFound)
  err := json.WriteJsonToResponseWriter(json.Envelope{"error": "The data couldn't be found on server"}, w)
  if err != nil{
    writer.ErrResponse(err, w)
  }
}



