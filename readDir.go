////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"os"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// ReadDir reads the directory named by path and returns
// a list of directory entries. If verbose is true,
// it prints diagnostic messages before and after the read.
func ReadDir(path string, verbose bool) ([]os.DirEntry, error) {
	if verbose {
		fmt.Printf("Attempting to read directory: %s\n", path)
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, horus.NewCategorizedHerror(
			"read directory",
			"directory_read_error",
			"failed to read directory",
			err,
			map[string]any{"path": path},
		)
	}

	if verbose {
		fmt.Printf("Directory successfully read: %s (%d entries)\n", path, len(entries))
	}
	return entries, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
