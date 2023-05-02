package components

import (
	"fyne.io/fyne/v2/widget"
	"github.com/zarasfara/to-do-list/internal/file"
)

var nameEntry = widget.NewEntry()
var descEntry = widget.NewMultiLineEntry()
var doneCheck = widget.NewCheck("Выполнено", func(checked bool) {
	// Здесь можно добавить код, который будет выполнен при изменении состояния doneCheck
})

func NewDetailsTab() *widget.Form {

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Название", Widget: nameEntry},
			{Text: "Описание", Widget: descEntry},
			{Text: "Выполнено", Widget: doneCheck},
		},
		OnSubmit: func() {
			name := nameEntry.Text
			desc := descEntry.Text
			done := doneCheck.Checked

			file.UpdateTask(CurrentTaskId, name, desc, done)

			GetTasksTable().RefreshTable()
		},
	}
	form.SubmitText = "Подтвердить"

	return form
}

func UpdateForm(id int) {

	task, _ := file.GetTaskById(id)

	if id < 0 || task == nil {
		nameEntry.SetText("")
		descEntry.SetText("")
		doneCheck.Checked = false
		return
	}

	nameEntry.SetText(task.Title)
	descEntry.SetText(task.Description)
	if task.Completed {
		doneCheck.Checked = true
	} else {
		doneCheck.Checked = false
	}
}
