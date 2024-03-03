package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/priyanshupatra02/task-ninja-backend/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting the server at port 8080 ✨✨")

	log.Fatal(http.ListenAndServe(":8080", r))
}
