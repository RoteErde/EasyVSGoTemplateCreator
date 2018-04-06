package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_createRustProject(t *testing.T) {
	os.RemoveAll("Test_RustLangGitClone")
	createRustProject("Test_RustLangGitClone", "rust")
	_, err := os.Stat("Test_RustLangGitClone")
	fmt.Println("Testing Rust Cloning")
	if err == nil {
		os.RemoveAll("Test_RustLangGitClone")
	} else {
		t.Error("Failed")
	}

}

func Test_createGoProject(t *testing.T) {
	os.RemoveAll("Test_GoLangGitClone")
	createGoProject("Test_GoLangGitClone", "golang")
	_, err := os.Stat("Test_GoLangGitClone")
	fmt.Println("Testing Go Cloning")
	if err == nil {
		os.RemoveAll("Test_GoLangGitClone")
	} else {
		t.Error("Failed")
	}

}

func Test_createTypeScript(t *testing.T) {
	os.RemoveAll("Test_TSGitClone")
	createTypeScriptProject("Test_TSGitClone", "ts")
	_, err := os.Stat("Test_TSGitClone")
	fmt.Println("Testing TS Cloning")
	if err == nil {
		os.RemoveAll("Test_TSGitClone")
	} else {
		t.Error("Failed")
	}

}

func Test_createJavaScriptProject(t *testing.T) {
	os.RemoveAll("Test_JSGitClone")
	createJavaScriptProject("Test_JSGitClone", "js")
	_, err := os.Stat("Test_JSGitClone")
	fmt.Println("Testing JS Cloning")
	if err == nil {
		os.RemoveAll("Test_JSGitClone")
	} else {
		t.Error("Failed")
	}

}
