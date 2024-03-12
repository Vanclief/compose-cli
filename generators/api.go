package generators

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func NewResourceAPI() error {
	path := "application/resources"

	color.White("Generating a new resource")

	err := folderExists(path)
	if err != nil {
		errMsg := fmt.Sprintf(`Resources folder "%s" doesn't exist`, path)
		color.Red(errMsg)
		return err
	}

	color.White("Resource name (plural): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	color.White("Generating new resource for %s\n", input)

	return nil
}

func NewAPIMethod() error {
	color.Green("Generating api method...")
	return nil
}
