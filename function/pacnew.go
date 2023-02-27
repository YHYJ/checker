/*
File: pacnew.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-27 12:59:42

Description: 子命令`pacnew`功能函数
*/

package function

func PacnewChecker() {
	pacnewArgs := []string{}
	RunCommand("pacdiff", pacnewArgs)
}
