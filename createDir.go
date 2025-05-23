////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"os"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// CreateDir returns a NotFoundAction that attempts to create a directory if it doesn't exist.
// If the directory is successfully created, the action returns (true, nil). If it fails,
// it returns (false, error) with diagnostic details.
func CreateDir(dirPath string) horus.NotFoundAction {
	return func(address string) (bool, error) {
		fmt.Printf("Attempting to create directory: %s\n", address)
		err := os.Mkdir(address, 0755)
		if err != nil {
			return false, horus.NewCategorizedHerror(
				"create directory",
				"directory_creation_error",
				"failed to create directory",
				err,
				map[string]any{"path": address},
			)
		}
		fmt.Printf("Directory successfully created: %s\n", address)
		return true, nil
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
