package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shauryagupta3/recruitment-management-sys/models"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	pass := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func (h handler) Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatalln(err)
	}

	if user.UserType!="admin" && user.UserType!="applicant" {
		log.Fatal("wrong type passed")
	}

	user.PasswordHash = hashPassword(user.PasswordHash)
	if result := h.DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
