package components

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/zarasfara/to-do-list/internal/file"
)

func NewButtonsContainer(win fyne.Window, table *TaskTable) *fyne.Container {
	return container.NewVBox(
		layout.NewSpacer(),
		container.NewHBox(
			// Кнопка вызова модалки создания таски
			widget.NewButton("Создать задачу", func() {
				// Создаем диалог с формой
				NewCreateModalForm(win, table)
			}),
			widget.NewButton("Удалить задачу", func() {
				err := file.DeleteTask(CurrentTaskId)
				if err != nil {
					_ = fmt.Errorf("ошибка")
				}
				CurrentTaskId = -1
				table.RefreshTable()
			}),
		))
}
