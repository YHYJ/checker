/*
File: update.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-12-15 13:33:21

Description: 执行子命令 'update'
*/

package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/yhyj/checker/cli"
	"github.com/yhyj/checker/general"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check pending updates in Arch Linux and AUR",
	Long:  `Safely check the list of pending updates, including Arch Linux official repositories and AUR.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取配置文件路径
		configFile, _ := cmd.Flags().GetString("config")
		// 解析参数
		emptyFlag, _ := cmd.Flags().GetBool("empty")
		fileFlag, _ := cmd.Flags().GetBool("file")

		var (
			writeMode string = "t" // 记录文件写入模式，t 表示覆盖写入，a 表示追加写入
		)

		// 读取配置文件
		configTree, err := general.GetTomlConfig(configFile)
		if err != nil {
			fileName, lineNo := general.GetCallerInfo()
			color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
			return
		}
		// 获取配置项
		config, err := general.LoadConfigToStruct(configTree)
		if err != nil {
			fileName, lineNo := general.GetCallerInfo()
			color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
			return
		}

		// 清空记录文件
		if emptyFlag {
			general.EmptyFile(config.Update.ArchRecordFile)
			general.EmptyFile(config.Update.AurRecordFile)
			return
		}

		// 获取更新信息
		archUpdateList := cli.CheckArchUpdates(config)
		aurUpdateList := cli.CheckAurUpdates(config)
		// 根据 file 参数决定输出到文件还是终端
		if fileFlag {
			general.WriteFile(config.Update.ArchRecordFile, archUpdateList, writeMode)
			general.WriteFile(config.Update.AurRecordFile, aurUpdateList, writeMode)
		} else {
			color.Printf("%s\n%s\n", general.PrimaryText("Arch Official Repository:"), general.LightText(archUpdateList)) // 打印更新列表
			color.Println()
			color.Printf("%s\n%s\n", general.PrimaryText("Arch User Repository:"), general.LightText(aurUpdateList)) // 打印更新列表
		}
	},
}

func init() {
	updateCmd.Flags().Bool("empty", false, "Clear the file that records update information")
	updateCmd.Flags().Bool("file", false, "The list of pending updates will be recorded in the file")

	updateCmd.Flags().BoolP("help", "h", false, "help for update command")
	rootCmd.AddCommand(updateCmd)
}
