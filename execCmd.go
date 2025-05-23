////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"os"
	"os/exec"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// ExecCmd executes the specified command with the provided arguments.
// It configures the command to forward its output (Stdout and Stderr)
// to the operating system's standard streams, preserving any color codes.
// If executing the command fails, ExecCmd returns a wrapped error using Horus.
func ExecCmd(command string, args ...string) error {
	cmdRun := exec.Command(command, args...)

	// Forward command output to the console.
	cmdRun.Stdout = os.Stdout
	cmdRun.Stderr = os.Stderr

	// Run the command.
	err := cmdRun.Run()
	if err != nil {
		return horus.NewCategorizedHerror(
			"exec command",
			"command_execution_error",
			"failed to execute command",
			err,
			map[string]any{
				"command": command,
				"args":    args,
			},
		)
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
