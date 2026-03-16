package cli

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

type MigrationData struct {
	Timestamp string
	Name      string
	FuncName  string
}

func MakeMigration(name string) {

	timestamp := time.Now().Format("20060102150405")

	funcName := toPascalCase(name)

	fileName := fmt.Sprintf(
		"internal/database/migrations/%s_%s.go",
		timestamp,
		name,
	)

	data := MigrationData{
		Timestamp: timestamp,
		Name:      name,
		FuncName:  funcName,
	}

	renderMigration("internal/templates/migration.tpl", fileName, data)

	fmt.Println("migration created:", fileName)
}

func renderMigration(templatePath, outputPath string, data MigrationData) {

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("template error:", err)
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("file error:", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Println("execute error:", err)
	}
}