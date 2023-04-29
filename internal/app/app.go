package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/zarasfara/to-do-list/internal/ui"
	"log"
)

type TaskApp struct {
	fyne.App
}

func NewTodoApp() *TaskApp {
	fyneApp := app.New()
	fyneApp.Settings().SetTheme(theme.DarkTheme())

	return &TaskApp{
		fyneApp,
	}
}

func (a *TaskApp) Run() {
	fyneApp := app.New()
	fyneApp.Settings().SetTheme(theme.DarkTheme())

	win := a.NewWindow("Todo List")

	icon, err := fyne.LoadResourceFromPath("assets/icon.png")
	if err != nil {
		log.Printf("ошибка при загрузке иконки: %s", err.Error())
	}
	win.SetIcon(icon)

	rightPart := ui.NewTaskForm()

	sidebar := ui.NewSideBar(a)

	content := container.NewBorder(nil, nil, sidebar, nil, rightPart)

	win.Resize(fyne.NewSize(800, 600))
	win.SetContent(content)
	win.CenterOnScreen()
	win.ShowAndRun()

}
