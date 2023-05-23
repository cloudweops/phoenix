package cmd

import (
	"fmt"
	"os"

	"github.com/CloudWeOps/phoenix/cmd/phoenix/cmd/generate"
	"github.com/CloudWeOps/phoenix/cmd/phoenix/cmd/project"
	"github.com/CloudWeOps/phoenix/cmd/phoenix/cmd/protobuf"
	"github.com/CloudWeOps/phoenix/version"
	"github.com/spf13/cobra"
)

var vers bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "phoenix",
	Short: "phoenix 分布式服务构建工具",
	Long:  `phoenix 分布式服务构建工具`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
			return nil
		}
		return cmd.Help()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.AddCommand(project.Cmd, generate.Cmd, protobuf.Cmd)
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "the phoenix version")
}
