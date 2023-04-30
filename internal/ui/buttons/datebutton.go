package buttons

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"time"
)

func NewDateButton() *widget.Entry {
	dateEntry := widget.NewEntry()
	dateEntry.SetPlaceHolder("DD.MM.YYYY")

	dateEntry.Validator = func(text string) error {

		if text == "" {
			return fmt.Errorf("необходимо ввести значение")

		}

		_, err := time.Parse("02.01.2006", text)
		if err != nil {
			return fmt.Errorf("неверный формат")
		}
		return nil
	}

	return dateEntry
}
