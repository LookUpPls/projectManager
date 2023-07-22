package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestBetterPrintFiles1(t *testing.T) {
	files, err := os.ReadDir("C:\\WorkSpace\\_home\\")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		if file.Name() != "_home" && file.IsDir() {
			fmt.Print(file.Name() + "    ")
		}
	}
}
func TestBetterPrintFiles(t *testing.T) {
	files, err := os.ReadDir("C:\\WorkSpace\\_home\\")
	if err != nil {
		fmt.Println(err)
	}
	all := make([]string, len(files))
	index := 0
	maxLen := 0
	for _, file := range files {
		if file.Name() != "_home" && file.IsDir() {
			all[index] = file.Name()
			if ll := len(file.Name()); ll > maxLen {
				maxLen = ll
			}
			index++
		}
	}
	all = all[0:index]

	for _, each := range all {
		fmt.Printf("%-"+strconv.Itoa(maxLen)+"s", each+"\t")
	}
}

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

func TestOpen(t *testing.T) {
	runWithArgs([]string{"test", "open", "english"})
}
func TestOpenWithIdea(t *testing.T) {
	runWithArgs([]string{"test", "open", "english", "idea"})
}
func TestOpenWithWeb(t *testing.T) {
	runWithArgs([]string{"test", "open", "english", "web"})
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
