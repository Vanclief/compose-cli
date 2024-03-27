package generators

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

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

	err = os.Mkdir(path, 0o755)
	if err != nil {
		errMsg := fmt.Sprintf("Error creating directory:", err)
		return ez.New(op, ez.ECONFLICT, errMsg, err)
	}

	return nil
}

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

	return nil
}

// contains checks if a string is present in a slice of strings.
func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

// singularize attempts to convert a plural noun to its singular form.
// This is a basic implementation and may not correctly handle all plural forms.
func singularize(word string) string {
	// Basic rules for regular plural forms
	if strings.HasSuffix(word, "ies") {
		return strings.TrimSuffix(word, "ies") + "y"
	} else if strings.HasSuffix(word, "ves") {
		return strings.TrimSuffix(word, "ves") + "f"
	} else if strings.HasSuffix(word, "s") {
		// Assumes words ending in 's' are plural
		return strings.TrimSuffix(word, "s")
	}

	// Return the word if no rules apply
	return word
}

// uppercaseFirst converts the first letter of the string to uppercase.
func uppercaseFirst(s string) string {
	if s == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}
