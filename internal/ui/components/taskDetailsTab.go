package components

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/zarasfara/to-do-list/internal/file"
)
var nameEntry = widget.NewEntry()
var descEntry = widget.NewMultiLineEntry()
var categoryEntry = widget.NewEntry()
var doneCheck = widget.NewCheck("Выполнено", func(checked bool) {})

func NewDetailsTab(window fyne.Window) *widget.Form {

	nameEntry.Validator = func(text string) error {
		if len(text) == 0 {
			return fmt.Errorf("поле не может быть пустым")
		}

		return nil
	}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Название", Widget: nameEntry},
			{Text: "Описание", Widget: descEntry},
			{Text: "Категория", Widget: categoryEntry},
			{Text: "Выполнено", Widget: doneCheck},
		},
		OnSubmit: func() {
			name := nameEntry.Text
			desc := descEntry.Text
			done := doneCheck.Checked

			success, _ := file.UpdateTask(CurrentTaskId, name, desc, done)

			if success {
				infoDialog := dialog.NewInformation(
					"Успех!",
					"Успешно обновлено",
					window)
				infoDialog.SetDismissText("Закрыть")
				infoDialog.Show()
			}

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
