package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/zarasfara/to-do-list/internal/ui/components"
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
	win := a.NewWindow("Todo List")

	icon, err := fyne.LoadResourceFromPath("assets/icon.png")
	if err != nil {
		log.Printf("ошибка при загрузке иконки: %s", err.Error())
	}
	win.SetIcon(icon)

	// Сайд бар
	sidebar := components.NewSideBar(a) // Sidebar

	// Таблица с задачами
	table := components.NewTasksTable()

	// Кнопка вызова модалки создания таски
	toggleFormBtn := container.NewPadded(widget.NewButton("Создать задачу", func() {
		// Создаем диалог с формой
		table.Refresh()
	}))

	// Задаем контент для задач
	taskContent := container.NewBorder(
		nil,
		nil,
		sidebar, // left
		nil,
		container.NewGridWithRows(2, container.NewHBox(container.NewVBox(toggleFormBtn)), table), // objects
	)

	// Задаем контент для напоминаний
	// ...

	content := container.NewAppTabs(
		container.NewTabItem("Задачи", taskContent),
		container.NewTabItem("Напоминания", widget.NewLabel("Напоминания")),
	)

	//---------------------------------------//
	win.Resize(fyne.NewSize(1200, 800))
	win.SetFixedSize(true)
	win.SetContent(content)
	win.CenterOnScreen()
	win.ShowAndRun()
}
