package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
)

var appTheme = true

func NewMainMenu(app fyne.App, window fyne.Window) *fyne.MainMenu {

	settingsAbout := fyne.NewMenuItem("О программе", func() {
		infoDialog := dialog.NewInformation(
			"О программе",
			"Версия: 1.0\nАвтор: Евгений Осипов\n",
			window)
		infoDialog.SetDismissText("Закрыть")
		infoDialog.Show()
	})

	// Создаем меню с пунктом "Настройки"
	settingsMenu := fyne.NewMenuItem("Поменять тему", func() {
		if appTheme {
			// Текущая тема - темная
			app.Settings().SetTheme(theme.LightTheme())
			appTheme = false
		} else {
			// Текущая тема - светлая
			app.Settings().SetTheme(theme.DarkTheme())
			appTheme = true
		}
	})

	return fyne.NewMainMenu(
		fyne.NewMenu("О программе", settingsAbout),
		fyne.NewMenu("Настройки", settingsMenu),
	)
}
