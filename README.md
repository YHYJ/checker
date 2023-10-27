# README

<!-- File: README.md -->
<!-- Author: YJ -->
<!-- Email: yj1516268@outlook.com -->
<!-- Created Time: 2023-02-27 11:40:09 -->

---

## Table of Contents

<!-- vim-markdown-toc GFM -->

* [Usage](#usage)
* [Compile](#compile)
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

适用于Arch Linux的系统检查工具

## Usage

- `aur`子命令

    该子命令检查本地安装的AUR包是否与源保持同步

- `pacnew`子命令

    该子命令检查并比较已更新包的新旧配置文件

- `version`子命令

    查看程序版本信息

- `help`

    查看程序帮助信息

## Compile

- 编译当前平台可执行文件：

```bash
go build main.go
```

- **交叉编译**指定平台可执行文件：

```bash
# 适用于Linux AArch64平台
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build main.go
```

```bash
# 适用于macOS amd64平台
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
```

```bash
# 适用于Windows amd64平台
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

## Dependencies

- bash
- coreutils
- package-query
- pacman
- pacman-contrib
- parallel
