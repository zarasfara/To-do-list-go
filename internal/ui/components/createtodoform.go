package components

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/zarasfara/to-do-list/internal/file"
)

func NewCreateModelForm(window fyne.Window) {

	// Создаем форму с необходимыми элементами
	titleEntry := widget.NewEntry()
	descriptionEntry := widget.NewEntry()
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
	formDialog := dialog.NewForm("Введите данные", "Подтвердить", "Закрыть", items, func(ok bool) {
		if ok {
			// Если нажата кнопка "Submit", выводим данные из формы в консоль
			// fmt.Printf("title: %s\ndescription: %s\ncategory: %s\n", titleEntry.Text, descriptionEntry.Text, categoryEntry.Text)

			err := file.WriteTasksToFile(titleEntry.Text, descriptionEntry.Text, categoryEntry.Text)
			if err != nil {
				fmt.Errorf("что-то пошло не так")
			}
		}
	}, window)

	formDialog.Resize(fyne.NewSize(500, 300))

	// Отображаем диалог с формой
	formDialog.Show()
}
