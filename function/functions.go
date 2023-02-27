/*
File: functions.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-27 12:49:55

Description: 自定义函数
*/

package function

import (
	"os"
	"os/exec"
	"strings"
)

// 运行指定命令并获取命令输出
func RunCommandGetReturn(command string, args []string) string {
	// 定义命令
	cmd := exec.Command(command, args...)
	// 执行命令并获取命令输出
	output, _ := cmd.Output()
	// 类型转换
	result := strings.TrimRight(string(output), "\n")

	return result
}

// 运行指定命令
func RunCommand(command string, args []string) {
	// 定义命令
	cmd := exec.Command(command, args...)
	// 定义标准输入/输出/错误
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 执行命令
	cmd.Run()
}
