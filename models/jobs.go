package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	PostedOn          time.Time `json:"posted_on"`
	TotalApplications int       `json:"total_applications"`
	CompanyName       string    `json:"company_name"`
	PostedByID        uint      `json:"posted_by_id"`
	User              User      `json:"user" gorm:"foreignKey:PostedByID"`
	Applicants        []User    `gorm:"many2many:job_applicants;"`
}
