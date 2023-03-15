package basics

import (
	"fmt"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!!")
	fmt.Println("Home Page")
}

func HadleRequest() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/res", ResponseCheck)
	log.Fatal(http.ListenAndServe(":1221", nil))
}

func ResponseCheck(w http.ResponseWriter, r *http.Request) {
	//  w.
	// WriteHeader(200)
	// json("Hello THis is response page")
	fmt.Println("Response printed")
}
