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

const RESOURCES_PATH = "application/resources"

func NewResourceAPI() error {
	const op = "generators.NewResourceAPI"

	color.Cyan("Generating a new resource")

	// Get the module path
	modulePath, err := getModulePath()
	if err != nil {
		return err
	}

	// Check the folder exists
	err = folderExists(RESOURCES_PATH)
	if err != nil {
		errMsg := fmt.Sprintf(`Resources folder "%s" doesn't exist`, RESOURCES_PATH)
		return ez.New(op, ez.ECONFLICT, errMsg, err)
	}

	// Get the resource name
	color.Yellow("Resource name (plural): ")
	reader := bufio.NewReader(os.Stdin)
	resourceName, err := reader.ReadString('\n')
	if err != nil {
		return ez.New(op, ez.EINTERNAL, "Error reading input", err)
	}

	resourceName = strings.ToLower(strings.TrimSpace(resourceName))
	color.Cyan("Generating new resource for %s\n", resourceName)

	// Create the resource folder
	dirPath := fmt.Sprintf("%s/%s", RESOURCES_PATH, resourceName)

	err = folderExists(dirPath)
	if err == nil {
		errMsg := fmt.Sprintf(`%s already exists`, dirPath)
		return ez.New(op, ez.ECONFLICT, errMsg, err)
	}

	err = os.Mkdir(dirPath, 0o755)
	if err != nil {
		errMsg := fmt.Sprintf("Error creating directory:", err)
		return ez.New(op, ez.ECONFLICT, errMsg, err)
	}

	// Create the file API File
	filePath := fmt.Sprintf("%s/%s/api.go", RESOURCES_PATH, resourceName)

	apiData := templates.APIData{
		PackageName: resourceName,
		ModulePath:  modulePath,
	}

	err = createFileFromTemplate(filePath, "api.go.tpl", apiData)
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

	err = createFileFromTemplate(filePath, "api_test.go.tpl", apiTestData)
	if err != nil {
		return ez.Wrap(op, err)
	}

	return nil
}

func NewResourceMethod() error {
	color.Green("Generating api method...")
	return nil
}
