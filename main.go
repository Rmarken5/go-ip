package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()


	http.Handle("/", router)
	http.HandleFunc("/get-ip", ipFunc)
	fmt.Printf("Port: %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func ipFunc(writer http.ResponseWriter, request *http.Request) {
	_, err := io.WriteString(writer, request.RemoteAddr)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}
