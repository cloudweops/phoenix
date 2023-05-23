package project

import (
	"fmt"

	"github.com/CloudWeOps/phoenix/cmd/phoenix/project"
	"github.com/spf13/cobra"

	"github.com/AlecAivazis/survey/v2/terminal"
)

// addCmd 初始化系统
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "初始化",
	Long:  `初始化一个phoenix项目`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := project.LoadConfigFromCLI()
		if err != nil {
			if err == terminal.InterruptErr {
				fmt.Println("项目初始化取消")
				return nil
			}
			return err
		}

		err = p.Init()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	Cmd.AddCommand(addCmd)
}
