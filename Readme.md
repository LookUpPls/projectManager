# welcome to project manager

> 管理繁多的不同语言不同用途的project。 所有数据存储在文件夹内， 命令执行用于操作文件，基于tag实现， space内每个文件夹表示一个tag点，文件夹内放置对应项目的快捷方式表示该项目上有这个tag。 一行命令就能打开project用不同的方式


## 文件夹结构
> _home是所有项目的存放地  

> 下面以.lnk结尾的是快捷方式， 无尾缀的是文件夹
* space                  
  * **_home** 
    * 项目1 
    * 项目2
  * tag1
  * tag2
    * 项目1.lnk
  * tag3
    * 项目1.lnk
    * 项目2.lnk

## 命令
 pj config spaceLocation location   : set the space location

 pj create gitUrl  [openWith]       : create project use the git url with git clone command

 pj open name [openWith]            : open project, can open with java.idea.web.go.py.python.e.explorer

 pj tag                   : list all tags
 pj tag ls  [tags]        : list projects with given tags
 pj tag add name tags     : add tags to project,multi tags split with space
 pj tag rm  name tags     : rm  tags to project,multi tags split with space

 pj ls                    :list all the project
 pj name                  :list project's tags

 pj tidy                  :clear tag which project deleted,只会删除空文件夹和lnk文件


## 安装
把pj.exe下载复制到哪里都行， 然后加入到系统变量 path 里。