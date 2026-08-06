package main

import "shihua_algo/proxy"

func main() {
	proxy_instance := proxy.NewProxy()
	proxy_instance.Execute("yes")
}

//$ go run main.go
//Proxy Service: yes
