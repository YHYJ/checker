<h1 align="center">Checker</h1>
<h3 align="center">用于 Arch Linux 的系统检查工具</h3>

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

* [适配](#适配)
* [安装](#安装)
  * [一键安装](#一键安装)
  * [编译安装](#编译安装)
    * [当前平台](#当前平台)
    * [交叉编译](#交叉编译)
* [用法](#用法)
* [依赖](#依赖)

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

## 适配

- Linux: 适配 Arch Linux 及基于 Arch Linux 的发行版
- macOS: 不适配
- Windows: 不适配

## 安装

### 一键安装

```bash
curl -fsSL https://raw.githubusercontent.com/YHYJ/checker/main/install.sh | sudo bash -s
```

也可以从 [GitHub Releases](https://github.com/YHYJ/checker/releases) 下载解压后使用

### 编译安装

#### 当前平台

如果要为当前平台编译，可以使用以下命令：

```bash
go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/checker/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/checker/general.BuildTime=`date +%s` -X github.com/yhyj/checker/general.BuildBy=$USER" -o build/checker main.go
```

#### 交叉编译

> 使用命令`go tool dist list`查看支持的平台
>
> Linux 和 macOS 使用命令`uname -m`，Windows 使用命令`echo %PROCESSOR_ARCHITECTURE%` 确认系统架构
>
> - 例如 x86_64 则设 GOARCH=amd64
> - 例如 aarch64 则设 GOARCH=arm64
> - ...

因为 checker 仅支持 Linux 系统，设置如下系统变量后使用 [编译安装](#编译安装) 的命令即可进行交叉编译：

- CGO_ENABLED: 不使用 CGO，设为 0
- GOOS: 仅支持 Linux，设为 linux
- GOARCH: 根据当前系统架构设置

## 用法

- `aur`子命令

  检查本地安装的 AUR 包是否与源保持同步

- `pacnew`子命令

  检查并比较已更新包的新旧配置文件

- `version`子命令

  查看程序版本信息

- `help`子命令

  查看程序帮助信息

## 依赖

- bash
- coreutils
- package-query
- pacman
- pacman-contrib
- parallel
