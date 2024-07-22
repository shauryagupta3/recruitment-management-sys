package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shauryagupta3/recruitment-management-sys/models"
)

func (h handler) AdminGetJobFromID(w http.ResponseWriter, r *http.Request) error{
	_, err := AdminProtectedHandler(w, r)
	if err != nil {
		return err
	}
	id := r.PathValue("id")
	var jobs models.Job

	if result := h.DB.Where("postedbyid = ?", id).First(&jobs); result.Error != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jobs)
	return nil
}
