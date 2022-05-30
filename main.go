package main

import (
	"encoding/json"
	"fmt"
    "log"
    "net/http"
)
type ipv4 struct{
    Ip_address string `json:"ip_address"`
    Netmask string `json:"netmask"`
    Gateway string `json:"gateway"`
}

type net_interface struct {
    Ipv4 ipv4 `json:"ipv4"`
    Anchor_ipv4 ipv4 `json:"anchor_ipv4,omitempty"`
    Mac string `json:"mac"`
    Interface_type string `json:"type"`
}

type interfaces struct {
    Private net_interface `json:"private"`
    Public net_interface `json:"public"`
}

type metadata struct {
    Hostname string `json:"hostname"`
    Region string `json:"region"`
    Interfaces interfaces `json:"interfaces"`
}
func main() {

    mydata := &metadata{
		Hostname: "Pigeon",
        Region: "fra",
	}

	data, _ := json.Marshal(mydata)

	fmt.Println(string(data))

    // myJsonString := `{"hostname":"centos-s-1vcpu-1gb-fra1-01","region":"fra1","interfaces":{"private":[{"ipv4":{"ip_address":"10.114.0.2","netmask":"255.255.240.0","gateway":"10.114.0.1"},"mac":"ae:8a:a5:54:06:27","type":"private"}],"public":[{"ipv4":{"ip_address":"142.93.108.93","netmask":"255.255.240.0","gateway":"142.93.96.1"},"anchor_ipv4":{"ip_address":"10.19.0.5","netmask":"255.255.0.0","gateway":"10.19.0.1"},"mac":"ca:cd:a7:1a:4d:7e","type":"public"}]},"floating_ip":{"ipv4":{"active":false}}}`
    http.HandleFunc("/metadata/v1.json", func(w http.ResponseWriter, r *http.Request) {
        //fmt.Fprintf(w, "[{'droplet_id': 2756294,'hostname': 'sample-droplet','vendor_data': '#cloud-config','public_keys': [],'region': 'nyc3','interfaces': {'private': [],'public': []},'floating_ip': {},'dns': {},'features': {}}]" )
        fmt.Fprint(w,string(data))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}