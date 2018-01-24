package main

import (
	"encoding/json"
	"io/ioutil"
)

type GitConfig struct {
	NAME   string `json:"name"`
	GITURL string `json:"giturl"`
	DESC   string `json:"desc"`
}

func loadConfigFile(filename string) []GitConfig {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		//TODO: handle error
		return nil
	}
	var jsonGitConfig []GitConfig
	json.Unmarshal(raw, &jsonGitConfig)
	return jsonGitConfig
}
