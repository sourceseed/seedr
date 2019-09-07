package seed

import (
	"testing"
)

func Test_isGitRepo(t *testing.T) {
	tests := []struct {
		repo string
		want bool
	}{
		{"https://github.com/sourceseed/seed-golang.git", true},
		{"git@github.com:sourceseed/seed-golang.git", true},
		{"git://host.com:sourceseed.git", true},
		{"git@github.com:sourceseed/seed-golang", false},
		{"git@github.com:sourceseed.bla", false},
	}
	for _, tt := range tests {
		t.Run(tt.repo, func(t *testing.T) {
			if got := isGitRepo(tt.repo); got != tt.want {
				t.Errorf("isGitRepo(%s) = %v, want %v", tt.repo, got, tt.want)
			}
		})
	}
}

func Test_fileExists(t *testing.T) {
	tests := []struct {
		filename string
		want bool
	}{
		{"/tmp", false},
		{"/bin/bash", true},
	}
	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			if got := fileExists(tt.filename); got != tt.want {
				t.Errorf("fileExists(%s) = %v, want %v", tt.filename, got, tt.want)
			}
		})
	}
}
