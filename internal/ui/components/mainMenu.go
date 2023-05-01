package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
)

var themeVar = true

func NewMainMenu(app fyne.App, window fyne.Window) *fyne.MainMenu {

	settingsAbout := fyne.NewMenuItem("О программе", func() {
		mydialog := dialog.NewInformation(
			"О программе",
			"Версия: 1.0\nАвтор: Евгений Осипов\n",
			window)
		mydialog.SetDismissText("Закрыть")
		mydialog.Show()
	})

	// Создаем меню с пунктом "Настройки"
	settingsMenu := fyne.NewMenuItem("Поменять тему", func() {
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

	return fyne.NewMainMenu(
		fyne.NewMenu("О программе", settingsAbout),
		fyne.NewMenu("Настройки", settingsMenu),
	)
}
