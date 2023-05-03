package components

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/zarasfara/to-do-list/internal/ui/buttons"
)

type Remind struct {
	Title string
	Date  time.Time
}

var Reminders []Remind

func NewReminderForm(window fyne.Window) *widget.Form {
	datetest := buttons.NewDateEntry()

	title := widget.NewEntry()

	title.Validator = func(text string) error {
		// Проверяем, что значение не пустое
		if len(text) == 0 {
			return fmt.Errorf("поле не может быть пустым")
		}

		return nil
	}

	titleEntry := widget.NewFormItem("Название", title)

	dateEntry := widget.NewFormItem("Дата", datetest)

	reminderForm := widget.NewForm(
		titleEntry,
		dateEntry,
	)

	reminderForm.SubmitText = "Подтвердить"

	reminderForm.OnSubmit = func() {
		dateTimeLayout := "02.01.2006 15:04"
		dateTimeString := dateEntry.Widget.(*widget.Entry).Text

		date, err := time.Parse(dateTimeLayout, dateTimeString)
		if err != nil {
			fmt.Println(err)
			return
		}

		reminder := Remind{
			Title: titleEntry.Widget.(*widget.Entry).Text,
			Date:  date,
		}

		Reminders = append(Reminders, reminder)

		infoDialog := dialog.NewInformation(
			"Успех!",
			"Успешно создано",
			window)
		infoDialog.SetDismissText("Закрыть")
		infoDialog.Show()
	}

	return reminderForm
}

func RemoveItem(slice []Remind, index int) []Remind {
    if index < 0 || index >= len(slice) {
        return slice
    }
    if len(slice) == 1 {
        return []Remind{}
    }
    return append(slice[:index], slice[index+1:]...)
}
