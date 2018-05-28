package util

import (
	"encoding/json"
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
