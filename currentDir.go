////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"path/filepath"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// CurrentDir returns the base name (last component) of the current directory.
// It obtains the absolute path of the current directory and then extracts its base name.
// If an error occurs while determining the absolute path, it returns a wrapped error using Horus.
func CurrentDir() (string, error) {
	currentDir, err := filepath.Abs(".")
	if err != nil {
		return "", horus.NewCategorizedHerror(
			"get current directory",
			"path_error",
			"failed to get absolute path of current directory",
			err,
			map[string]any{"function": "CurrentDirBase"},
		)
	}
	baseName := filepath.Base(currentDir)
	return baseName, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
