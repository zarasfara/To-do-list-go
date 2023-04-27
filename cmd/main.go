package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Todo List")

	// Создание виджетов для todo list
	taskEntry := widget.NewEntry()

	taskList := container.NewHBox()

	addButton := widget.NewButton("Добавить", func() {
		task := taskEntry.Text
		if task != "" {
			taskItem := widget.NewLabel(task)
			taskList.Add(taskItem)
			taskEntry.SetText("")
		}
	})

	// Создание сайдбара
	sidebar := container.NewVBox(
		widget.NewLabel("Категории"),
		widget.NewButton("All", func() {

		}),
		widget.NewButton("Work", func() {

		}),
		widget.NewButton("Personal", func() {

		}),
		widget.NewButton("Shopping", func() {

		}),
		layout.NewSpacer(),
	)

	// Создание главного контейнера
	content := container.NewBorder(
		nil, // top
		nil,
		sidebar, // bottom
		nil,     // left
		container.NewVBox(
			taskEntry,
			addButton,
			taskList,
		), // right
	)

	window.Resize(fyne.NewSize(800, 600))
	window.SetContent(content)
	window.ShowAndRun()
}
