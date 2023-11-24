/*
File: pacnew.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-27 12:59:42

Description: 子命令`pacnew`功能函数
*/

package cli

import "github.com/yhyj/checker/general"

// PacnewChecker 检测是否有 pacnew 文件
//
//   - pacnew 文件是 Arch Linux 系统软件包更新后其中的配置文件等比旧版有更改时将新版的文件添加后缀名 .pacnew
//
// 返回：
//   - 错误信息
func PacnewChecker() error {
	pacnewArgs := []string{}
	if err := general.RunCommand("pacdiff", pacnewArgs); err != nil {
		return err
	}
	return nil
}
