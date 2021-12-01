package main

import (
	"todo/app"

	"todo/core"
)

func main() {
	config := core.NewConfig()
	app := &app.App{}
	app.Start(config)

}
