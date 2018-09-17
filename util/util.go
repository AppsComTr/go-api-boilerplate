package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

var JWTSigningKey []byte = []byte("zxVbRh2Req.qweewq90*")

type ErrResponse struct {
	Message string `json:"message"`
}

func SendJSONResponse(w http.ResponseWriter, res interface{}, statusCode int) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}

}

func ParseRequestBody(body io.ReadCloser, v interface{}) (error, string) {

	content, err := ioutil.ReadAll(io.LimitReader(body, 1048576))
	if err != nil {
		return err, "Request body izin verilen boyutu aştı"
	}
	if err := body.Close(); err != nil {
		return err, "Request body parse edilemedi, veri doğru değil"
	}
	if err := json.Unmarshal(content, v); err != nil {
		return err, "Request body parse edilemedi, veri doğru değil"
	}

	return nil, ""

}
