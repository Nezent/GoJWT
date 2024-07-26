package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

var users = map[string]string {
	"user1": "password1",
	"user2": "password2",
}


type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expectedPassword ,ok := users[credentials.Username]

	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute*5)

	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.Println(err)
		// w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// log.Println(tokenString)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: expirationTime,
	})


}

// func Home(w http.ResponseWriter, r *http.Request) {
	
// }

// func Refresh(w http.ResponseWriter, r *http.Request) {
	
// }