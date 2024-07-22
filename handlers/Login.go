package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var SECRET = []byte("this-is-my-secret-for-jwt")

type userLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	var user userLogin
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatalln(err)
	}

	userFromDB, err := h.GetUserFromEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(err)
	}

	checkPass := CheckPasswordHash(user.Password, userFromDB.PasswordHash)
	if !checkPass {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode("password wrong")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   userFromDB.ID,
		"type": userFromDB.UserType,
		"exp":  time.Now().Add(24*time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(SECRET)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(err)
	}

	cookie := &http.Cookie{
		Name:  "token",
		Value: tokenString,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(200)
	w.Write([]byte("user loged in : " + tokenString))
}
