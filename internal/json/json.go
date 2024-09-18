package json

import (
	"encoding/json"
	"io"
	"net/http"
)

type Envelope map[string]any

func ToJSONString(data any) (string, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(d[:]), nil
}

func ReadJSONString(data interface {}, reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&data)
	if err != nil {
		return err
	}

	return nil
}

func WriteJsonToResponseWriter(data Envelope, w http.ResponseWriter) error {
	// convert to json data
	json_data, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// write the json data
	w.Header().Add("Content-Type", "application/json")
	w.Write(json_data)

	return nil
}

func ReadJsonFromReq(data interface{}, w http.ResponseWriter, r *http.Request) error {
  return ReadJSONString(data, r.Body)
}
