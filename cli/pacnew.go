/*
File: pacnew.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-27 12:59:42

Description: 子命令 'pacnew' 的实现
*/

package cli

import "github.com/yhyj/checker/general"

// PacnewChecker 检测是否有 pacnew 文件
//
// 返回：
//   - 错误信息
func PacnewChecker() error {
	pacnewArgs := []string{}
	return general.RunCommandToOS("pacdiff", pacnewArgs)
}
