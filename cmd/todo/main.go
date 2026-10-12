package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"todo-app/internal/storage"
	"todo-app/internal/todo"
)

func main() {
	var cmd string
	var args []string

	if len(os.Args) > 1 {
		cmd = os.Args[1]
		args = os.Args[2:]
	} else {
		log.Println("Ошибка: не введена команда")
		os.Exit(1)
	}

	currentTasks, err := storage.LoadJSON("./test.json")
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		os.Exit(1)
	}

	switch cmd {
	case "add":
		listCmd := flag.NewFlagSet("add", flag.ExitOnError)
		desc := listCmd.String("desc", "", "task description")
		err := listCmd.Parse(args)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			os.Exit(1)
		}
		currentTasks = todo.Add(currentTasks, *desc)
	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		filter := listCmd.String("filter", "all", "all, done, pending")
		err := listCmd.Parse(args)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			os.Exit(1)
		}
		list := todo.List(currentTasks, *filter)
		for _, task := range list {
			var status string
			if task.Done {
				status = "Выполнена"
			} else {
				status = "Не выполнена"
			}
			fmt.Printf("ID: %d / %s - %s\n", task.ID, task.Description, status)
		}
	case "complete":
		listCmd := flag.NewFlagSet("complete", flag.ExitOnError)
		id := listCmd.Int("id", 0, "task id")
		err := listCmd.Parse(args)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			os.Exit(1)
		}
		currentTasks, err = todo.Complete(currentTasks, *id)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			os.Exit(1)
		}
	case "delete":
		listCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := listCmd.Int("id", 0, "task id")
		err := listCmd.Parse(args)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			os.Exit(1)
		}
		currentTasks, err = todo.Delete(currentTasks, *id)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			os.Exit(1)
		}
	case "load":
		listCmd := flag.NewFlagSet("load", flag.ExitOnError)
		file := listCmd.String("file", "", "file name")
		err := listCmd.Parse(args)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			os.Exit(1)
		}
		if strings.HasSuffix(*file, "csv") {
			currentTasks, err = storage.LoadCSV(*file)
			if err != nil {
				fmt.Printf("Ошибка: %s\n", err)
				os.Exit(1)
			}
		} else if strings.HasSuffix(*file, "json") {
			currentTasks, err = storage.LoadJSON(*file)
			if err != nil {
				fmt.Printf("Ошибка: %s\n", err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("Ошибка: неверно указан формат")
			os.Exit(1)
		}
	case "export":
		listCmd := flag.NewFlagSet("export", flag.ExitOnError)
		format := listCmd.String("format", "json", "csv or json")
		out := listCmd.String("out", "", "file name without format")
		err := listCmd.Parse(args)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			os.Exit(1)
		}
		switch *format {
		case "json":
			err = storage.SaveJSON(fmt.Sprintf("%s.%s", *out, *format), currentTasks)
			if err != nil {
				fmt.Printf("Ошибка: %s\n", err)
				os.Exit(1)
			}
		case "csv":
			err = storage.SaveCSV(fmt.Sprintf("%s.%s", *out, *format), currentTasks)
			if err != nil {
				fmt.Printf("Ошибка: %s\n", err)
				os.Exit(1)
			}
		default:
			fmt.Printf("Ошибка: неверно указан формат")
			os.Exit(1)
		}
	default:
		fmt.Println("Ошибка: введенной команды не существует")
		os.Exit(1)
	}

	err = storage.SaveJSON("test.json", currentTasks)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		os.Exit(1)
	}

}
