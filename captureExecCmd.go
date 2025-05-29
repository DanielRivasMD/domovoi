////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"os/exec"

	"github.com/DanielRivasMD/horus"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// CaptureExecCmd executes the specified command with provided arguments and
// captures its standard output and error output into separate strings.
// If an error occurs during execution, it returns a Horus-wrapped error with detailed context.
func CaptureExecCmd(command string, args ...string) (string, string, error) {
	// Prepare the command.
	cmdRun := exec.Command(command, args...)

	// Capture output using buffers.
	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	cmdRun.Stdout = &stdoutBuf
	cmdRun.Stderr = &stderrBuf

	// Run the command.
	err := cmdRun.Run()
	if err != nil {
		return "", "", horus.NewCategorizedHerror(
			"capture exec command",
			"command_capture_error",
			chalk.Red.Color("failed to execute command and capture output"),
			err,
			map[string]any{
				"command": command,
				"args":    args,
			},
		)
	}

	// Return captured output.
	return stdoutBuf.String(), stderrBuf.String(), nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
