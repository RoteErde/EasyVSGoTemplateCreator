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
	/*
		type args struct {
			filename string
		}
		tests := []struct {
			name string
			args args
			want []GitConfig
		}{
		// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := loadConfigFile(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("loadConfigFile() = %v, want %v", got, tt.want)
				}
			})
		}
	*/
}
