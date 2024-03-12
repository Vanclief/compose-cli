package generators

import (
	"os"

	"github.com/vanclief/ez"
)

func folderExists(path string) error {
	const op = "folderExists"

	if info, err := os.Stat(path); err == nil && info.IsDir() {
		return nil
	} else if os.IsNotExist(err) {
		return ez.New(op, ez.EINVALID, "Folder does not exist.", nil)
	} else {
		return ez.New(op, ez.EINTERNAL, "Error checking folder.", err)
	}
}
