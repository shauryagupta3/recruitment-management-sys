package handlers

import (
	"net/http"
	"strconv"

	"github.com/shauryagupta3/recruitment-management-sys/models"
)

func (h handler) ApplyForJob(w http.ResponseWriter, r *http.Request) error {
	claims, err := ApplicantProtectedHandler(w, r)
	if err != nil {
		return err
	}

	query := r.URL.Query()
	jobIDStr := query.Get("job_id")

	if jobIDStr == "" {
		return NewAPIError(http.StatusBadRequest, "Missing job id")
	}

	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		return NewAPIError(http.StatusBadRequest, "Invalid job_id parameter")
	}
	var job models.Job
	if result := h.DB.First(&job, jobID); result.Error != nil {
		return result.Error
	}

	userId, err := getIDFromClaims(claims)
	if err != nil {
		return err
	}

	var user models.User
	if result := h.DB.Where("id = ?", userId).First(&user); result.Error != nil {
		return result.Error
	}
	h.DB.Model(&job).Association("Applicants").Append(&user)
	return nil
}
