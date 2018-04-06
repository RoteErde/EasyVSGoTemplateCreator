package main

import "testing"

func Test_deleteExistingGoDirectory(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteExistingGoDirectory(); got != tt.want {
				t.Errorf("deleteExistingGoDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteExistingRustDirectory(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteExistingRustDirectory(); got != tt.want {
				t.Errorf("deleteExistingRustDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_purgeExistingGitDirectory(t *testing.T) {
	type args struct {
		newProjectName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := purgeExistingGitDirectory(tt.args.newProjectName); got != tt.want {
				t.Errorf("purgeExistingGitDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_renameRepo(t *testing.T) {
	type args struct {
		templateName string
		newName      string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renameRepo(tt.args.templateName, tt.args.newName); got != tt.want {
				t.Errorf("renameRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}
