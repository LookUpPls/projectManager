package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"project/cfg"
	"project/jb"
	"project/shortcut"
	"strings"
)

var (
	helpText = `
welcome to project manager

pj config projectLocation location

pj create gitUrl  [openWith]
pj open name [openWith]

pj tag list  [tags_split_with_space]        :list with tags
pj tag add name  tags_split_with_space
pj tag remove name  tags_split_with_space
pj tag name         :list project with the tag

pj list             :list all the project
pj name             :list tags with the name

pj tidy
	clear tag which project deleted
					`
	Shortcuter    = shortcut.NewShortcutCreator()
	spacePath     = "C:\\WorkSpace1\\"
	spaceHomePath = spacePath + "_home\\"
	gitUrl        = "https://github.com/1357885013/english.git"
	config        = cfg.Config{}
	projectConfig = cfg.ProjectConfig{}
	inited        = true

	repoName    = ""
	openWith    = ""
	projectName = ""
)

func main() {
	// 检查命令行参数
	fmt.Println("welcome to project manager")
	// 加载程序配置
	config = *config.LoadConfig()
	if config.ProjectLocation == "" {
		fmt.Println("项目地址未配置， 请用命令pj config projectLocation location配置地址")
		inited = false
	}
	// 加载项目配置
	if inited {
		projectConfig = *projectConfig.LoadConfig(config.ProjectLocation)
	}

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
		fmt.Println(helpText)
		return
	}

	switch strings.ToLower(args[0]) {
	case "cfg":
		fallthrough
	case "config":
		goto setConfig
	}

	if !inited {
		fmt.Println("请先配置项目地址")
		return
	}

	//处理第二个参数
	switch strings.ToLower(args[0]) {
	case "open":
		goto openProject
	case "create":
		fallthrough
	case "new":
		goto newProject
	case "tag":
		if len == 1 {
			// list all tag
			fmt.Println("all tags ")
			printFiles(spacePath)
			return
		}

		switch strings.ToLower(args[1]) {
		case "list":
			fallthrough
		case "ls":
			fmt.Println("all tags ")
			printFiles(spacePath)
			return
		case "add":
			fallthrough
		case "new":
			fallthrough
		case "create":
			goto addTag
		case "delete":
			fallthrough
		case "remove":
			fallthrough
		case "rm":
			goto deleteTag

		}
	case "list":
		fallthrough
	case "all":
		fmt.Println("all projects:")
		printFiles(spaceHomePath)
	default:

	}

	//todo: operation log

setConfig:
	if len < 3 {
		fmt.Println("参数长度小于3， 请提供完整命令")
	}
	switch strings.ToLower(args[1]) {
	case "projectlocation":
		if strings.HasSuffix(args[2], "\\") {
			config.ProjectLocation = args[2]
		} else {
			config.ProjectLocation = args[2] + "\\"
		}
		config.SaveConfig()
	}
	return

openProject:
	if len == 1 {
		fmt.Println("请提供name")
		return
	}
	repoName = args[1]
	openWith = "idea"
	if len == 3 {
		openWith = args[2]
		projectConfig.SetOpenMethod(repoName, openWith)
		projectConfig.SaveConfig(config.ProjectLocation)
	} else {
		if method := projectConfig.GetOpenMethod(repoName); method != "" {
			openWith = method
		}
	}
	if jb.Open(openWith, spaceHomePath+repoName+"\\") == "成功" {
		fmt.Println("成功打开")
	}
	return
newProject:
	if len == 1 {
		fmt.Println("请提供git仓库地址")
		return
	}
	gitUrl = args[1]
	repoName = getNameFromGitUrl(gitUrl)
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
	return
addTag:
	fmt.Println("111")
	if len < 3 {
		fmt.Println("请输入project name")
	}
	projectName = args[2]
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
	return
deleteTag:
	if len < 3 {
		fmt.Println("请输入project name")
	}
	projectName = args[2]
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
	return
}

func printFiles(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		if file.Name() != "_home" {
			fmt.Print(file.Name() + " ")
		}
	}
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
