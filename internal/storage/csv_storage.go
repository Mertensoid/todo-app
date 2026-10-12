package storage

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"todo-app/internal/todo"
)

func LoadCSV(path string) ([]todo.Task, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("файл не найден")
	}
	var lines [][]string
	lines, err = csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, errors.New("не удалось прочитать файл CSV")
	}
	tasks := []todo.Task{}
	for i := 1; i < len(lines); i++ {
		id, err := strconv.Atoi(lines[i][0])
		if err != nil {
			return nil, errors.New("не удалось извлечь значение ID из файла CSV")
		}
		var done bool
		done, err = strconv.ParseBool(lines[i][2])
		if err != nil {
			return nil, errors.New("не удалось извлечь значение Done из файла CSV")
		}
		task := todo.Task{
			ID:          id,
			Description: lines[i][1],
			Done:        done,
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func SaveCSV(path string, tasks []todo.Task) error {

}
