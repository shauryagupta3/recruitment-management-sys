package models

import "gorm.io/gorm"

type Profile struct{
	gorm.Model
    Applicant        User   `json:"applicant" gorm:"foreignKey:ApplicantID"`
    ApplicantID      uint   `json:"applicant_id"`  // Foreign key for User
    ResumeFileAddress string `json:"resume_file_address"`
    Skills           string `json:"skills"`
    Education        string `json:"education"`
    Experience       string `json:"experience"`
    Name             string `json:"name"`
    Email            string `json:"email"`
    Phone            string `json:"phone"`
}