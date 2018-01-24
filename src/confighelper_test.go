package main

import (
	"testing"
)

func Test_loadConfigFile(t *testing.T) {

	gitconfig := loadConfigFile("../sample_helper.json")
	//fmt.Print(len(gitconfig))
	if gitconfig == nil {
		t.Error("unable to load")
	} else if len(gitconfig) != 2 {
		t.Error("Wrong length")
	}
}
