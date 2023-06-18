package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudweops/phoenix/cmd/generate/enum"
	"github.com/spf13/cobra"
)

// Cmd 枚举生成器
var Cmd = &cobra.Command{
	Use:   "enum",
	Short: "枚举生成器",
	Long:  `枚举生成器`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("input file is mandatory, see: -help")
		}

		matchedFiles := []string{}
		for _, v := range args {
			files, err := filepath.Glob(v)
			if err != nil {
				return err
			}
			// 只匹配Go源码文件
			if strings.HasSuffix(v, ".go") {
				matchedFiles = append(matchedFiles, files...)
			}
		}

		if len(matchedFiles) == 0 {
			return fmt.Errorf("no file matched")
		}

		for _, path := range matchedFiles {
			// 生成代码
			code, err := enum.G.Generate(path)
			if err != nil {
				return err
			}

			if len(code) == 0 {
				continue
			}

			var genFile = ""
			if strings.HasSuffix(path, ".pb.go") {
				genFile = strings.ReplaceAll(path, ".pb.go", "_enum.pb.go")
			} else {
				genFile = strings.ReplaceAll(path, ".go", "_enum.go")
			}

			// 写入文件
			err = os.WriteFile(genFile, code, 0644)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	Cmd.PersistentFlags().BoolVarP(&enum.G.Marshal, "marshal", "m", false, "is generate json MarshalJSON and UnmarshalJSON method")
	Cmd.PersistentFlags().BoolVarP(&enum.G.ProtobufExt, "protobuf_ext", "p", false, "is generate protobuf extention method")
}
