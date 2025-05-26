////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"os"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// CreateDir ensures that the directory at 'dir' exists.
// If it does not exist, it attempts to create it using os.MkdirAll.
// All errors are wrapped and propagated via horus.PropagateErr.
func CreateDir(dir string) error {
	exists, err := DirExist(dir, func(path string) (bool, error) {
		// Attempt to create the directory (and any necessary parent directories).
		if err := os.MkdirAll(path, 0755); err != nil {
			return false, err
		}
		return true, nil
	}, true)
	if err != nil {
		return horus.PropagateErr(
			"CreateDir",
			"dir_exist_error",
			"Failed to verify or create directory",
			err,
			map[string]any{"directory": dir},
		)
	}
	if !exists {
		return fmt.Errorf("directory %q does not exist after creation attempt", dir)
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
