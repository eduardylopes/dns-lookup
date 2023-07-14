package main

import (
	"log"
	"main/app"
	"os"
)

func main() {
	application := app.Generate()

	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}

}
