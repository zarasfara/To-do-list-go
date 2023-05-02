package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
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

func GetTasksTable() *TaskApp {
	return &TaskApp{}
}

func (a *TaskApp) Run() {
	win := a.NewWindow("Todo List")

	if desk, ok := a.App.(desktop.App); ok {
		m := fyne.NewMenu("myApp",
			fyne.NewMenuItem("Показать", func() {
				win.Show()
			}))

		desk.SetSystemTrayMenu(m)
	}

	// Таблица с задачами
	table := components.NewTasksTable()

	// Контейнер с задачами
	buttonContainer := components.NewButtonsContainer(win, table)

	// Главное меню
	menu := components.NewMainMenu(a, win)

	// Шаблон для задач
	taskContent := container.NewGridWithRows(2, buttonContainer, table)

	taskMoreContent := components.NewDetailsTab(win)

	tabItems := []*container.TabItem{
		container.NewTabItemWithIcon("Задачи", theme.ContentCopyIcon(), taskContent),
		// container.NewTabItemWithIcon("Напоминания", theme.ErrorIcon(), widget.NewLabel("Напоминания")),
		container.NewTabItemWithIcon("Подробнее", theme.DocumentCreateIcon(), taskMoreContent),
	}

	/* Tabs */
	content := container.NewAppTabs(tabItems...)

	content.OnSelected = func(tabItem *container.TabItem) {
		if tabItem.Text == "Подробнее" {
			components.UpdateForm(components.CurrentTaskId)
		}
	}

	//---------------------------------------//
	icon, err := fyne.LoadResourceFromPath("assets/icon.png")
	if err != nil {
		log.Printf("ошибка при загрузке иконки: %s", err.Error())
	}
	win.SetIcon(icon)
	win.Resize(fyne.NewSize(1200, 800))
	win.SetMainMenu(menu)
	win.SetCloseIntercept(func() {
		win.Hide()
	})
	win.SetContent(content)
	win.CenterOnScreen()
	win.ShowAndRun()
}
