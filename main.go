package main

import (
    "fmt"
//    "html"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/metadata/v1.json", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "[{'droplet_id': 2756294,'hostname': 'sample-droplet','vendor_data': '#cloud-config','public_keys': [],'region': 'nyc3','interfaces': {'private': [],'public': []},'floating_ip': {},'dns': {},'features': {}}]" )
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}