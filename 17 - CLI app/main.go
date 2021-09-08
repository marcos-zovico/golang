package main

import (
	"cli-app/app"
	"log"
	"os"
)

func main() {
	appication := app.Gerar()
	if erro := appication.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
