package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shauryagupta3/recruitment-management-sys/models"
)

func (h handler) AdminGetApplicantfromID(w http.ResponseWriter, r *http.Request) error {
	_, err := AdminProtectedHandler(w, r)
	if err != nil {
		return err
	}
	id := r.PathValue("id")
	var Profile models.Profile

	if result := h.DB.Preload("User").Where("user_id = ?", id).First(&Profile); result.Error != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Profile)
	return nil
}
