////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"strings"

	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// PrintCentered prints msg centered in a 100-column width,
// surrounding it with “=” fills and one space padding on each side.
func PrintCentered(msg string) {
	const width = 100
	// total width occupied by msg + two spaces
	padded := len(msg) + 2
	// remaining fill slots
	fill := width - padded
	if fill <= 0 {
		// msg too long—just print it
		fmt.Println(chalk.Blue.Color(msg))
		return
	}
	// distribute fill equally (extra goes on right)
	left := fill / 2
	right := fill - left

	line := strings.Repeat("=", left) +
		" " + msg + " " +
		strings.Repeat("=", right)

	fmt.Println(chalk.Blue.Color(line))
}

////////////////////////////////////////////////////////////////////////////////////////////////////
