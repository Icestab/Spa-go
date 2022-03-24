package main

import (
	"spa-go/cmd/app/handlers/router"
	"spa-go/internal/utils/conf"
	"spa-go/internal/utils/db/sqlite"
)

func main() {
	run()
}

func run() {
	conf.Init()
	sqlite.Init()
	//auth.Init()
	router.Init()
}
