### 介绍
本程序所有代码都是基于https://github.com/Readify/httpstatus 完全实现的golang版本。

原网站平时用的挺频繁的，但速度比较慢，尤其在家用起来很慢，此版本是本地部署的，速度不是问题。

### 优势
也就是golang的优势啦，不用装visual studio,不用还原nuget,不用装.net framework,一个文件搞定一切

### 怎么用
- 方法一： 下载单独[HttpStatusServer.exe](https://github.com/yongfa365/HttpStatusServer/releases)运行

- 方法二： 源码
```
git clone https://github.com/yongfa365/HttpStatusServer.git  --depth=1 %gopath%\pkg\mod\github.com\yongfa365\HttpStatusServer2

然后用vscode打开目录：
%gopath%\pkg\mod\github.com\yongfa365\HttpStatusServer

在终端运行：
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go build
```


### 改图标
参见：https://github.com/josephspurrier/goversioninfo

或直接：
```
go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo
go generate
```
### History
- 2020-07-12 v0.9.1.0 程序运行后直接打开URL
- 2020-07-12 v0.9.1.0 初始版本，.net转go首个版本
