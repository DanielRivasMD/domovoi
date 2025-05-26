////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"os"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// EnsureDirExist uses domovoi.DirExist to check whether the directory at dirPath exists.
// If it does not exist, the anonymous NotFoundAction attempts to create it.
// The verbose flag controls whether diagnostic messages are printed.
func EnsureDirExist(dirPath string, verbose bool) error {
	exists, err := DirExist(dirPath, func(path string) (bool, error) {
		// Attempt to create the directory (and any necessary parent directories)
		if err := os.MkdirAll(path, 0755); err != nil {
			return false, err
		}
		return true, nil
	}, verbose)
	if err != nil {
		return err
	}
	if !exists {
		// If exists is false even after our notFoundAction, we can wrap and return an error.
		return fmt.Errorf("directory %q does not exist after creation attempt", dirPath)
	}
	return nil
}

func main() {
	dirPath := "./my_directory"
	// Ensure the directory exists, creating it if not.
	if err := EnsureDirExist(dirPath, true); err != nil {
		// Handle or log the error.
		horus.CheckErr(err)
	}

	fmt.Printf("Directory %q is ready to use.\n", dirPath)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

