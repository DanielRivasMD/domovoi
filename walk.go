////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"io/fs"
	"path/filepath"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// fileProcessor defines the type for a file processing function.
type fileProcessor func(string)

// walk returns an fs.WalkDirFunc that applies the given processor function to each non-directory file.
// Hidden files are ignored, and any encountered errors are wrapped and propagated using Horus.
func walk(processor fileProcessor) fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		// If an error occurred while walking this path, wrap and return it.
		if err != nil {
			return horus.PropagateErr(
				"walk",
				"walk_error",
				"Error encountered during directory walk",
				err,
				map[string]any{"path": path},
			)
		}

		// Only process non-directory entries.
		if !d.IsDir() {
			// Split the path into directory and file name.
			dir, file := filepath.Split(path)

			// Check if the file is hidden.
			hidden, err := isHiddenFile(file)
			if err != nil {
				return horus.PropagateErr(
					"walk",
					"hidden_file_check_error",
					"Failed to determine if file is hidden",
					err,
					map[string]any{"file": file},
				)
			}

			// If the file is not hidden, process it.
			if !hidden {
				processor(filepath.Join(dir, file))
			}
		}
		return nil
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
