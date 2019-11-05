package main

import (
    "flag"
    "log"
    "net/http"
)

func main() {
    port := flag.String("p", "8080", "Port for the server")
    dir := flag.String("d", ".", "Directory to server")

    flag.Parse()

    portDef := ":" + *port

    log.Println("Listening at", "localhost" + portDef)    
    log.Fatal(http.ListenAndServe(portDef, http.FileServer(http.Dir(*dir))))
}
