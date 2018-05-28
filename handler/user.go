package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kayalardanmehmet/go-api-boilerplate/util"
)

type loginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq loginRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &loginReq); err != nil {
		util.SendJSONResponse(w, util.ErrResponse{Message: "Request body parse edilemedi, veri doğru değil"}, http.StatusBadRequest)
		return
	}

	if loginReq.UserName == "" || loginReq.Password == "" {
		util.SendJSONResponse(w, util.ErrResponse{Message: "\"username\" ve \"password\" alanı boş bırakılamaz"}, http.StatusBadRequest)
		return
	}

	//TODO: db kontrolleri sağlanıp ilgili kullanıcı çekilecek
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    "test",
		"createdAt": time.Now().Unix(),
		"nbf":       time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(util.JWTSigningKey)

	if err != nil {
		panic(err)
	}

	var loginRes = loginResponse{Token: tokenString}

	util.SendJSONResponse(w, loginRes, http.StatusOK)

}
