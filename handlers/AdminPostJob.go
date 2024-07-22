package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/shauryagupta3/recruitment-management-sys/models"
)

func (h handler) PostJob(w http.ResponseWriter, r *http.Request) error {

	claims, err := AdminProtectedHandler(w, r)
	if err != nil {
		return err
	}

	userId, _ := getIDFromClaims(claims)

	var job models.Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		return err
	}
	u64, err := strconv.ParseUint(userId, 10, 64)
	job.PostedByID = uint(u64)

	if result := h.DB.Create(&job); result.Error != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Job Created")
	return nil
}
