package main

import (
	"fmt"
	"os"
	"testing"
)

func TestTagCreate(t *testing.T) {
	fmt.Println("run here")
	runWithArgs([]string{"tag", "add", "english", "java", "my"})
	fmt.Println("run here1")
	if _, err := os.Stat("C:\\WorkSpace1\\java\\english.lnk"); os.IsNotExist(err) {
		t.Errorf("The java tag was not created")
	}
	if _, err := os.Stat("C:\\WorkSpace1\\my\\english.lnk"); os.IsNotExist(err) {
		t.Errorf("The my tag was not created")
	}
}

func TestClone(t *testing.T) {
	runWithArgs([]string{"create", "https://github.com/1357885013/english.git"})
	// Check if the shortcut was created
	if _, err := os.Stat("C:\\WorkSpace1"); os.IsNotExist(err) {
		t.Errorf("The workspace was not created")
	}
	if _, err := os.Stat("C:\\WorkSpace1\\_home\\english"); os.IsNotExist(err) {
		t.Errorf("The english was not created")
	}
	if _, err := os.Stat("C:\\WorkSpace1\\english.lnk"); os.IsNotExist(err) {
		t.Errorf("The english.lnk was not created")
	}
}

func TestCloneAndOpen(t *testing.T) {
	runWithArgs([]string{"create", "https://github.com/1357885013/english.git", "idea"})
	// Check if the shortcut was created
	if _, err := os.Stat("C:\\WorkSpace1"); os.IsNotExist(err) {
		t.Errorf("The workspace was not created")
	}
	if _, err := os.Stat("C:\\WorkSpace1\\_home\\english"); os.IsNotExist(err) {
		t.Errorf("The english was not created")
	}
	if _, err := os.Stat("C:\\WorkSpace1\\english.lnk"); os.IsNotExist(err) {
		t.Errorf("The english.lnk was not created")
	}
}

func TestHelp(t *testing.T) {
	runWithArgs([]string{})
}
