package protobuf

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "proto",
	Short: "项目 protocol buffer 管理类工具",
	Long:  `项目 protocol buffer 管理类工具`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
