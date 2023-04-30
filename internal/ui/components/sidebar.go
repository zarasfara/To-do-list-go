package components

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var themeVar = true

func NewSideBar(app fyne.App) *fyne.Container {

	// Сайд бар левая часть
	list := container.NewVBox(
		&widget.Label{
			Text:      "Фильтры",
			Alignment: 1,
		},
		widget.NewButton("Не выполненные", func() {
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
		if themeVar {
			// Текущая тема - темная
			app.Settings().SetTheme(theme.LightTheme())
			themeVar = false
		} else {
			// Текущая тема - светлая
			app.Settings().SetTheme(theme.DarkTheme())
			themeVar = true
		}
	})

	return container.NewVBox(
		list,
		layout.NewSpacer(),
		toggleThemeButton,
	)
}
