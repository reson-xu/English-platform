package main

import (
	"fmt"

	"github.com/reson-xu/english-platform/internal/app"
)

func main() {
	desktopApp := app.New()
	fmt.Println(desktopApp.Name())
}
