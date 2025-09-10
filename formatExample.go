////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"strings"

	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// FormatExample builds a multi‐line example block
// each usage is a slice of "tokens": [ command, flagOrArg, flagOrArg, ... ].
func FormatExample(app string, usages ...[]string) string {
	var b strings.Builder

	for i, usage := range usages {
		if len(usage) == 0 {
			continue
		}

		// first token is the subcommand
		b.WriteString(
			chalk.White.Color(app) + " " +
				chalk.White.Color(chalk.Bold.TextStyle(usage[0])),
		)

		// remaining tokens are either flags (--foo) or args
		for _, tok := range usage[1:] {
			switch {
			case strings.HasPrefix(tok, "--"):
				b.WriteString(" " + chalk.Italic.TextStyle(chalk.White.Color(tok)))
			default:
				b.WriteString(" " + chalk.Dim.TextStyle(chalk.Italic.TextStyle(tok)))
			}
		}

		if i < len(usages)-1 {
			b.WriteRune('\n')
		}
	}

	return b.String()
}

////////////////////////////////////////////////////////////////////////////////////////////////////
