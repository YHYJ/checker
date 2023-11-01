<h1 align="center">Checker</h1>

<!-- File: README.md -->
<!-- Author: YJ -->
<!-- Email: yj1516268@outlook.com -->
<!-- Created Time: 2023-02-27 11:40:09 -->

---

<p align="center">
  <a href="https://github.com/YHYJ/checker/actions/workflows/release.yml"><img src="https://github.com/YHYJ/checker/actions/workflows/release.yml/badge.svg" alt="Go build and release by GoReleaser"></a>
</p>

---

## Table of Contents

<!-- vim-markdown-toc GFM -->

* [Usage](#usage)
* [Compile](#compile)
  * [当前平台](#当前平台)
  * [交叉编译](#交叉编译)
    * [Linux](#linux)
    * [macOS](#macos)
    * [Windows](#windows)
* [Dependencies](#dependencies)

<!-- vim-markdown-toc -->

---

<!------------------------------------------>
<!--       _               _              -->
<!--   ___| |__   ___  ___| | _____ _ __  -->
<!--  / __| '_ \ / _ \/ __| |/ / _ \ '__| -->
<!-- | (__| | | |  __/ (__|   <  __/ |    -->
<!--  \___|_| |_|\___|\___|_|\_\___|_|    -->
<!------------------------------------------>

---

适用于 Arch Linux 的系统检查工具

## Usage

- `aur`子命令

  该子命令检查本地安装的 AUR 包是否与源保持同步

- `pacnew`子命令

  该子命令检查并比较已更新包的新旧配置文件

- `version`子命令

  查看程序版本信息

- `help`

  查看程序帮助信息

## Compile

### 当前平台

```bash
go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/checker/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/checker/general.BuildTime=`date +%s` -X github.com/yhyj/checker/general.BuildBy=$USER" -o build/checker main.go
```

### 交叉编译

使用命令`go tool dist list`查看支持的平台

#### Linux

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/checker/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/checker/general.BuildTime=`date +%s` -X github.com/yhyj/checker/general.BuildBy=$USER" -o build/checker main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64

#### macOS

```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/checker/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/checker/general.BuildTime=`date +%s` -X github.com/yhyj/checker/general.BuildBy=$USER" -o build/checker main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64

#### Windows

```powershell
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -H windowsgui -X github.com/yhyj/checker/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/checker/general.BuildTime=`date +%s` -X github.com/yhyj/checker/general.BuildBy=$USER" -o build/checker.exe main.go
```

> 使用`echo %PROCESSOR_ARCHITECTURE%`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64

## Dependencies

- bash
- coreutils
- package-query
- pacman
- pacman-contrib
- parallel
