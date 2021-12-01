package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
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
	re := regexp.MustCompile("\"(.*?)\"")
	match := re.FindStringSubmatch(forwardedString)
	fmt.Println(match[1])
	_, err := io.WriteString(writer, match[1])
	if err != nil {
		log.Fatalf("error WriteString data  forwarded[\"for\"]: %v", err)
	}
}
