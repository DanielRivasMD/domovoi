////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"os"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// FileExist checks if filePath exists. If it does, returns (true, nil).
// If it does not exist and notFoundAction is nil, returns (false, nil).
// If notFoundAction is non-nil, it’s invoked on absence.
// Any real I/O error (other than “not exist”) or action-error bubbles up.
func FileExist(
	filePath string,
	notFoundAction horus.NotFoundAction,
	verbose bool,
) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			if verbose {
				fmt.Printf("File '%s' does not exist.\n", filePath)
			}
			if notFoundAction != nil {
				ok, actionErr := notFoundAction(filePath)
				if actionErr != nil {
					return false, horus.NewHerror(
						"check file",
						fmt.Sprintf("notFoundAction failed for '%s'", filePath),
						actionErr,
						map[string]any{"path": filePath},
					)
				}
				return ok, nil
			}
			// file missing, but we treat that as a non-error
			return false, nil
		}
		// some other Stat error
		return false, horus.NewHerror(
			"check file",
			fmt.Sprintf("error checking '%s'", filePath),
			err,
			map[string]any{"path": filePath},
		)
	}

	if verbose {
		fmt.Printf("File '%s' exists.\n", filePath)
	}
	return true, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
