package main

import "log"
import "net/http"
import "github.com/robel-yemane/go-in-action/Testing/listing17/handlers"

// main is the entry point for the application
func main() {
    handlers.Routes()
    
    log.Println("listener : Started : Listening on :4000")
    http.ListenAndServe(":4000", nil)
}

