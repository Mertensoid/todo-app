package todo

import "errors"

func Add(tasks []Task, desc string) []Task {
	return append(tasks, Task{
		tasks[len(tasks)-1].ID + 1,
		desc,
		false,
	})
}

func List(tasks []Task, filter string) []Task {
	switch filter {
	case "done":
		filtered := []Task{}
		for _, task := range tasks {
			if task.Done {
				filtered = append(filtered, task)
			}
		}
		return filtered
	case "pending":
		filtered := []Task{}
		for _, task := range tasks {
			if !task.Done {
				filtered = append(filtered, task)
			}
		}
		return filtered
	default:
		return tasks
	}
}

func Complete(tasks []Task, id int) ([]Task, error) {
	result := tasks
	for i, task := range tasks {
		if task.ID == id {
			if task.Done {
				return result, errors.New("задача уже выполнена")
			}
			result[i].Done = true
			return result, nil
		}
	}
	return result, errors.New("задача не найдена")
}

func Delete(tasks []Task, id int) ([]Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...), nil
		}
	}
	return tasks, errors.New("задача не найдена")
}
