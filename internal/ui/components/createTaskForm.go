package components

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/zarasfara/to-do-list/internal/file"
)

func NewCreateModelForm(window fyne.Window, table *TaskTable) {

	// Создаем форму с необходимыми элементами
	titleEntry := &widget.Entry{
		Validator: func(text string) error {
			if len(text) == 0 {
				return fmt.Errorf("необходимо ввести значение")
			}
			return nil
		},
	}

	descriptionEntry := &widget.Entry{
		Validator: func(text string) error {
			if len(text) > 50 {
				return fmt.Errorf("слишком длинное предложение")
			}
			return nil
		}}
	categoryEntry := widget.NewEntry()

	items := []*widget.FormItem{
		{
			Text:   "Название",
			Widget: titleEntry,
		},
		{
			Text:   "Описание",
			Widget: descriptionEntry,
		},
		{
			Text:   "Категория",
			Widget: categoryEntry,
		},
	}

	// Создаем диалог с формой
	formDialog := dialog.NewForm("Создать задачу", "Подтвердить", "Закрыть", items, func(ok bool) {
		if ok {
			// Если нажата кнопка "Submit", выводим данные из формы в консоль

			err := file.WriteTaskToFile(titleEntry.Text, descriptionEntry.Text, categoryEntry.Text)
			if err != nil {
				_ = fmt.Errorf("что-то пошло не так")
			}

			table.RefreshTable() // вызываем метод RefreshTable у таблицы
		}
	}, window)

	formDialog.Resize(fyne.NewSize(600, 400))

	// Отображаем диалог с формой
	formDialog.Show()
}
