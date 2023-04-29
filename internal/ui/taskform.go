package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewTaskForm() *fyne.Container {
	nameEntry := &widget.Entry{
		PlaceHolder: "Введите название",
	}

	descriptionEntry := &widget.Entry{
		PlaceHolder: "Введите описание",
	}

	return container.NewVBox(nameEntry, descriptionEntry)
}
