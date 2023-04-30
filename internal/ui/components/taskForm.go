package components

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/zarasfara/to-do-list/internal/ui/buttons"
)

func NewTaskForm() *fyne.Container {

	titleEntry := widget.NewFormItem("", &widget.Entry{
		PlaceHolder: "Введите название",
		Validator:   validateTitle,
	})

	descriptionEntry := widget.NewFormItem("", &widget.Entry{
		PlaceHolder: "Введите описание",
	})

	dateEntry := widget.NewFormItem("", buttons.NewDateEntry())

	form := widget.NewForm(
		titleEntry,
		descriptionEntry,
		dateEntry,
	)

	form.SubmitText = "Создать"

	form.OnSubmit = func() {
		fmt.Println("click!")

	}
	return container.NewVBox(form)
}

func validateTitle(text string) error {
	if len(text) == 0 {
		return fmt.Errorf("необходимо ввести значение")
	}

	return nil
}
