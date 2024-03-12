package generators

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/fatih/color"
	"github.com/vanclief/compose-cli/generators/templates"
)

const RESOURCES_PATH = "application/resources"

func NewResourceAPI() error {
	color.Cyan("Generating a new resource")

	// Get the module path
	modulePath, err := getModulePath()
	if err != nil {
		color.Red(err.Error())
		return err
	}

	// Check the folder exists
	err = folderExists(RESOURCES_PATH)
	if err != nil {
		errMsg := fmt.Sprintf(`Resources folder "%s" doesn't exist`, RESOURCES_PATH)
		color.Red(errMsg)
		return err
	}

	// Get the resource name
	color.Yellow("Resource name (plural): ")
	reader := bufio.NewReader(os.Stdin)
	resourceName, err := reader.ReadString('\n')
	if err != nil {
		color.Red("Error reading input")
		return err
	}

	resourceName = strings.ToLower(strings.TrimSpace(resourceName))
	color.Cyan("Generating new resource for %s\n", resourceName)

	// Create the resource folder
	dirPath := fmt.Sprintf("%s/%s", RESOURCES_PATH, resourceName)

	err = folderExists(dirPath)
	if err == nil {
		errMsg := fmt.Sprintf(`%s already exists`, dirPath)
		color.Red(errMsg)
		return err
	}

	err = os.Mkdir(dirPath, 0755)
	if err != nil {
		errMsg := fmt.Sprintf("Error creating directory:", err)
		color.Red(errMsg)
		return err
	}

	// Create the file
	tmpl, err := template.ParseFS(templates.FS, "api.go.tpl")
	if err != nil {
		errMsg := fmt.Sprintf("Error parsing template: %v", err)
		color.Red(errMsg)
		return err
	}

	filePath := fmt.Sprintf("%s/%s/%s.go", RESOURCES_PATH, resourceName, "api")
	f, err := os.Create(filePath)
	if err != nil {
		errMsg := fmt.Sprintf("Error creating file: %v", err)
		color.Red(errMsg)
		return err
	}

	defer f.Close()

	apiData := templates.APIData{
		PackageName: resourceName,
		ModulePath:  modulePath,
	}

	err = tmpl.Execute(f, apiData)
	if err != nil {
		errMsg := fmt.Sprintf("Error executing template: %v", err)
		color.Red(errMsg)
		return err
	}

	return nil
}

func NewAPIMethod() error {
	color.Green("Generating api method...")
	return nil
}
