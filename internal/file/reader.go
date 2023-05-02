package file

import (
	"encoding/json"
	"fmt"
	"github.com/zarasfara/to-do-list/internal/models"
	"io"
	"os"
	"sort"
)

// получить все задачи
func ReadTasksFromFile() ([]*models.Task, error) {
	// Открываем файл с задачами
	file, err := os.Open("tasks.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Читаем содержимое файла
	var tasks []*models.Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil && err != io.EOF {
		return nil, err
	}

	return tasks, nil
}

// добавить задачу
func WriteTaskToFile(title, description, category string) error {
	// Получаем следующее значение id
	nextId, err := getNextId()
	if err != nil {
		return fmt.Errorf("ошибка: %s", err)
	}

	// Создаем новую задачу
	todo := &models.Task{
		Id:          nextId,
		Title:       title,
		Description: description,
		Category:    category,
		Completed:   false,
	}

	// Открываем файл с задачами
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Читаем содержимое файла
	var tasks []*models.Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil && err != io.EOF {
		return err
	}

	// Добавляем новую задачу в конец массива
	tasks = append(tasks, todo)

	// Перезаписываем содержимое файла
	file.Seek(0, 0)  // перемещаем позицию указателя чтения/записи в начало файла.
	file.Truncate(0) //  используется для очистки файла перед перезаписью его содержимого.
	if err := json.NewEncoder(file).Encode(&tasks); err != nil {
		return err
	}

	return nil
}

// получить задачу
func GetTaskById(id int) (*models.Task, error) {
	tasks, err := ReadTasksFromFile()

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(tasks); i++ {
		if id == tasks[i].Id {
			return tasks[i], nil
		}
	}

	return nil, fmt.Errorf("task with ID %d not found", id)
}

// удалить задачу
func DeleteTask(id int) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	// Находим индекс задачи по её ID
	index := -1
	for i, task := range tasks {
		if task.Id == id {
			index = i
			break
		}
	}

	if index < 0 {
		return fmt.Errorf("задача с id %d не найдена", id)
	}

	// Удаляем задачу из списка задач
	tasks = append(tasks[:index], tasks[index+1:]...)

	data, _ := json.Marshal(tasks)

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// инкремент id
func getNextId() (int, error) {
	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return 1, nil
		}
		return 0, err
	}

	defer file.Close()

	var tasks []*models.Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil && err != io.EOF {
		return 0, err
	}

	if len(tasks) == 0 {
		return 1, nil
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Id < tasks[j].Id
	})

	return tasks[len(tasks)-1].Id + 1, nil
}

// обновить задачу
func UpdateTask(id int, title, description string, completed bool) (bool, error) {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return false, err
	}

	// Находим задачу по её ID
	task := FindTaskById(tasks, id)
	if task == nil {
		return false, fmt.Errorf("задача с id %d не найдена", id)
	}

	// Обновляем поля задачи
	task.Title = title
	task.Description = description
	task.Completed = completed

	// Перезаписываем все задачи в файл
	data, err := json.Marshal(tasks)
	if err != nil {
		return false, err
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return false, err
	}

	return true, nil
}

// функция для поиска задачи в срезе задач по её ID
func FindTaskById(tasks []*models.Task, id int) *models.Task {
	for _, task := range tasks {
		if task.Id == id {
			return task
		}
	}
	return nil
}

// изменить статус
func ChangeTaskStatus(CurrentTaskId int) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if tasks[i].Id == CurrentTaskId {
			if task.Completed {
				tasks[i].Completed = false
				break
			} else {
				tasks[i].Completed = true
				break
			}
		}
	}

	data, _ := json.Marshal(tasks)

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}
