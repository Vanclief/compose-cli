package generators

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/vanclief/compose-cli/generators/templates"
	"github.com/vanclief/ez"
)

func NewResourceModel() error {
	const op = "generators.NewResourceModel"

	color.Blue("Generate resource Model")

	// Check the folder exists
	err := dirExists(MODELS_PATH)
	if err != nil {
		errMsg := fmt.Sprintf(`Models folder "%s" doesn't exist do you want to create it?`, MODELS_PATH)

		createDirectory, errConf := getUserConfirmation(errMsg)
		if errConf != nil || !createDirectory {
			return ez.New(op, ez.ECONFLICT, errMsg, err)
		}

		err = createDir(MODELS_PATH)
		if err != nil {
			return ez.New(op, ez.ECONFLICT, errMsg, err)
		}
	}

	// Get the model name
	modelName, err := getUserInput("Enter model name (plural)")
	if err != nil {
		return ez.New(op, ez.EINTERNAL, "Error reading input", err)
	}

	color.Blue("Generating Model for %s\n", modelName)

	// Create the model File
	filePath := fmt.Sprintf("%s/%s.go", MODELS_PATH, modelName)

	modelData := templates.ModelData{
		PackageName: modelName,
		ModelStruct: uppercaseFirst(singularize(modelName)),
	}

	err = createFileFromTemplate(filePath, "models/model.go.tpl", modelData, false)
	if err != nil {
		return ez.Wrap(op, err)
	}

	color.Green("Succesfully generated model for %s\n", modelName)

	return nil
}
