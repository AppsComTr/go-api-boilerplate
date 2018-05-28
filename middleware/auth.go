package middleware

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/kayalardanmehmet/go-api-boilerplate/util"
)

func AuthMiddleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if len(authorization) > 0 && strings.HasPrefix(authorization, "Bearer") {

			tokens := strings.Split(authorization, " ")

			if len(tokens) == 2 {

				tokenString := tokens[1]
				//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkQXQiOjE1Mjc1MjI0NjEsIm5iZiI6MTQ0NDQ3ODQwMCwidXNlcklEIjoidGVzdCJ9.hEIaeC97l1YPv8tP3EnA2JQLUp5fNnU_yWRjajgUqec

				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
					return util.JWTSigningKey, nil
				})

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					fmt.Println(claims)
					context.Set(r, "userID", claims["userID"])
					inner.ServeHTTP(w, r)
				} else {
					fmt.Println(err)
					util.SendJSONResponse(w, util.ErrResponse{Message: "Token okunurken hata: " + err.Error()}, http.StatusUnauthorized)
				}

			} else {

				fmt.Println("Authorization eksik veya doğru değil sayi")
				util.SendJSONResponse(w, util.ErrResponse{Message: "Doğrulama yapılamadı"}, http.StatusUnauthorized)

			}

		} else {

			fmt.Println("Authorization eksik veya doğru değil")
			util.SendJSONResponse(w, util.ErrResponse{Message: "Doğrulama yapılamadı"}, http.StatusUnauthorized)

		}
	})
}
