//
//  CAS Server (beta)
//
//  Created by Marko Tunjic on 15/07/16.
//  Copyright © 2016 Marko Tunjic. All rights reserved.
//
package main

import (
	"log"
	"net/http"
)

//User for login credentials
type User struct {
	UserName   string `json:"username"`
	Password   string `json:"password"`
	LT         string `json:"lt"`
	RememberMe bool   `json:"rememberme"`
}

// login - credential requestor / acceptor
func loginHandler(u User) http.Handler {
	// validate user credentials
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if u.UserName != "demo" && u.Password != "pass" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("We don't recognize that user. Please try again. ."))
		} else {

			w.Write([]byte(u.UserName + " you have successfully logged in."))
		}
	})
}

// logout - destroy CAS session (logout)
func logoutHandler(u string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(u + " you have successfully been logged out."))
	})
}

// validate - service ticket validation
func validateHandler(u string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("service ticket validation"))
	})
}

// serviceValidate - service ticket validation
func serviceValidateHandler(u string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("service ticket validation"))
	})
}

// proxyValidate - service/proxy ticket validation
func proxyValidateHandler(u string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("service/proxy ticket validation"))
	})
}

// /proxy - proxy ticket service
func proxyHandler(u string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("proxy ticket service"))
	})
}

func main() {

	user := User{"demo", "pass", "27718", true}
	http.Handle("/login", loginHandler(user))

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
