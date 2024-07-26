package main

import (
	"log"
	"net/http"

	"github.com/Nezent/GoJWT/handlers"
)

func main() {
	http.HandleFunc("/login",handlers.Login)
	// http.HandleFunc("/home",Home)
	// http.HandleFunc("/refresh",Refresh)

	log.Fatal(http.ListenAndServe(":8080",nil))
}