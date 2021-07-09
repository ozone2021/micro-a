package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {

    port := os.Getenv("MICRO_A_PORT")

    if port == "" {
        log.Fatalln("MICRO_A_PORT env var not set.")
    }

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

        w.WriteHeader(http.StatusOK)
        msg := fmt.Sprintf("Runnable MICRO-A is OK.")
        log.Println("logging")
        w.Write([]byte(msg))
    })

    addr := fmt.Sprintf(":%s", port)
    http.ListenAndServe(addr, nil)
}