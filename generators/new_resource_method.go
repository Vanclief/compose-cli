package generators

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/vanclief/compose-cli/generators/templates"
	"github.com/vanclief/ez"
)

func NewResourceMethod(force bool) error {
	const op = "generators.NewResourceMethod"

	color.Cyan("Generate a new resource method")

	// Check the folder exists
	err := dirExists(RESOURCES_PATH)
	if err != nil {
		errMsg := fmt.Sprintf(`Resources folder "%s" doesn't exist`, RESOURCES_PATH)
		return ez.New(op, ez.ECONFLICT, errMsg, err)
	}

	// Get the directories from the resources folder
	var dirs []string

	err = filepath.Walk(RESOURCES_PATH, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip the root directory itself
		if path == RESOURCES_PATH {
			return nil
		}
		if info.IsDir() {
			dirs = append(dirs, info.Name())
		}
		return nil
	})

	// Print the directories
	color.Cyan("Existing resources:")
	for _, dir := range dirs {
		fmt.Println(dir)
	}

	// Get the resource name
	color.Yellow("Enter resource name: ")
	reader := bufio.NewReader(os.Stdin)
	resourceName, err := reader.ReadString('\n')
	if err != nil {
		return ez.New(op, ez.EINTERNAL, "Error reading input", err)
	}

	resourceName = strings.ToLower(strings.TrimSpace(resourceName))
	if !contains(dirs, resourceName) {
		return ez.New(op, ez.EINVALID, "Resource name doesn't exist, select one from the list", nil)
	}

	// Get the method
	methods := []string{"list", "get", "create", "update", "delete"}

	color.Yellow("Enter resource method: ")
	reader = bufio.NewReader(os.Stdin)
	resourceMethod, err := reader.ReadString('\n')
	if err != nil {
		return ez.New(op, ez.EINTERNAL, "Error reading input", err)
	}
	resourceMethod = strings.ToLower(strings.TrimSpace(resourceMethod))

	var templateName string
	if !contains(methods, resourceMethod) {
		templateName = "generic"
	} else {
		templateName = resourceMethod
	}

	// Create the file API File
	modulePath, err := getModulePath()
	if err != nil {
		return ez.Wrap(op, err)
	}

	color.Cyan("Generating %s method for %s\n", resourceMethod, resourceName)
	filePath := fmt.Sprintf("%s/%s/%s.go", RESOURCES_PATH, resourceName, resourceMethod)

	methodData := templates.MethodData{
		PackageName:   resourceName,
		ModulePath:    modulePath,
		ModelPackage:  resourceName,
		ModelStruct:   uppercaseFirst(singularize(resourceName)),
		ModelVariable: singularize(resourceName),
		ModelSlice:    uppercaseFirst(resourceName),
		MethodName:    uppercaseFirst(resourceMethod),
	}

	templatePath := fmt.Sprintf("%s.go.tpl", templateName)
	err = createFileFromTemplate(filePath, templatePath, methodData, force)
	if err != nil {
		return ez.Wrap(op, err)
	}

	// Create the file API Test File
	filePath = fmt.Sprintf("%s/%s/%s_test.go", RESOURCES_PATH, resourceName, resourceMethod)

	apiTestData := templates.MethodTestData{
		PackageName: resourceName,
		ModulePath:  modulePath,
		SuiteName:   uppercaseFirst(resourceName) + "Suite",
		TestFunc:    methodData.MethodName + "Suite",
	}

	templatePath = "generic_test.go.tpl"
	err = createFileFromTemplate(filePath, templatePath, apiTestData, force)
	if err != nil {
		return ez.Wrap(op, err)
	}

	return nil
}
