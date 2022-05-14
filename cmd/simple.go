package main

import (
	"os"

	"github.com/narumiruna/wolframalpha/pkg/simple"
)

func main() {
	appID := os.Getenv("WOLFRAMALPHA_APP_ID")
	w := simple.New(appID)
	w.QueryToFile("taiwan", "taiwan.jpg")
}