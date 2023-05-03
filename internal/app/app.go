package app

import (
	"fmt"
	"log"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"github.com/zarasfara/to-do-list/internal/ui/components"
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

	// Шаблон для нопоминаний
	reminderContent := container.NewGridWithRows(2,
		components.NewReminderForm(),
	)

	var mutex sync.Mutex
	go func() {
		for range time.Tick(time.Minute) {
			nowFormatted := time.Now().Format("02.01.2006 15:04")

			mutex.Lock()
			for i, v := range components.Reminders {
				dateFormatted := v.Date.Format("02.01.2006 15:04")
				if dateFormatted <= nowFormatted {
					win.Show()
					dialog.ShowInformation("Напоминание", v.Title, win)
					// Удаляем элемент из slice
					components.Reminders = components.RemoveItem(components.Reminders, i)
					fmt.Println(components.Reminders)
				}
			}
			mutex.Unlock()
		}
	}()

	taskMoreContent := components.NewDetailsTab(win)

	tabItems := []*container.TabItem{
		container.NewTabItemWithIcon("Задачи", theme.ContentCopyIcon(), taskContent),
		container.NewTabItemWithIcon("Напоминания", theme.ErrorIcon(), reminderContent),
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
