package main

import (
	"os"

	"github.com/narumiruna/wolframalpha/pkg/simple"
)

func main() {
	appID := os.Getenv("WOLFRAMALPHA_APP_ID")
	c := simple.New(appID)
	c.QueryFile("taiwan", "taiwan.jpg", nil)
}
