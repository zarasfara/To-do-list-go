package file

import (
	"encoding/json"
	"fmt"
	"github.com/zarasfara/to-do-list/pkg/models"
	"io"
	"os"
)

/*
Parameters:

	filePath (string): Имя файла

Returns:

	[]models.Task: Срез структур задач
*/
func ReadTasksFromFile(filePath string) ([]models.Task, error) {
	file, err := os.Open(filePath)
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

func WriteTasksToFile(title, description, category string) error {

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

/*
*
Получаем id следующей структуры

Return:

	(int)id: айдшник следующей записи
*/
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
