package main

import "form-manager/pkg/di"

func main() {
	server := di.InitDI()
	server.Start()
}
