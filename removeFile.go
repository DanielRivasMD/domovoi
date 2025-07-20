////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"os"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// RemoveFile returns a NotFoundAction that attempts to delete a file if it exists.
// If the file is successfully removed (or already missing), it returns (true, nil).
// Otherwise, it returns (false, error) with detailed diagnostic information.
func RemoveFile(filePath string, verbose bool) horus.NotFoundAction {
	return func(address string) (bool, error) {
		if verbose {
			fmt.Printf("Attempting to remove file: %s\n", address)
		}
		err := os.Remove(address)
		if err != nil {
			if os.IsNotExist(err) {
				// nothing to remove
				fmt.Printf("File not found, nothing to remove: %s\n", address)
				return true, nil
			}
			return false, horus.NewCategorizedHerror(
				"remove file",
				"file_removal_error",
				"failed to remove file",
				err,
				map[string]any{"path": address},
			)
		}
		if verbose {
			fmt.Printf("File successfully removed: %s\n", address)
		}
		return true, nil
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
