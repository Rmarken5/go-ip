package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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
	log.Printf("header: %v\n", request.Header)
	forwardedString := request.Header.Get("Forwarded")
	log.Println(forwardedString)
	ipAddr := strings.TrimRight(strings.TrimLeft(forwardedString, "\""), "\"")
	log.Println(ipAddr)
	_, err := io.WriteString(writer, ipAddr)
	if err != nil {
		log.Fatalf("error WriteString data  forwarded[\"for\"]: %v", err)
	}
}
