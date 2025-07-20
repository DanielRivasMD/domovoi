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
// a list of directory entries.
// It prints diagnostic messages and wraps errors with detailed context.
func ReadDir(path string) ([]os.DirEntry, error) {
	fmt.Printf("Attempting to read directory: %s\n", path)
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
	fmt.Printf("Directory successfully read: %s (%d entries)\n", path, len(entries))
	return entries, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
