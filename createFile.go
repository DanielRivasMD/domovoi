////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"os"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// CreateFile returns a NotFoundAction that attempts to create a file if it doesn't exist.
// If the file is successfully created, the action returns (true, nil).
// Otherwise, it returns (false, error) with detailed diagnostic information.
func CreateFile(filePath string, verbose bool) horus.NotFoundAction {
	return func(address string) (bool, error) {
		if verbose {
			fmt.Printf("Attempting to create file: %s\n", address)
		}
		file, err := os.Create(address)
		if err != nil {
			return false, horus.NewCategorizedHerror(
				"create file",
				"file_creation_error",
				"failed to create file",
				err,
				map[string]any{"path": address},
			)
		}
		// Ensure the file is closed after creation.
		if err := file.Close(); err != nil {
			return false, horus.NewCategorizedHerror(
				"create file",
				"file_closing_error",
				"failed to close file after creation",
				err,
				map[string]any{"path": address},
			)
		}
		if verbose {
			fmt.Printf("File successfully created: %s\n", address)
		}
		return true, nil
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
