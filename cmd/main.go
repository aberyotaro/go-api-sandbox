package main

import (
	"github.com/aberyotaro/sample_api/internal/app/user"
	"github.com/aberyotaro/sample_api/internal/infrastructure"
	"github.com/aberyotaro/sample_api/internal/usecase"
	"log"
	"net/http"
)

func main() {
	uh := user.NewHandler(usecase.NewUser(infrastructure.NewUser()))

	http.HandleFunc("/user", uh.GetUserByID)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
