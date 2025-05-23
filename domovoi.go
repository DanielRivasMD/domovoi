/*
Copyright © 2025 Daniel Rivas <danielrivasmd@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

////////////////////////////////////////////////////////////////////////////////////////////////////

/*
Package domovoi provides file system utility functions that interface
with the horus error-handling library. It offers a clean and modular API
for operations like listing files, listing directories, checking file existence,
and creating directories.

Each function is implemented in its own source file for maintainability.
For example:

 - ListFiles: returns the regular file names in a given directory.
 - ListDirs: returns the subdirectory names in a given directory.
 - FileExists: checks if a file or directory exists at the provided path.
 - CreateDir: creates a directory if it does not exist, with proper error handling.

Import this package as follows:

import "github.com/yourusername/domovoi"

Then you can call its functions directly, for example:

files, err := domovoi.ListFiles("/some/directory")
if err != nil {
	// horus-styled error
	log.Fatalf("Error listing files: %v", err)
}
fmt.Println("Files:", files)
*/

////////////////////////////////////////////////////////////////////////////////////////////////////

package domovoi

////////////////////////////////////////////////////////////////////////////////////////////////////
