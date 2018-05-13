package main

import (
	"os"
	"strings"
)

/*
remove current cached directory (in case previous operations failed)
*/

func deleteExistingDirectory(deldir string) bool {
	if err := os.RemoveAll(deldir); err != nil {
		log(err.Error())
		return false
	}
	return true
}

func deleteExistingGoDirectory() bool {
	if err := os.RemoveAll("VSCodeGoLangStarterTemplate"); err != nil {
		log(err.Error())
		return false
	}
	return true
}

func deleteExistingRustDirectory() bool {
	if err := os.RemoveAll("RustVSCodeTemplate"); err != nil {
		log(err.Error())
		return false
	}
	return true
}

func purgeExistingGitDirectory(newProjectName string) bool {
	projPath := strings.Join([]string{newProjectName, "/.git"}, "")
	if err := os.RemoveAll(projPath); err != nil {
		log(err.Error())
		return false
	}
	return true
}

func renameRepo(templateName string, newName string) bool {
	//if err := os.Rename("VSCodeGoLangStarterTemplate", newName); err != nil {
	if err := os.Rename(templateName, newName); err != nil {
		log("failed to rename")
		log(err.Error())
		return false
	}
	return true
}
