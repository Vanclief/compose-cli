package generators

import (
	"fmt"
	"os"
	"text/template"

	"github.com/fatih/color"
	"github.com/vanclief/compose-cli/generators/templates"
	"github.com/vanclief/ez"
)

func createFileFromTemplate(filePath, templatePath string, templateData interface{}, force bool) error {
	const op = "generators.createFileFromTemplate"

	// Check if the file already exists
	_, err := os.Stat(filePath)
	if !force && err == nil {
		errMsg := fmt.Sprintf("File already exists: %s", filePath)
		return ez.New(op, ez.ECONFLICT, errMsg, nil)
	}

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

	color.Green("Created %s\n", filePath)

	return nil
}

func appendFileFromTemplate(filePath, templatePath string, templateData interface{}) error {
	const op = "generators.appendFileFromTemplate"

	// Parse the template
	tmpl, err := template.ParseFS(templates.FS, templatePath)
	if err != nil {
		errMsg := fmt.Sprintf("Error parsing template: %v", err)
		return ez.New(op, ez.EINTERNAL, errMsg, err)
	}

	var f *os.File
	// Open the file in append mode if it exists, create it otherwise.
	f, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		errMsg := fmt.Sprintf("Error opening or creating file: %v", err)
		return ez.New(op, ez.EINTERNAL, errMsg, err)
	}
	defer f.Close()

	// Execute the template, writing to the file.
	err = tmpl.Execute(f, templateData)
	if err != nil {
		errMsg := fmt.Sprintf("Error executing template: %v", err)
		return ez.New(op, ez.EINTERNAL, errMsg, err)
	}

	return nil
}
