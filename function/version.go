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
	Name    = "Checker"
	Version = "v0.3.0"
	Path    = "github.com/yhyj/skynet"
)

func ProgramInfo(only bool) string {
	programInfo := fmt.Sprintf("%s\n", Version)
	if !only {
		programInfo = fmt.Sprintf("%s version %s\n", Name, Version)
	}
	return programInfo
}
