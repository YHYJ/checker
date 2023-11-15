/*
File: version.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-27 12:50:26

Description: 子命令`version`功能实现
*/

package general

import (
	"fmt"
	"strconv"
	"time"
)

// 程序信息
const (
	Name    string = "Checker"
	Version string = "v0.4.6"
	Project string = "github.com/yhyj/checker"
)

// 编译信息
var (
	GitCommitHash string = "Unknown"
	BuildTime     string = "Unknown"
	BuildBy       string = "Unknown"
)

// ProgramInfo 返回程序信息
func ProgramInfo(only bool) string {
	programInfo := fmt.Sprintf("%s\n", Version)
	if !only {
		BuildTimeConverted := "Unknown"
		timestamp, err := strconv.ParseInt(BuildTime, 10, 64)
		if err == nil {
			BuildTimeConverted = time.Unix(timestamp, 0).String()
		}
		programInfo = fmt.Sprintf("%s %s - Build rev %s\nBuilt on: %s\nBuilt by: %s\n", Name, Version, GitCommitHash, BuildTimeConverted, BuildBy)
	}
	return programInfo
}
