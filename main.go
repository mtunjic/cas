package main

import (
	"log"
	"net/http"
)

func loginHandler(u string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(u + " you have successfully logged in."))
	})
}

func main() {

	http.Handle("/login", loginHandler("demo"))

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
