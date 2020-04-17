package main

import (
	"go-http-example/app"
	_ "go-http-example/app"
)

func main() {
	application := app.App{}
	application.Initialize()
	application.Run(":8080")
}
