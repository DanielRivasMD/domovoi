////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"os"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// ListFiles returns a slice of names for all regular files in the provided directory.
// If the directory cannot be read, it returns a wrapped error using horus.
func ListFiles(directory string) ([]string, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, horus.NewCategorizedHerror(
			"list files",
			"io_error",
			"failed to list files in directory",
			err,
			map[string]any{"directory": directory},
		)
	}

	var files []string
	for _, entry := range entries {
		if entry.Type().IsRegular() {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
