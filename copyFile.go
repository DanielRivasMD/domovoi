////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"io"
	"os"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// CopyFile copies the contents of src to dst.
// If dst exists, it will be overwritten.
func CopyFile(src, dst string) error {
	op := "domovoi.copy_file"

	in, err := os.Open(src)
	if err != nil {
		return horus.PropagateErr(
			op,
			"open_source_error",
			"failed to open source file",
			err,
			map[string]any{"source": src},
		)
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return horus.PropagateErr(
			op,
			"create_dest_error",
			"failed to create destination file",
			err,
			map[string]any{"destination": dst},
		)
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return horus.PropagateErr(
			op,
			"copy_error",
			"failed to copy file contents",
			err,
			map[string]any{
				"source":      src,
				"destination": dst,
			},
		)
	}

	if err := out.Sync(); err != nil {
		return horus.PropagateErr(
			op,
			"sync_error",
			"failed to flush destination file",
			err,
			map[string]any{"destination": dst},
		)
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
