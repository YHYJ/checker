/*
File: aur.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-27 13:33:21

Description: 程序子命令'aur'时执行
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yhyj/checker/cli"
	"github.com/yhyj/checker/general"
)

// aurCmd represents the aur command
var aurCmd = &cobra.Command{
	Use:   "aur",
	Short: "Check AUR package synchronization",
	Long:  `Check if the locally installed AUR package is kept in sync with the source.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cli.AURChecker(); err != nil {
			fmt.Printf(general.ErrorBaseFormat, err)
		}
	},
}

func init() {
	aurCmd.Flags().BoolP("help", "h", false, "help for aur command")
	rootCmd.AddCommand(aurCmd)
}
