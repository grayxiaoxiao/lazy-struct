package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/grayxiaoxiao/lazy-struct/cover"
	"github.com/grayxiaoxiao/lazy-struct/global"
	"github.com/grayxiaoxiao/lazy-struct/utils"
	"github.com/spf13/cobra"
)

var (
  GenCmd = &cobra.Command{
    Use:   "gen",
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
  Classify string // 用于数据映射 还是 配置映射
)

func init() {
  GenCmd.Flags().StringVarP(&Dir, "dir", "d", "", "结构体文件存储路径")
  GenCmd.Flags().StringVarP(&Pkg, "pkg", "p", "", "结构体package")
  GenCmd.Flags().StringVarP(&Name, "name", "n", "", "结构体名")
  GenCmd.Flags().StringVarP(&Classify, "classify", "c", "dm", "选择用于数据映射还是配置映射结构(dm数据映射 sm配置映射)")
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
  fields   := utils.GenerateFields(Classify, args)
  template := cover.Template{PackageName: Pkg, StructureName: Name, Fields: fields}
  bufstr, err := template.Parse()
  if err != nil {
    utils.PrintLog(global.ERROR_MODE, err.Error())
    return
  }
  filename  := fmt.Sprintf("%s.go", strings.ToLower(Name))
  file, err := utils.CreateFile(Dir, filename)
  if err != nil {
    utils.PrintLog(global.ERROR_MODE, err.Error())
    return
  }
  _, err     = file.WriteString(strings.Replace(bufstr, "\n\n", "\n", -1))
  if err != nil {
    utils.PrintLog(global.ERROR_MODE, err.Error())
  }
  file.Close()
}