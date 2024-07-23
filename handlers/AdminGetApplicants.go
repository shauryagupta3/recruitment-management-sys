package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shauryagupta3/recruitment-management-sys/models"
)

func (h handler) AdminGetApplicants(w http.ResponseWriter, r *http.Request) error{
	_, err := AdminProtectedHandler(w, r)
	if err != nil {
		return err
	}
	var Profiles []models.Profile

	if result := h.DB.Preload("User").Find(&Profiles); result.Error != nil {
		return result.Error
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Profiles)
	return nil
}
