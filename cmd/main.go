package main

import . "eventure_backend/internal/infrastructure"

func main() {
	server := InitializeApp()
	server.Start()
}
