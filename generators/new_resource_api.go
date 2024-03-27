package generators

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/vanclief/compose-cli/generators/templates"
	"github.com/vanclief/ez"
)

const (
	RESOURCES_PATH    = "application/resources"
	REST_HANDLER_PATH = "interfaces/rest/handler"
)

func NewResourceAPI() error {
	const op = "generators.NewResourceAPI"

	color.Cyan("Generate a new resource")

	// Check the folder exists
	err := dirExists(RESOURCES_PATH)
	if err != nil {
		errMsg := fmt.Sprintf(`Resources folder "%s" doesn't exist`, RESOURCES_PATH)
		return ez.New(op, ez.ECONFLICT, errMsg, err)
	}

	// Get the resource name
	color.Yellow("Enter resource name (plural): ")
	reader := bufio.NewReader(os.Stdin)
	resourceName, err := reader.ReadString('\n')
	if err != nil {
		return ez.New(op, ez.EINTERNAL, "Error reading input", err)
	}

	resourceName = strings.ToLower(strings.TrimSpace(resourceName))
	color.Cyan("Generating new resource for %s\n", resourceName)

	// Create the resource folder
	dirPath := fmt.Sprintf("%s/%s", RESOURCES_PATH, resourceName)

	err = createDir(dirPath)
	if err != nil {
		return ez.Wrap(op, err)
	}

	// Create the file API File
	modulePath, err := getModulePath()
	if err != nil {
		return ez.Wrap(op, err)
	}

	filePath := fmt.Sprintf("%s/%s/api.go", RESOURCES_PATH, resourceName)

	apiData := templates.APIData{
		PackageName: resourceName,
		ModulePath:  modulePath,
	}

	err = createFileFromTemplate(filePath, "api/api.go.tpl", apiData, false)
	if err != nil {
		return ez.Wrap(op, err)
	}

	// Create the file API Test File
	filePath = fmt.Sprintf("%s/%s/api_test.go", RESOURCES_PATH, resourceName)

	apiTestData := templates.APITestData{
		PackageName: resourceName,
		ModulePath:  modulePath,
		SuiteName:   strings.ToUpper(resourceName[:1]) + resourceName[1:] + "Suite",
	}

	err = createFileFromTemplate(filePath, "api/api_test.go.tpl", apiTestData, false)
	if err != nil {
		return ez.Wrap(op, err)
	}

	return nil
}
