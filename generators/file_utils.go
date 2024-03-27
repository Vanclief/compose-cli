package generators

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

func dirExists(path string) error {
	const op = "dirExists"

	if info, err := os.Stat(path); err == nil && info.IsDir() {
		return nil
	} else if os.IsNotExist(err) {
		return ez.New(op, ez.EINVALID, "Folder does not exist.", nil)
	} else {
		return ez.New(op, ez.EINTERNAL, "Error checking folder.", err)
	}
}

func createDir(path string) error {
	const op = "createDir"

	err := dirExists(path)
	if err == nil {
		errMsg := fmt.Sprintf(`%s already exists`, path)
		return ez.New(op, ez.ECONFLICT, errMsg, err)
	}

	err = os.MkdirAll(path, 0o755)
	if err != nil {
		errMsg := fmt.Sprintf("Error creating directory:", err)
		return ez.New(op, ez.ECONFLICT, errMsg, err)
	}

	return nil
}
