package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/zarasfara/to-do-list/internal/file"
	"github.com/zarasfara/to-do-list/pkg/models"
	"strconv"
)

var TaskId int

var tasks []models.Task

type TaskTable struct {
	widget.Table
}

func (t *TaskTable) RefreshTable() {
	tempTasks, err := file.ReadTasksFromFile()
	if err != nil {
		return
	}

	tasks = tempTasks

	t.Refresh()
}

func NewTasksTable() *TaskTable {
	table := &TaskTable{}
	table.ExtendBaseWidget(table)

	// Создаем заголовки таблицы
	headers := [...]string{"id", "Название", "Описание", "Категория", "Статус"}

	tasks, _ = file.ReadTasksFromFile()

	table.Length = func() (int, int) {
		return len(tasks) + 1, 5 // 5 строки, 5 столбца
	}
	table.CreateCell = func() fyne.CanvasObject {
		return widget.NewLabel("Cell")
	}
	table.UpdateCell = func(cell widget.TableCellID, cellView fyne.CanvasObject) {

		if cell.Row == 0 { // Если первая строка выводим данные из массива заголовков
			cellView.(*widget.Label).SetText(headers[cell.Col])

		} else if cell.Row <= len(tasks) { // выбранная строка в таблице находится в диапазоне допустимых индексов

			// Устанавливаем значения ячеек таблицы, если индекс не выходит за пределы массива
			task := tasks[cell.Row-1]
			switch cell.Col {
			case 0:
				cellView.(*widget.Label).SetText(strconv.Itoa(task.Id))
			case 1:
				cellView.(*widget.Label).SetText(task.Title)
			case 2:
				cellView.(*widget.Label).SetText(task.Description)
			case 3:
				cellView.(*widget.Label).SetText(task.Category)
			case 4:
				var text string
				if task.Completed {
					text = "Завершено"
				} else {
					text = "В работе"
				}
				cellView.(*widget.Label).SetText(text)
			}
		}
	}

	table.OnSelected = func(id widget.TableCellID) {
		if id.Row == 0 { // Если выбрана первая строка, то ничего не делаем
			return
		}

		TaskId = tasks[id.Row-1].Id
	}

	//columnWidth := 1000 / len(headers) // вычисляем ширину колонки
	//for i := 0; i < len(headers); i++ {
	//	table.SetColumnWidth(i, float32(columnWidth))
	//}

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 250)
	table.SetColumnWidth(2, 450)
	table.SetColumnWidth(3, 100)
	table.SetColumnWidth(4, 150)

	return table
}