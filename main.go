package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/metadata/v1.json", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "{NUTANIX_HOSTNAME:abc}" )
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}