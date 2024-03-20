package internal

import (
	"fmt"
	"os"

	"github.com/grayxiaoxiao/lazy-struct/global"
	"github.com/grayxiaoxiao/lazy-struct/utils"
	"github.com/spf13/cobra"
)

var (
  Gorm = &cobra.Command{
    Use:   "gorm",
    Short: "生成用于数据映射的结构体，使用gorm",
    Long:  `
    示例1：不带存储路径
    > lazy-struct gorm --pkg=business customer id:uint name:string
    - 'pkg=business' 结构体package business
    - 'customer' 结构体名
    - 'id:uint'  字段名:字段类型
    示例2： 带存储路径
    > lazy-struct gorm --path=dir --pkg=business customer id:uint
    - 'dir' 结构体文件将要存放的地方
    `,
    Run:   run,
  }
  Dir  string // 存储地址
  Pkg  string // 包名
  Name string // 结构体名
)

func init() {
  Gorm.Flags().StringVarP(&Dir, "dir", "d", "", "结构体文件存储路径")
  Gorm.Flags().StringVarP(&Pkg, "pkg", "p", "", "结构体package")
  Gorm.Flags().StringVarP(&Name, "name", "n", "", "结构体名")
}

func run(cmd *cobra.Command, args []string) {
  utils.PrintLog(global.INFO_MODE, fmt.Sprintf("lazy-struct gorm running with %v", args))
  if len(args) == 0 {
    utils.PrintLog(global.ERROR_MODE, "structure name is required.\033[m Example: lazy-struct gorm customer")
    return
  }
  curDir, err := os.Getwd()
  if err != nil {
    utils.PrintLog(global.ERROR_MODE, err.Error())
    return
  }
  storeDir := fmt.Sprintf("%s/%s", curDir, Dir)
  utils.PrintLog(global.INFO_MODE, fmt.Sprintf("存放路径 %s", storeDir))
  gorms := utils.GenerateFields(args)
  err    = utils.CreateStructure(Dir, Pkg, Name, gorms)
  if err != nil {
    utils.PrintLog(global.ERROR_MODE, err.Error())
  }
}