package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/shauryagupta3/recruitment-management-sys/db"
	"github.com/shauryagupta3/recruitment-management-sys/handlers"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("recieved request : " + id))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db_url := os.Getenv("POSTGRES")
	postgres := db.Init(db_url)
	h := handlers.New(postgres)

	router := http.NewServeMux()

	router.HandleFunc("POST /signup", h.Signup)
	router.HandleFunc("POST /login", h.Login)

	router.HandleFunc("POST /uploadresume", handlers.Make(h.UploadResume))

	router.HandleFunc("POST /admin/job", handlers.Make(h.PostJob))
	router.HandleFunc("GET /admin/job/{id}", handlers.Make(h.AdminGetJobFromID))
	router.HandleFunc("GET /admin/applicants", handlers.Make(h.AdminGetApplicants))
	router.HandleFunc("GET /admin/applicant/{id}", handlers.Make(h.AdminGetApplicantfromID))

	router.HandleFunc("GET /jobs", handlers.Make(h.GetAllJobs))
	router.HandleFunc("GET /jobs/apply", handlers.Make(h.ApplyForJob))

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	log.Println("server starting at :3000")
	server.ListenAndServe()
}
