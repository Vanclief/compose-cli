package generators

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/vanclief/ez"
)

func getUserConfirmation(prompt string) (bool, error) {
	const op = "getUserConfirmation"
	reader := bufio.NewReader(os.Stdin)

	for {
		msg := fmt.Sprintf("%s (Y/N): ", prompt)
		color.Yellow(msg)
		input, err := reader.ReadString('\n')
		if err != nil {
			return false, ez.New(op, ez.EINTERNAL, "Error reading input", err)
		}

		input = strings.ToLower(strings.TrimSpace(input))

		switch input {
		case "yes":
			fallthrough
		case "y":
			return true, nil
		case "no":
			fallthrough
		case "n":
			return false, nil
		default:
			color.Yellow("Do you want to proceed? (Y/N): ")
		}
	}
}

func getUserInput(prompt string) (string, error) {
	const op = "getUserInput"

	color.Yellow(fmt.Sprintf("%s: ", prompt))

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", ez.New(op, ez.EINTERNAL, "Error reading user input", err)
	}

	input = strings.ToLower(strings.TrimSpace(input))

	return input, nil
}
