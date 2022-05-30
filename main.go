package main

import (
    "fmt"
//    "html"
    "log"
    "net/http"
)

func main() {

    myJsonString := `{"droplet_id":301981521,"hostname":"centos-s-1vcpu-1gb-fra1-01","public_keys":["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVFOL2fXS5A4Lf5/MgZGuj9CVyRzjND2hFXZ/D2LR+CZ8dNoCTEEuAKzWdH4OjYQvNZOtHzWsd4v5R1CddfpPFEce3F8gCLelchFz+Sy1Zcua3q81qnj6kDMszcP6NuwF8FtIc5/M2oR0S00wAd7DOV+ZIFFUUS9uBdYaGKVK9MGj/M/zVmg2ic/PPvfRnMrL08oRZWah+DaK8X8sI0xdWATKtX0xgkrcrdbnelGbUY6vrIm7YNsRnhxfQbgTguGiU6QIzKrQT8UlXSbzImwbyo48HsMpNCaadOy8pep54bu6acVebeJF7QcZP6drmv670xIbNvvUfBkUYYsGo1DpLpp9Ia6YWfYnPmwIina+LHhkKv04/+dreH99ptFcDIcN3uL6VclntFJaQJlIBgFUyoJVq/DH2yKIgZwkQ6eAoVzbzdLR+IpKUJR8/ShicGhHQQzsibwT0ydq5DPP6VaF3vPE8ULONwVuR63sry42jLxyqbl7oZvwuiCcSjdZsTh8= wolfgang.huse@C02WF0Z7HV2Q"],"auth_key":"6703d9044e28270ec024727b899b1766","region":"fra1","interfaces":{"private":[{"ipv4":{"ip_address":"10.114.0.2","netmask":"255.255.240.0","gateway":"10.114.0.1"},"mac":"ae:8a:a5:54:06:27","type":"private"}],"public":[{"ipv4":{"ip_address":"142.93.108.93","netmask":"255.255.240.0","gateway":"142.93.96.1"},"anchor_ipv4":{"ip_address":"10.19.0.5","netmask":"255.255.0.0","gateway":"10.19.0.1"},"mac":"ca:cd:a7:1a:4d:7e","type":"public"}]},"floating_ip":{"ipv4":{"active":false}},"dns":{"nameservers":["67.207.67.2","67.207.67.3"]},"tags":null,"features":{"dhcp_enabled":false},"modify_index":146566526,"dotty_keys":["{\"os_user\":\"root\",\"ssh_key\":\"ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBNK99I8N/oyik3EDTYanxRdSHzMZ22aw8w1Lh5Y0zPdQjV4KG/PglUZS95TJcsJzJTxRPKRzXwYomkV6Z5uCp0=\",\"actor_email\":\"wolfgang.huse@nutanix.com\",\"ttl\":165}","{\"os_user\":\"root\",\"ssh_key\":\"ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGkoUc9ztsvGiHNewGI7VOnI6QorZCLTEh2hz8BDXK0g0Bdeq6plsfQkCG0alQPbCZrg6SQKtkKv245p9qbQdOk=\",\"actor_email\":\"wolfgang.huse@nutanix.com\",\"ttl\":180}"],"dotty_status":"running","ssh_info":{"port":22}}`

    
    http.HandleFunc("/metadata/v1.json", func(w http.ResponseWriter, r *http.Request) {
        //fmt.Fprintf(w, "[{'droplet_id': 2756294,'hostname': 'sample-droplet','vendor_data': '#cloud-config','public_keys': [],'region': 'nyc3','interfaces': {'private': [],'public': []},'floating_ip': {},'dns': {},'features': {}}]" )
        fmt.Fprint(w,myJsonString)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}