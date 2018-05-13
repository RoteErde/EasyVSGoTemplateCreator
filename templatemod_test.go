package main

import (
	"testing"
)

func TestNewProjectInfoList(t *testing.T) {
	plist := NewProjectInfoList()
	if len(plist) == 0 {
		t.Error("Empty list, expecting more than 0 elements")
	}
}

func Test_createProject(t *testing.T) {
	plist := NewProjectInfoList()
	createProject(plist[0], "ds")
}

func Test_getProjectType(t *testing.T) {
	type args struct {
		pinfo       []ProjectInfo
		projecttype string
	}
	validArgs := args{NewProjectInfoList(), "ts"}
	invalidArgs := args{NewProjectInfoList(), "tss"}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test for valid list", validArgs, 2, false},
		{"test for valid list", invalidArgs, -1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getProjectType(tt.args.pinfo, tt.args.projecttype)
			if (err != nil) != tt.wantErr {
				t.Errorf("getProjectType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getProjectType() = %v, want %v", got, tt.want)
			}
		})
	}
}
