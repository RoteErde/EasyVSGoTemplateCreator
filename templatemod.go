package main

import (
	"errors"
	"time"

	"gopkg.in/src-d/go-git.v4"
)

//ProjectInfo contains keyword and url to pull the respective repo
type ProjectInfo struct {
	Keyword     string
	Url         string
	Description string
}

//NewProjectInfoList returns an list of projects
func NewProjectInfoList() []ProjectInfo {
	return []ProjectInfo{
		ProjectInfo{Keyword: "golang", Url: "https://github.com/RoteErde/VSCodeGoLangStarterTemplate"},
		ProjectInfo{Keyword: "rust", Url: "https://github.com/RoteErde/RustVSCodeTemplate"},
		ProjectInfo{Keyword: "ts", Url: "https://github.com/rebooting/TypeScriptVSCodeTemplate"},
		ProjectInfo{Keyword: "js", Url: "https://github.com/rebooting/JsWebDeVSCodeTemplate"},
		ProjectInfo{Keyword: "winpython", Url: "https://github.com/RoteErde/PythonWin32Template"},
		ProjectInfo{Keyword: "webgae", Url: "https://github.com/rebooting/sampleGAEWeb", Description: "Google App Engine Python Sample"},
		ProjectInfo{Keyword: "electron", Url: "https://github.com/RoteErde/VSElectronReactStarter", Description: "Electron React App"},
	}
}

func getProjectType(pinfo []ProjectInfo, projecttype string) (int, error) {
	for index := range pinfo {
		if pinfo[index].Keyword == projecttype {
			return index, nil
		}
	}
	//not found
	return -1, errors.New("project type not found")
}

func createProject(pinfo ProjectInfo, newProjectName string) error {
	/*
		gitUrlTokens := strings.Split(pinfo.Url, "/")
		if len(gitUrlTokens) <= 0 {

			return errors.New("expecting WWW")

		}
		newProjectName := gitUrlTokens[len(gitUrlTokens)-1]
	*/
	deleteExistingDirectory(newProjectName) // if it exists
	log("cloning github template")

	_, cloneerr := git.PlainClone(newProjectName, false, &git.CloneOptions{
		URL:               pinfo.Url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	if cloneerr != nil {
		return cloneerr
	}
	time.Sleep(time.Millisecond * 200)
	purgeExistingGitDirectory(newProjectName)
	return nil
}
