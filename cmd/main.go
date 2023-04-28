package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

func main() {
	todoApp := app.New()
	window := todoApp.NewWindow("Todo List")

	icon, err := fyne.LoadResourceFromPath("assets/icon.png")
	if err != nil {
		log.Printf("ошибка при загрузке иконки: %s", err.Error())
	}

	window.SetIcon(icon)

	nameEntry := &widget.Entry{
		PlaceHolder: "Введите название",
	}

	descriptionEntry := &widget.Entry{
		PlaceHolder: "Введите описание",
	}

	// Сайд бар левая часть
	list := container.NewVBox(
		&widget.Label{
			Text:      "Типы",
			Alignment: 1,
		},
		widget.NewButton("Элемент 1", func() {
			fmt.Print("1")
		}),
		widget.NewButton("Элемент 2", func() {
			fmt.Print("2")
		}),
		widget.NewButton("Элемент 2", func() {
			fmt.Print("2")
		}))

	// Кнопка "смена темы"
	toggleThemeButton := widget.NewButton("Сменить тему", func() {
		if todoApp.Settings().ThemeVariant() == theme.VariantDark {
			todoApp.Settings().SetTheme(theme.LightTheme())
		} else {
			todoApp.Settings().SetTheme(theme.DarkTheme())
		}
	})

	sidebar := container.NewVBox(
		list,
		layout.NewSpacer(),
		toggleThemeButton,
	)

	// Правая часть
	rightPart := container.NewVBox(nameEntry, descriptionEntry)

	content := container.NewBorder(nil, nil, sidebar, nil, rightPart)

	window.Resize(fyne.NewSize(800, 600))
	window.SetContent(content)
	window.CenterOnScreen()
	window.ShowAndRun()
}
