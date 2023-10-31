/*
File: version.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-27 12:50:26

Description: 子命令`version`功能函数
*/

package function

import "fmt"

// 程序信息
const (
	name    = "Checker"
	version = "v0.3.2"
	project = "github.com/yhyj/checker"
)

// 编译信息
var (
	gitCommitHash string = "unknown"
	buildTime     string = "unknown"
	buildBy       string = "unknown"
)

func ProgramInfo(only bool) string {
	programInfo := fmt.Sprintf("%s\n", version)
	if !only {
		programInfo = fmt.Sprintf("%s version: %s\nGit commit hash: %s\nBuilt on: %s\nBuilt by: %s\n", name, version, gitCommitHash, buildTime, buildBy)
	}
	return programInfo
}
