package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewSideBar(app fyne.App) *fyne.Container {

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
		if app.Settings().ThemeVariant() == theme.VariantDark {
			app.Settings().SetTheme(theme.LightTheme())
		} else {
			app.Settings().SetTheme(theme.DarkTheme())
		}
	})

	return container.NewVBox(
		list,
		layout.NewSpacer(),
		toggleThemeButton,
	)

}
