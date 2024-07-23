package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/shauryagupta3/recruitment-management-sys/models"
)

func (h handler) UploadResume(w http.ResponseWriter, r *http.Request) error {

	claims, err := ApplicantProtectedHandler(w, r)
	if err != nil {
		return err
	}
	userId, err := getIDFromClaims(claims)
	fmt.Println(userId)
	if err != nil {
		return err
	}
	var User models.User
	if result := h.DB.First(&User, userId); result.Error != nil {
		return result.Error
	}

	// Maximum upload of 2 MB files
	r.ParseMultipartForm(2 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return err
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	destDir := "/home/shaurya/code/git/recruitment-management-sys/uploads/"
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return err
	}
	uniqueName := strconv.FormatUint(uint64(User.ID), 10) + filepath.Ext(handler.Filename)
	filePath := filepath.Join(destDir, uniqueName)
	fmt.Println(filePath)
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")

	file, err = os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Create a buffer with the file content
	fileBuffer := bytes.NewBuffer(fileContent)

	apiURL := os.Getenv("API_URL")
	apiKey := os.Getenv("API_KEY")

	// Create the request
	req, err := http.NewRequest("POST", apiURL, fileBuffer)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("apikey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return err
	}

	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", responseBody)

	type Education struct {
		Name string `json:"name"`
	}

	type ResponseProfile struct {
		Skills     []string    `json:"skills"`
		Education  []Education `json:"education"`
		Experience []string    `json:"experience"`
		Name       string      `json:"name"`
		Email      string      `json:"email"`
		Phone      string      `json:"phone"`
	}
	var respProfile ResponseProfile

	if err := json.Unmarshal(responseBody, &respProfile); err != nil {
		fmt.Println(err)
		return err
	}

	education := ""
	for _, edu := range respProfile.Education {
		if education == "" {
			education = edu.Name
		} else {
			education = education + ", " + edu.Name
		}
	}
	profile := models.Profile{
		Skills:     strings.Join(respProfile.Skills, ", "),
		Experience: strings.Join(respProfile.Experience, ", "),
		Name:       respProfile.Name,
		Email:      respProfile.Email,
		Phone:      respProfile.Phone,
	}

	profile.Education = education
	profile.User = User
	profile.ResumeFileAddress = filePath

	if err := h.DB.Create(&profile).Error; err != nil {
		return err
	}

	return nil
}
