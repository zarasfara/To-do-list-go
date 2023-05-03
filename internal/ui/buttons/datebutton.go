package buttons

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"time"
)

func NewDateEntry() *widget.Entry {
	dateEntry := widget.NewEntry()
	dateEntry.SetPlaceHolder("02.01.2006 15:04")

	dateEntry.Validator = func(text string) error {

		if len(text) == 0 {
			return fmt.Errorf("необходимо ввести значение")
		}

		// Парсим дату
		date, err := time.Parse("02.01.2006 15:04", text)
		if err != nil {
			return fmt.Errorf("неверный формат")
		}

		// Получаем текущую дату
		now := time.Now()

		// Сравниваем даты
		if date.Before(now) {
			return fmt.Errorf("дата не может быть прошедшей")
		}

		return nil
	}

	return dateEntry
}

