package main

import (
	"go_mockery/cmd/watch/app"
	"go_mockery/cmd/watch/repository"
)

var (
	watchRepo *repository.WatchRepo
	runApp    *app.App
)

func init() {
	watchRepo = repository.NewWatchRepo()
	runApp = app.NewApp(watchRepo)
	runApp.Run()
}

func main() {
}
