package main

import (
	"mvc-goALUMNO/app"
	"mvc-goALUMNO/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
