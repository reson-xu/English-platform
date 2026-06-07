package main

import (
	"os"

	"github.com/reson-xu/english-platform/internal/app"
)

func main() {
	if err := app.RunAPI(); err != nil {
		os.Exit(1)
	}
}
