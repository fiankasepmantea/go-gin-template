package main

import (
	"fmt"
	"os"

	"github.com/fiankasepman/go-gin-template/internal/cli"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("command not found")
		return
	}

	switch os.Args[1] {

	case "make:module":
		if len(os.Args) < 3 {
			fmt.Println("module name required")
			return
		}
		cli.MakeModule(os.Args[2])

	case "make:migration":
		if len(os.Args) < 3 {
			fmt.Println("migration name required")
			return
		}
		cli.MakeMigration(os.Args[2])

	default:
		fmt.Println("unknown command")
	}
}