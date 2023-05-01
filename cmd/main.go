package main

import "github.com/zarasfara/to-do-list/internal/app"

func main() {

	myApp := app.NewTodoApp()

	myApp.Run()

}
