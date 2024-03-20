package internal

import (
	"fmt"

	"github.com/grayxiaoxiao/lazy-struct/global"
	"github.com/grayxiaoxiao/lazy-struct/utils"
	"github.com/spf13/cobra"
)

var (
  ImpCmd = &cobra.Command{
    Use:   "imp",
    Short: "从MySQL的表描述文件中，生成对应的映射结构",
    Long:  "",
    Run: impRun,
  }
)

func impRun(cmd *cobra.Command, args []string) {
  utils.PrintLog(global.INFO_MODE, fmt.Sprintf("lazy-struct imp running with %v", args))
}