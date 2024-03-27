package generators

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/vanclief/compose-cli/generators/templates"
	"github.com/vanclief/ez"
)

func NewResourceMethod(force bool) error {
	const op = "generators.NewResourceMethod"

	color.Blue("Generate a new resource method")

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
	color.Blue("Existing resources:")
	for i, dir := range dirs {
		line := fmt.Sprintf("(%d) %s", i+1, dir)
		fmt.Println(line)
	}

	// Get the resource name
	resourceName, err := getUserInput("Enter resource name from list")
	if err != nil {
		return ez.New(op, ez.EINTERNAL, "Error reading input", err)
	}

	if !contains(dirs, resourceName) {
		errMsg := fmt.Sprintf("Resource %s does not exist in the list, create it first")
		return ez.New(op, ez.EINVALID, errMsg, nil)
	}

	// Get the method
	methods := []string{"list", "get", "create", "update", "delete"}

	resourceMethod, err := getUserInput("Enter resource method (i.e. create)")
	if err != nil {
		return ez.New(op, ez.EINTERNAL, "Error reading input", err)
	}

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

	color.Blue("Generating %s method for %s\n", resourceMethod, resourceName)
	filePath := fmt.Sprintf("%s/%s/%s.go", RESOURCES_PATH, resourceName, resourceMethod)

	methodData := templates.MethodData{
		PackageName:   resourceName,
		ModulePath:    modulePath,
		ModelStruct:   uppercaseFirst(singularize(resourceName)),
		ModelVariable: singularize(resourceName),
		ModelSlice:    uppercaseFirst(resourceName),
		MethodName:    uppercaseFirst(resourceMethod),
	}

	templatePath := fmt.Sprintf("methods/%s.go.tpl", templateName)
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

	templatePath = "methods/generic_test.go.tpl"
	err = createFileFromTemplate(filePath, templatePath, apiTestData, force)
	if err != nil {
		return ez.Wrap(op, err)
	}

	// Check if the user wants to create a handler
	createHandler, err := getUserConfirmation("Do you want to create a handler?")
	if err != nil {
		return ez.Wrap(op, err)
	}

	if !createHandler {
		return nil
	}

	color.Blue("Generating %s handler for %s\n", resourceMethod, resourceName)

	err = dirExists(REST_HANDLER_PATH)
	if err != nil {
		errMsg := fmt.Sprintf(`Handler folder "%s" doesn't exist do you want to create it?`, REST_HANDLER_PATH)

		createDirectory, errConf := getUserConfirmation(errMsg)
		if errConf != nil || !createDirectory {
			return ez.New(op, ez.ECONFLICT, errMsg, err)
		}

		err = createDir(REST_HANDLER_PATH)
		if err != nil {
			return ez.New(op, ez.ECONFLICT, errMsg, err)
		}
	}

	filePath = fmt.Sprintf("%s/%s.go", REST_HANDLER_PATH, resourceName)

	// Check if the file already exists
	fileExists := true
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			fileExists = false
		}
	}

	if !fileExists {
		templatePath = "handlers/imports.go.tpl"
		err = createFileFromTemplate(filePath, templatePath, methodData, false)
		if err != nil {
			return ez.Wrap(op, err)
		}
	}

	templatePath = fmt.Sprintf("handlers/%s.go.tpl", templateName)
	err = appendFileFromTemplate(filePath, templatePath, methodData)
	if err != nil {
		return ez.Wrap(op, err)
	}

	color.Green("Succesfully generated method %s for %s\n", resourceMethod, resourceName)

	return nil
}
