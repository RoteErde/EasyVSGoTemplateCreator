package main

import (
	"fmt"
	"time"

	"gopkg.in/src-d/go-git.v4"
)

/**
Alvin Ng

Logic for cloning
*/

// function prototype for map
type sub func(string, string)

func setupMappedFunctions() map[string]sub {
	mappedFunctions := map[string]sub{
		// add more templates here
		"golang": createGoProject,
		"rust":   createRustProject,
		"ts":     createTypeScriptProject,
		"js":     createJavaScriptProject,
	}

	return mappedFunctions
}

func pullTemplateRepo(projectType string, newProjectName string) bool {
	//var cmd = "git"
	var args []string

	switch projectType {
	case "golang":
		args = []string{"clone", "https://github.com/RoteErde/VSCodeGoLangStarterTemplate", newProjectName}
	case "rust":
		args = []string{"clone", "https://github.com/RoteErde/RustVSCodeTemplate", newProjectName}
	case "ts":
		args = []string{"clone", "https://github.com/rebooting/TypeScriptVSCodeTemplate", newProjectName}
	case "js":
		args = []string{"clone", "https://github.com/rebooting/JsWebDeVSCodeTemplate", newProjectName}
	default:
		fmt.Println("No valid repository specify")
		return false
	}

	git.PlainClone(newProjectName, false, &git.CloneOptions{
		URL:               args[1],
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	/*
		if cmdout, err := exec.Command(cmd, args...).Output(); err != nil {
			log(string(cmdout))
			log(err.Error())
			return false
		}*/

	return true
}

func createRustProject(newProjectName string, projectType string) {
	deleteExistingRustDirectory()
	log("cloning github template")
	if pullTemplateRepo(projectType, newProjectName) {
		log("cloning done, renaming")
		//allow a pause as I've encountered directory lock in windows
		time.Sleep(time.Millisecond * 200)
		log("removing git")
		purgeExistingGitDirectory(newProjectName)
	} else {
		log("error occured")
	}

}

func createGoProject(newProjectName string, projectType string) {
	deleteExistingGoDirectory()
	log("cloning github template")
	if pullTemplateRepo(projectType, newProjectName) {
		log("cloning done, renaming")
		//allow a pause as I've encountered directory lock in windows
		time.Sleep(time.Millisecond * 200)
		log("removing git")
		purgeExistingGitDirectory(newProjectName)
	} else {
		log("error occured")
	}

}

func createTypeScriptProject(newProjectName string, projectType string) {
	deleteExistingGoDirectory()
	log("cloning github template")
	if pullTemplateRepo(projectType, newProjectName) {
		log("cloning done, renaming")
		//allow a pause as I've encountered directory lock in windows
		time.Sleep(time.Millisecond * 200)
		log("removing git")
		purgeExistingGitDirectory(newProjectName)
	} else {
		log("error occured")
	}
}

func createJavaScriptProject(newProjectName string, projectType string) {
	deleteExistingGoDirectory()
	log("cloning github template")
	if pullTemplateRepo(projectType, newProjectName) {
		log("cloning done, renaming")
		//allow a pause as I've encountered directory lock in windows
		time.Sleep(time.Millisecond * 200)
		log("removing git")
		purgeExistingGitDirectory(newProjectName)
	} else {
		log("error occured")
	}

}
