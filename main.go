package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8080", "Port for the server")
	dir := flag.String("d", ".", "Directory to server")
	quiet := flag.Bool("q", false, "Set for quiet mode")

	flag.Parse()

	portDef := ":" + *port

	if !*quiet {
		log.Println("Listening at", "localhost"+portDef)
	}

	fileHandler := http.FileServer(http.Dir(*dir))

	log.Fatal(http.ListenAndServe(portDef, logHandler(fileHandler, *quiet)))
}

func logHandler(h http.Handler, quiet bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !quiet {
			log.Printf("%s %s\n", r.Method, r.URL.Path)
		}

		h.ServeHTTP(w, r)
	})
}
