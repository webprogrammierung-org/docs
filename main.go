package main

import "net/http"

func main() {

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("build/html"))
	mux.Handle("/", http.StripPrefix("/", fileServer))
	error := http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/webprogrammierung.org/cert.pem", "/etc/letsencrypt/live/webprogrammierung.org/privkey.pem", mux)
	if error != nil {
		http.ListenAndServe(":80", mux)

	}
}
