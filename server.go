package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	//"encoding/json"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	port := "3000"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	fmt.Println("Starting server Oauth2 Server on " + port)
	http.ListenAndServe(":"+port, router)

}
