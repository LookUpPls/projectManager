package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"project/jb"
	"project/shortcut"
	"strings"
)

var Shortcuter = shortcut.NewShortcutCreator()
var spaceHomePath = "C:\\WorkSpace1\\_home\\"
var gitUrl = "https://github.com/1357885013/english.git"

func main() {
	// 检查命令行参数
	runWithArgs(os.Args)

	repoName := getNameFromGitUrl(gitUrl)
	fmt.Println(repoName)

	// 执行git clone命令
	clone(repoName, gitUrl, spaceHomePath+repoName+"\\")

	// 创建快捷方式
	defer Shortcuter.Close()

	fmt.Println("开始创建快捷方式...")
	Shortcuter.CreateShortcut("C:\\WorkSpace1\\ss.lnk", "C:\\WorkSpace1\\_home\\english")

	// 用IDEA打开仓库文件夹
	fmt.Println("用IDEA打开仓库...")
	fmt.Println(jb.Open("idea", spaceHomePath+repoName+"\\"))
}

func runWithArgs(args []string) {

	// 显示 help
	len := len(args)
	if len == 0 {
		fmt.Println(`
					welcome to project manager
					pj create gitUrl  [openWith]
					pj tag add name  tags split with space
					pj tag remove name  tags split with space
					pj open name [openWith]
					pj tag name
						list
					pj name
						list
					`)
		return
	}
	if len < 2 {
		fmt.Println("请提供git仓库地址")
		// return
	} else {
		gitUrl = args[1]
	}
	//todo: operation log
}

func clone(name string, url string, path string) {

	fmt.Println("开始克隆仓库...")
	cmd := exec.Command("git", "clone", gitUrl, spaceHomePath+name+"\\")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	cmd.Wait()
}

func getNameFromGitUrl(gitUrl string) string {
	slices := strings.Split(gitUrl, "/")
	repoName := strings.TrimSuffix(slices[len(slices)-1], ".git")
	//fmt.Println(repoName)
	return repoName
}
