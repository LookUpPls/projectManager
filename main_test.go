package main

import (
	"os"
	"testing"
)

func TestCreateShortcut(t *testing.T) {
	spaceHomePath := "C:\\WorkSpace1\\_home\\"
	repoName := "english"

	// Call your function
	createShortcut(spaceHomePath + repoName + "\\")

	// Check if the shortcut was created
	if _, err := os.Stat("C:\\WorkSpace1\\ss.lnk"); os.IsNotExist(err) {
		t.Errorf("The shortcut was not created")
	}
}

func createShortcut(path string) {
	// ... Your code here
}
