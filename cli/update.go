/*
File: update.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-12-15 13:59:16

Description: 子命令 'update' 的实现
*/
package cli

import "github.com/yhyj/checker/general"

// CheckArchUpdates 检查 Arch Linux 官方库更新
//
// 参数：
//   - config: 解析 toml 配置文件得到的配置项
//
// 返回：
//   - 标准输出
func CheckArchUpdates(config *general.Config) string {
	checkArgs := []string{"-c", config.Update.ArchChecker}
	stdout, _, _ := general.RunCommandToBuffer("bash", checkArgs)
	return stdout
}

// CheckAurUpdates 检查 AUR 库更新
//
// 参数：
//   - config: 解析 toml 配置文件得到的配置项
//
// 返回：
//   - 标准输出
func CheckAurUpdates(config *general.Config) string {
	checkArgs := []string{"-c", config.Update.AurChecker}
	stdout, _, _ := general.RunCommandToBuffer("bash", checkArgs)
	return stdout
}
