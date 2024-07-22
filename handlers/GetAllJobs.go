package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shauryagupta3/recruitment-management-sys/models"
)

func (h handler) GetAllJobs(w http.ResponseWriter, r *http.Request) error {
	var jobs []models.Job

	if result := h.DB.Find(&jobs); result.Error != nil {
		return result.Error
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jobs)
	return nil
}
