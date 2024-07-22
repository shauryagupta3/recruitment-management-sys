package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func (h handler)UploadResume(w http.ResponseWriter, r *http.Request) error{

	claims, err := ApplicantProtectedHandler(w, r)
	if err != nil {
		return err
	}
	userId, err := getIDFromClaims(claims)
	fmt.Println(userId)
	if err!=nil {
		return err
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

	destDir := "./uploads" 
    if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
        return err
    }

    // Create the file path for the destination
    filePath := filepath.Join(destDir, handler.Filename)

	// Create file
	dst, err := os.Create(filePath)
	if err!=nil {
		return err
	}
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
	return nil
}