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
var spacePath = "C:\\WorkSpace1\\"
var spaceHomePath = spacePath + "_home\\"
var gitUrl = "https://github.com/1357885013/english.git"

func main() {
	// 检查命令行参数
	fmt.Println("welcome to project manager")
	defer Shortcuter.Close()
	runWithArgs(os.Args)

}

func runWithArgs(args []string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover with error", err)
		}
	}()

	// 显示 help
	args = args[1:]
	len := len(args)
	fmt.Printf("收到%d个参数\n", len)
	if len == 0 {
		fmt.Println(`
welcome to project manager

pj create gitUrl  [openWith]
pj tag add name  tags split with space
pj tag remove name  tags split with space
pj open name [openWith]
pj tag name
	- list
pj name
	- list
pj tidy
	clear tag which project deleted
					`)
		return
	}
	switch strings.ToLower(args[0]) {
	case "create":
		fallthrough
	case "new":
		if len == 1 {
			fmt.Println("请提供git仓库地址")
			return
		}
		gitUrl = args[1]
		repoName := getNameFromGitUrl(gitUrl)
		fmt.Println("仓库名字：" + repoName)

		// 执行git clone命令
		clone(repoName, gitUrl, spaceHomePath+repoName+"\\")

		// 创建快捷方式
		Shortcuter.CreateShortcut("C:\\WorkSpace1\\"+repoName+".lnk", "C:\\WorkSpace1\\_home\\"+repoName)

		if len >= 3 {
			// 用IDEA打开仓库文件夹
			if jb.Open(args[2], spaceHomePath+repoName+"\\") == "成功" {
				fmt.Println("成功打开")
			}
		}
	case "tag":
		if len == 1 {
			// list all tag
			files, err := os.ReadDir(spacePath)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("all tags ")
			for _, file := range files {
				if file.Name() != "_home" {
					fmt.Print(file.Name() + " ")
				}
			}
			return
		}

		switch strings.ToLower(args[1]) {
		case "add":
			fallthrough
		case "new":
			fallthrough
		case "create":
			fmt.Println("111")
			if len < 3 {
				fmt.Println("请输入project name")
			}
			projectName := args[2]
			if _, err := os.Stat(spaceHomePath + projectName); os.IsNotExist(err) {
				fmt.Println("project not exist")
				return
			}
			if len < 4 {
				fmt.Println("请输入tag name")
			}
			for i := 3; i < len; i++ {
				err := os.MkdirAll(spacePath+args[i], 0755)
				if err != nil {
					fmt.Println(err)
				}
				Shortcuter.CreateShortcut(spacePath+args[i]+"\\"+projectName+".lnk", spaceHomePath+projectName)
			}
		case "delete":
			fallthrough
		case "remove":
			fallthrough
		case "rm":
			if len < 3 {
				fmt.Println("请输入project name")
			}
			projectName := args[2]
			if len < 4 {
				fmt.Println("请输入tag name")
			}
			for i := 3; i < len; i++ {
				fmt.Println("222")
				err := os.Remove(spacePath + args[i] + "\\" + projectName + ".lnk")
				if err != nil {
					if os.IsNotExist(err) {
						fmt.Printf("project %s 不存在tag %s \n", projectName, args[i])
						err = nil
					}
				}
				if err != nil {
					fmt.Println("fail to delete tag " + args[i])
				} else {
					fmt.Println("success delete tag " + args[i])
				}
			}

		}

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
