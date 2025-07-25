////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"os"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// CreateDir ensures that the directory at `dir` exists. If `verbose` is true,
// it prints diagnostic messages before and after checking / creating.
// Any errors are wrapped and propagated via horus.
func CreateDir(dir string, verbose bool) error {
	if verbose {
		fmt.Printf("Ensuring directory exists: %s\n", dir)
	}

	exists, err := DirExist(dir, func(path string) (bool, error) {
		if verbose {
			fmt.Printf("Directory not found; creating: %s\n", path)
		}
		if mkErr := os.MkdirAll(path, 0755); mkErr != nil {
			return false, mkErr
		}
		return true, nil
	}, verbose)
	if err != nil {
		return horus.PropagateErr(
			"CreateDir",
			"dir_exist_error",
			"failed to verify or create directory",
			err,
			map[string]any{"directory": dir},
		)
	}

	if !exists {
		return horus.NewCategorizedHerror(
			"CreateDir",
			"dir_exist_error",
			fmt.Sprintf("directory %q does not exist after creation attempt", dir),
			nil,
			map[string]any{"directory": dir},
		)
	}

	if verbose {
		fmt.Printf("Directory is present: %s\n", dir)
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
