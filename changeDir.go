////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"os"

	"github.com/DanielRivasMD/horus"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// ChangeDir changes the current working directory to the specified path and reports errors using horus.
func ChangeDir(newDir string) error {
	err := os.Chdir(newDir)
	if err != nil {
		// Create a categorized error using horus with a colored message.
		herr := horus.NewCategorizedHerror(
			"change directory",
			"directory_error",
			chalk.Red.Color(fmt.Sprintf("failed to change directory to '%s'", newDir)),
			err,
			map[string]any{"target_directory": newDir},
		)

		// Optionally, log the error for debugging.
		fmt.Println(horus.FormatError(herr, horus.JSONFormatter))
		return herr
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
