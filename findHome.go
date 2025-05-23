////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"

	"github.com/DanielRivasMD/horus"
	"github.com/atrox/homedir"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// FindHome attempts to retrieve the current user's home directory.
// It prints diagnostic messages and returns the home directory string.
// On failure, it returns a wrapped error with detailed diagnostic information.
func FindHome() (string, error) {
	fmt.Printf("Attempting to find home directory...\n")
	home, err := homedir.Dir()
	if err != nil {
		return "", horus.NewCategorizedHerror(
			"find home",
			"home_directory_error",
			"failed to find home directory",
			err,
			map[string]any{"function": "FindHome"},
		)
	}
	fmt.Printf("Home directory successfully found: %s\n", home)
	return home, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
