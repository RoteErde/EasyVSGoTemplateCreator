package main

/*
Copyrights 2017 Alvin Ng


  This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/
import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	args := os.Args

	plist := NewProjectInfoList()

	if len(args) > 2 {
		var newProjectName = args[1]
		/*
			handles any invalid entry for project type
			and prints help message
		*/
		defer func() {
			if r := recover(); r != nil {
				printHelp(plist)
			}
		}()

		if isAlphanumeric(newProjectName) {
			var projectType = strings.ToLower(args[2])
			plistindex, indexerr := getProjectType(plist, projectType)
			if indexerr != nil {
				printHelp(plist)
				return
			}
			createProject(plist[plistindex], newProjectName)
		} else {
			printHelp(plist)
		}
	} else {
		printHelp(plist)
	}

}

func printHelp(plist []ProjectInfo) {
	args := os.Args
	fmt.Println("Usage:")
	fmt.Print(args[0] + " <project> <")
	for _, eachproj := range plist {
		fmt.Print(eachproj.Keyword)
		fmt.Print("/")
	}
	fmt.Println(">")
	fmt.Printf("Name\t\t\tDescription\n")
	fmt.Println("===========================")
	for _, eachproj := range plist {
		fmt.Printf("%s\t\t\t%s\n", eachproj.Keyword, eachproj.Description)

	}

}

/*
	returns true if character is either alphabet or number
	we don't want folders with funny characters
*/
func isAlphanumeric(foldername string) bool {
	for _, character := range foldername {
		if !unicode.IsDigit(character) && !unicode.IsLetter(character) {
			return false
		}
	}
	return true
}
