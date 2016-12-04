package oauth2server

import (
    "fmt"
    "net/http"
    //"encoding/json"
    //"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
    
}
func cleanup(){
	fmt.Println("Closing and destroying")
}

func main() {
	defer cleanup()
	fmt.Println("Starting server")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":3080", nil)
    
}