package jb

import (
	"fmt"
	"os"
	"os/exec"
	"pj/shortcut"
	"strings"
)

var idea string
var web string
var golang string
var python string
var c string
var android = ""
var androidLnk = "C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs\\Android Studio\\Android Studio.lnk"
var vscode string = "C:\\Program Files\\Microsoft VS Code\\Code.exe"
var exeHomePath = "C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs\\JetBrains"

func findExe() {

	files, err := os.ReadDir(exeHomePath)
	if err != nil {
		fmt.Println(err)
	}
	shortcutCreator := shortcut.NewShortcutCreator()
	android = shortcutCreator.LoadShortcutTarget(androidLnk)
	for _, file := range files {
		filename := file.Name()
		//fmt.Println(filename)
		shortcutCreator := shortcut.NewShortcutCreator()
		if strings.HasPrefix(filename, "GoLand") {
			golang = exeHomePath + "\\" + filename
			golang = shortcutCreator.LoadShortcutTarget(golang)
		} else if strings.HasPrefix(filename, "IntelliJ IDEA") {
			idea = exeHomePath + "\\" + filename
			idea = shortcutCreator.LoadShortcutTarget(idea)
		} else if strings.HasPrefix(filename, "PyCharm") {
			python = exeHomePath + "\\" + filename
			python = shortcutCreator.LoadShortcutTarget(python)
		} else if strings.HasPrefix(filename, "WebStorm") {
			web = exeHomePath + "\\" + filename
			web = shortcutCreator.LoadShortcutTarget(web)
		} else if strings.HasPrefix(filename, "CLion") {
			c = exeHomePath + "\\" + filename
			c = shortcutCreator.LoadShortcutTarget(c)
		}

	}

}

func Open(app string, path string) {
	var truePath string
	var openType string
	if true {
		switch strings.ToLower(app) {
		case "idea":
			fallthrough
		case "java":
			if idea == "" {
				findExe()
			}
			openType = "idea"
			truePath = idea

		case "web":
			fallthrough
		case "webstorm":
			if web == "" {
				findExe()
			}
			openType = "webstorm"
			truePath = web

		case "go":
			fallthrough
		case "golang":
			if golang == "" {
				findExe()
			}
			openType = "golang"
			truePath = golang

		case "python":
			fallthrough
		case "py":
			if python == "" {
				findExe()
			}
			openType = "python"
			truePath = python

		case "explorer":
			fallthrough
		case "e":
			openType = "explorer"
			truePath = "explorer"

		case "clion":
			fallthrough
		case "c":
			if c == "" {
				findExe()
			}
			openType = "clion"
			truePath = c

		case "vs":
			fallthrough
		case "vscode":
			openType = "vscode"
			truePath = vscode

		case "android":
			fallthrough
		case "a":
			if truePath == "" {
				findExe()
			}
			openType = "android"
			truePath = android
		}

	}
	fmt.Printf("用打%s开仓库%s中...\n", openType, path)

	if truePath == "" {
		fmt.Println("找不到路径")
		return
	} else {
		//fmt.Println("open with jb")
		//truePath = "C:\\Program returnFiles\\JetBrains\\IntelliJ IDEA 2021.3.3\\bin\\idea64.exe"
		//truePath = "C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs\\JetBrains\\IntelliJ IDEA 2021.3.3.lnk"
		//fmt.Println(truePath)
		//fmt.Println(path)
		//fmt.Println("调试至此")
		//return "调试至此"
		err := exec.Command(truePath, path).Start()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("正在打开...")
		return
	}
}
