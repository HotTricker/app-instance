package main

import (
	miniapp "app-instance/internal/mini-app"
	"fmt"
	"log"
)

func main() {
	fmt.Println("welcome")
	miniapp.OutputInfo("Version", miniapp.Version)
	if err := miniapp.App.Init("configs/mini-app.ini"); err != nil {
		log.Fatal(err)
	}

	if err := miniapp.App.Start(); err != nil {
		log.Fatal(err)
	}
}
