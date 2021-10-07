package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("client request ip : %s",r.RemoteAddr)

	w.Header().Set("VERSION",runtime.Version())
	for k,v := range r.Header {
		w.Header().Set(k,v[0])
	}

    fmt.Fprintf(w, "Welcome")
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("client request ip : %s",r.RemoteAddr)

	w.Header().Set("VERSION",runtime.Version())
	for k,v := range r.Header {
		w.Header().Set(k,v[0])
	}
	
	w.WriteHeader(200)
	fmt.Fprintf(w, "OK")
}


func main () {
    http.HandleFunc("/",RootHandler)
	http.HandleFunc("/healthz",HealthHandler)

    http.ListenAndServe(":8000", nil)
}

