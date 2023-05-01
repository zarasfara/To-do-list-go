package file

import (
	"encoding/json"
	"fmt"
	"github.com/zarasfara/to-do-list/pkg/models"
	"io"
	"os"
)

func ReadTasksFromFile() ([]models.Task, error) {
	file, err := os.Open("tasks.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var tasks []models.Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		panic(err)
	}

	return tasks, nil
}

func WriteTaskToFile(title, description, category string) error {

	// Получаем следующее значение id
	nextId, err := getNextId()
	if err != nil {
		fmt.Errorf("ошибка: %s", err)
	}

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

	var tasks []*models.Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil && err != io.EOF {
		return err
	}

	// Добавляем новый объект в конец массива
	tasks = append(tasks, todo)

	// Перезаписываем содержимое файла
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}
	err = file.Truncate(0)
	if err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(&tasks)
	if err != nil {
		return err
	}

	return nil
}

func GetTaskById(id int) (*models.Task, error) {
	tasks, err := ReadTasksFromFile()

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(tasks); i++ {
		if id == tasks[i].Id {
			return &tasks[i], nil
		}
	}

	return nil, fmt.Errorf("task with ID %d not found", id)
}

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

func getNextId() (int, error) {
	// Считываем содержимое файла с сохраненными todo
	file, err := os.Open("tasks.json")
	if err != nil {
		return 0, err
	}

	data, _ := io.ReadAll(file)

	// Распарсиваем содержимое файла из JSON-формата в массив todoList
	var tasks []*models.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return 0, err
	}

	if len(tasks) == 0 {
		return 1, nil
	}

	// Ищем максимальное значение id среди всех todo в срезе
	var maxId int
	for _, todo := range tasks {
		if todo.Id > maxId {
			maxId = todo.Id
		}
	}

	// Возвращаем максимальное значение id плюс один
	return maxId + 1, nil
}
