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

	toggleFormBtn := container.NewPadded(widget.NewButton("Создать задачу", func() {
		// Создаем диалог с формой
		components.NewCreateModelForm(win)
	}))

	sidebar := components.NewSideBar(a) // Sidebar

	// Создаем заголовки таблицы
	headers := []string{"Название", "Описание", "Дата"}

	table := widget.NewTable(
		func() (int, int) {
			return 3, 3 // 2 столбца, 3 строки
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Cell")
		},
		func(cell widget.TableCellID, cellView fyne.CanvasObject) {
			if cell.Row == 0 {
				// Устанавливаем заголовки
				cellView.(*widget.Label).SetText(headers[cell.Col])
			} else {
				// Устанавливаем значения ячеек таблицы
				cellView.(*widget.Label).SetText("тварь, мразь")
			}
		},
	)

	numColumns := 3                  // задаем количество колонок
	columnWidth := 1000 / numColumns // вычисляем ширину колонки
	for i := 0; i < numColumns; i++ {
		table.SetColumnWidth(i, float32(columnWidth))
	}

	taskContent := container.NewBorder(
		nil,
		nil,
		sidebar, // left
		nil,
		container.NewGridWithRows(2, container.NewHBox(container.NewVBox(toggleFormBtn)), table), // objects
	)
	t := container.NewAppTabs(container.NewTabItem("Задачи", taskContent), container.NewTabItem("Напоминания", widget.NewLabel("Напоминания")))

	//---------------------------------------//
	win.Resize(fyne.NewSize(1200, 800))
	win.SetFixedSize(true)
	win.SetContent(t)
	win.CenterOnScreen()
	win.ShowAndRun()
}
