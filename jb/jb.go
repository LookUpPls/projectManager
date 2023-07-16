package jb

import (
	"fmt"
	"os"
	"os/exec"
	"project/shortcut"
	"strings"
)

var idea string
var web string
var golang string
var python string
var exeHomePath = "C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs\\JetBrains"

func findExe() {

	files, err := os.ReadDir(exeHomePath)
	if err != nil {
		fmt.Println(err)
	}

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
		}

	}

}

func Open(app string, path string) string {
	fmt.Printf("用打%s开仓库%s中...\n", app, path)
	var truePath string
	switch strings.ToLower(app) {
	case "idea":
		fallthrough
	case "java":
		if idea == "" {
			findExe()
		}
		truePath = idea
	case "web":
		fallthrough
	case "webstorm":
		if web == "" {
			findExe()
		}
		truePath = web
	case "go":
		fallthrough
	case "golang":
		if golang == "" {
			findExe()
		}
		truePath = golang
	case "python":
		fallthrough
	case "py":
		if python == "" {
			findExe()
		}
		truePath = python
	}

	if truePath == "" {
		return "找不到路径"
	} else {
		//fmt.Println("open with jb")
		//truePath = "C:\\Program returnFiles\\JetBrains\\IntelliJ IDEA 2021.3.3\\bin\\idea64.exe"
		//truePath = "C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs\\JetBrains\\IntelliJ IDEA 2021.3.3.lnk"
		//fmt.Println(truePath)
		//fmt.Println(path)
		err := exec.Command(truePath, path).Start()
		if err != nil {
			return err.Error()
		}
		return "成功"
	}
}
