package main

import (
	"fangkm.demo/shortlink"
)

func main() {
	app:=shortlink.NewApp(":4567")
	app.Run()
}