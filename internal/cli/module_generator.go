package cli

import (
	"fmt"
	"os"
	"strings"
)

func MakeModule(name string) {

	moduleName := strings.ToLower(name)

	basePath := "internal/modules/" + moduleName

	os.MkdirAll(basePath, os.ModePerm)

	createFile(basePath+"/model.go", modelTemplate(moduleName))
	createFile(basePath+"/repository.go", repoTemplate(moduleName))
	createFile(basePath+"/service.go", serviceTemplate(moduleName))
	createFile(basePath+"/handler.go", handlerTemplate(moduleName))

	fmt.Println("module created:", moduleName)
}

func createFile(path string, content string) {

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(content)
}