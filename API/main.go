package main

import "triela/Routing"

func main() {
	app := Routing.CatchRouting()

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
