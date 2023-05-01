package components

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"github.com/zarasfara/to-do-list/internal/file"
)

var testId = -1

var nameEntry = widget.NewEntry()
var descEntry = widget.NewMultiLineEntry()
var doneCheck = widget.NewCheck("Выполнено", func(checked bool) {
	// Здесь можно добавить код, который будет выполнен при изменении состояния doneCheck
})

func NewDetailsTab(id int) *widget.Form {

	task, _ := file.GetTaskById(id)

	if task != nil {
		nameEntry.SetText(task.Title)
		nameEntry.SetText(task.Description)
	}

	return &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Название", Widget: nameEntry},
			{Text: "Описание", Widget: descEntry},
			{Text: "Выполнено", Widget: doneCheck},
		},
		OnSubmit: func() {
			name := nameEntry.Text
			desc := descEntry.Text
			done := doneCheck.Checked
			fmt.Printf("Название: %s\nОписание: %s\nВыполнено: %v\n", name, desc, done)
		},
	}
}

func UpdateForm(id int) {
	if id == 0 {
		return
	}

	task, _ := file.GetTaskById(id)

	if task == nil {
		// обработка ошибки, например:
		return
	}

	nameEntry.SetText(task.Title)
	descEntry.SetText(task.Description)
}
