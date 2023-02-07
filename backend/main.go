package main

import (
	"fmt"
	"log"

	"forum-backend/internal/app"
)

func main() {
	fmt.Println("asd")
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
