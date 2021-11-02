package main

import (
	"fmt"
	"net/http"

	"github.com/LuisHenriqueFA14/introducing/internal/api"
)

func main() {
	http.HandleFunc("/", api.Full)
	http.HandleFunc("/about", api.About)
	http.HandleFunc("/personal", api.Personal)
	http.HandleFunc("/technical", api.Technical)
	http.HandleFunc("/student", api.Student)
	
	fmt.Println("Server is running!")
	http.ListenAndServe(":3000", nil)
}
