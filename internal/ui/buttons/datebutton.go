package buttons

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"time"
)

func NewDateEntry() *widget.Entry {
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

		now := time.Now() // 03.05.2023 15:47
		nowFormatted := now.Format("02.01.2006 15:04")
		dateFormatted:= date.Format("02.01.2006 15:04")
		fmt.Println(dateFormatted)
		fmt.Println(nowFormatted)
		if dateFormatted < nowFormatted {
			return fmt.Errorf("aboba")
		}

		return nil
	}

	return dateEntry
}
