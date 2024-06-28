/*
File: aur.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-27 12:59:16

Description: 子命令 'aur' 的实现
*/

package cli

import "github.com/yhyj/checker/general"

// AURChecker 检测已安装的属于 AUR 的异常包
//
// 返回：
//   - 标准输出
func AURChecker() string {
	aurArgs := []string{"-c", `pacman -Qmq | parallel 'result=$(package-query -AQ -f "%v" "{}" | uniq -d | wc -l); [ $result -eq 0  ] && echo "{}"'`}
	stdout, _, _ := general.RunCommandToBuffer("bash", aurArgs)
	return stdout
}
