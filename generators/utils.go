package generators

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/vanclief/compose-cli/generators/templates"
	"github.com/vanclief/ez"
)

func getModulePath() (string, error) {
	const op = "getModulePath"

	file, err := os.Open("go.mod")
	if err != nil {
		return "", ez.New(op, ez.EINVALID, "Error opening go.mod", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() { // Read the first line

		firstLine := scanner.Text()
		modulePath := strings.TrimSpace(firstLine)

		if strings.HasPrefix(modulePath, "module") {
			modulePath = strings.TrimPrefix(modulePath, "module")
			modulePath = strings.TrimSpace(modulePath) // Remove any spaces after "module"

			return modulePath, nil
		}
	}

	return "", ez.New(op, ez.EINVALID, "Error reading from go.mod", nil)
}

func folderExists(path string) error {
	const op = "folderExists"

	if info, err := os.Stat(path); err == nil && info.IsDir() {
		return nil
	} else if os.IsNotExist(err) {
		return ez.New(op, ez.EINVALID, "Folder does not exist.", nil)
	} else {
		return ez.New(op, ez.EINTERNAL, "Error checking folder.", err)
	}
}

func createFileFromTemplate(filePath, templatePath string, templateData interface{}) error {
	const op = "generators.createFileFromTemplate"

	// Create the file
	tmpl, err := template.ParseFS(templates.FS, templatePath)
	if err != nil {
		errMsg := fmt.Sprintf("Error parsing template: %v", err)
		return ez.New(op, ez.EINTERNAL, errMsg, err)
	}

	f, err := os.Create(filePath)
	if err != nil {
		errMsg := fmt.Sprintf("Error creating file: %v", err)
		return ez.New(op, ez.EINTERNAL, errMsg, err)
	}

	defer f.Close()

	err = tmpl.Execute(f, templateData)
	if err != nil {
		errMsg := fmt.Sprintf("Error executing template: %v", err)
		return ez.New(op, ez.EINTERNAL, errMsg, err)
	}

	return nil
}
