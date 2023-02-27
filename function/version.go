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
var (
	name    string = "Checker"
	version string = "v0.1.1"
)

func ProgramInfo() string {
	programInfo := fmt.Sprintf("\033[1m%s\033[0m %s \033[1m%s\033[0m\n", name, "version", version)
	return programInfo
}
