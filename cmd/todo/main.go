package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var cmd string
	var args []string

	if len(os.Args) > 1 {
		cmd = os.Args[1]
		args = os.Args[2:]
		fmt.Println(cmd, args)
	} else {
		log.Println("Ошибка: не введена команда")
		os.Exit(1)
	}

	switch cmd {
	case "add":
		listCmd := flag.NewFlagSet("add", flag.ExitOnError)
		desc := listCmd.String("desc", "", "task description")
		err := listCmd.Parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(*desc)
	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		filter := listCmd.String("filter", "all", "all, done, pending")
		err := listCmd.Parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(*filter)
	case "complete":
		listCmd := flag.NewFlagSet("complete", flag.ExitOnError)
		id := listCmd.Int("complete", 0, "task id")
		err := listCmd.Parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(*id)
	case "delete":
		listCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := listCmd.Int("filter", 0, "task id")
		err := listCmd.Parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(*id)
	case "load":
		listCmd := flag.NewFlagSet("load", flag.ExitOnError)
		load := listCmd.String("load", "", "file name")
		err := listCmd.Parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(*load)
	case "export":
		listCmd := flag.NewFlagSet("export", flag.ExitOnError)
		export := listCmd.String("export", "", "file name")
		err := listCmd.Parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(*export)
	default:
		fmt.Println("Введенной команды не существует")
	}
}
