package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/grayxiaoxiao/lazy-struct/global"
	"github.com/grayxiaoxiao/lazy-struct/utils"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
  ImpCmd = &cobra.Command{
    Use:   "imp",
    Short: "从MySQL的表描述文件中，生成对应的映射结构",
    Long:  "",
    Run: impRun,
  }
  sqlfile string
  pkg     string
)

func init() {
  ImpCmd.Flags().StringVarP(&sqlfile, "file", "f", "", "通过mysqldump命令导出的不带数据的SQL文件")
  ImpCmd.Flags().StringVarP(&pkg, "pkg", "p", "", "结构体的package，也是最终存放的目录")
}

func impRun(cmd *cobra.Command, args []string) {
  utils.PrintLog(global.INFO_MODE, fmt.Sprintf("lazy-struct imp running with %v", args))
  parseFile()
}

func parseFile() {
  file, _ := os.Open(sqlfile)
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)
  set      := false
  var text string
  tablesql := make([]string, 0, 20)
  for scanner.Scan() {
    text = scanner.Text()
    if strings.Contains(text, "CREATE TABLE") && !set {
      set = true
      tablesql = tablesql[:0]
    }
    if set {
      tablesql = append(tablesql, text)
    }
    if strings.Contains(text, "CHARSET=utf8;") && set {
      set = false
      tableName, fields := generateFields(tablesql)
      fmt.Println("tableName = ", tableName, " and fields = ", fields)
    }
  }
  file.Close()
}
/*
CREATE TABLE `action_histories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `org_id` int DEFAULT NULL,
  `member_id` int DEFAULT NULL,
  `controller_name` varchar(255) DEFAULT NULL,
  `action_name` varchar(255) DEFAULT NULL,
  `visit_ip` varchar(255) DEFAULT NULL,
  `parameters` text,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_action_histories_on_org_id` (`org_id`),
  KEY `index_action_histories_on_member_id` (`member_id`),
  KEY `index_action_histories_on_org_id_and_member_id` (`org_id`,`member_id`),
  KEY `index_action_histories_on_controller_name` (`controller_name`),
  KEY `index_action_histories_on_action_name` (`action_name`),
  KEY `index_action_histories_on_controller_name_and_action_name` (`controller_name`,`action_name`),
  KEY `index_action_histories_on_created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=26258 DEFAULT CHARSET=utf8;
*/

func generateFields(tablesql []string) (table string, fields []([]string)) {
  table  = getTableName(tablesql)
  fields = getFields(tablesql)
  return
}

func getFields(tablesql []string) []([]string) {
  fields := make([]([]string), 0, 20) // [field_name, field_data]
  for _, line := range tablesql {
    line = strings.TrimLeft(line, " ")
    if strings.HasPrefix(line, "`") {
      _linesp := strings.Split(line, " ")
      fields = append(fields, []string{_linesp[0], _linesp[1]})
    }
  }
  return fields
}

func getTableName(tablesql []string) string {
  _str := strings.Split(tablesql[0], " ")
  return _str[2]
}

func titleField(fieldName string) string {
  caser := cases.Title(language.English)
  _name := ""
  for _, field := range strings.Split(fieldName, "_") {
    _name += caser.String(field)
  }
  return _name
}

func typeMapping(typeName string) string {
  switch typeName {
  case "int", "bigint":
    return "int64"
  case "smallint":
    return "int8"
  case "float", "double":
    return "float64"
  case "text":
    return "text"
  case "date":
    return "date"
  case "datetime":
    return "time"
  default:
    return "string"
  }
}
