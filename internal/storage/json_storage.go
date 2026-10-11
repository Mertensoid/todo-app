package storage

import (
	"encoding/json"
	"errors"
	"os"
	"todo-app/internal/todo"
)

func LoadJSON(path string) ([]todo.Task, error) {
	_, err := os.Stat(path)
	if err != nil {
		SaveJSON(path, []todo.Task{})
		return []todo.Task{}, errors.New("указанного файла не существует")
	}
	var file []byte
	file, err = os.ReadFile(path)
	if err != nil {
		return []todo.Task{}, errors.New("чтение файла не удалось")
	}
	var result []todo.Task
	err = json.Unmarshal(file, &result)
	return result, err
}

func SaveJSON(path string, tasks []todo.Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return errors.New("сериализация json не удалась")
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return errors.New("запись json в файл не удалась")
	}
	return nil
}
