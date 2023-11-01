/*
File: pacnew.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-21 11:24:15

Description: 程序子命令'pacnew'时执行
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yhyj/checker/cli"
)

// pacnewCmd represents the pacnew command
var pacnewCmd = &cobra.Command{
	Use:   "pacnew",
	Short: "Check and compare the old and new configuration files of the package",
	Long:  `Check and compare the old and new configuration files of installed packages.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cli.PacnewChecker(); err != nil {
			fmt.Printf("\x1b[31m%s\x1b[0m\n", err)
		}
	},
}

func init() {
	pacnewCmd.Flags().BoolP("help", "h", false, "help for pacnew command")
	rootCmd.AddCommand(pacnewCmd)
}
