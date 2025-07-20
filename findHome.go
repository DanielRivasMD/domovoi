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
// If verbose is true it prints diagnostics; otherwise it's silent.
// On error it returns a categorized *Herror.
func FindHome(verbose bool) (string, error) {
	if verbose {
		fmt.Println("Attempting to find home directory...")
	}

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

	if verbose {
		fmt.Printf("Home directory successfully found: %s\n", home)
	}
	return home, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
