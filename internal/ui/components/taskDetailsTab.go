package components

import (
	"fmt"
	"time"

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
		if len(text) == 0 && CurrentTaskId > 0 {
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
			category := categoryEntry.Text

			success, _ := file.UpdateTask(CurrentTaskId, name, desc, done, category)

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
		categoryEntry.SetText("")
		doneCheck.Checked = false
		return
	}

	nameEntry.SetText(task.Title)
	descEntry.SetText(task.Description)
	categoryEntry.SetText(task.Category)

	if task.Completed {
		doneCheck.Checked = true
	} else {
		doneCheck.Checked = false
	}
}

func CheckTime(win fyne.Window) {
	for range time.Tick(time.Minute) {
		nowFormatted := time.Now().Format("02.01.2006 15:04")

		for i := 0; i < len(Reminders); i++ {
			v := Reminders[i]
			dateFormatted := v.Date.Format("02.01.2006 15:04")
			if dateFormatted <= nowFormatted {
				win.Show()
				dialog.ShowInformation("Напоминание", v.Title, win)
				// Удаляем элемент из slice
				fmt.Println("До: ", Reminders)
				Reminders = RemoveItem(Reminders, i)
				fmt.Println("После: ", Reminders)
				i-- // делаем decrement для индексов
			}
		}
	}
}