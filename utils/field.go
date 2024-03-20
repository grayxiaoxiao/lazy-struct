package utils

import (
	"fmt"
	"strings"

	"github.com/grayxiaoxiao/lazy-struct/cover"
	"github.com/grayxiaoxiao/lazy-struct/global"
)

func getFieldName(str string) string {
  name := ""
  _sp  := strings.Split(str, "_")
  for _, es := range _sp {
    name += strings.Title(es)
  }
  return name
}

func GenerateFields(classify string, args []string) []cover.Field {
  fields := make([]cover.Field, 0, 16)
  for _, item := range args {
    nt := strings.Split(item, ":")
    if len(nt) != 2 {
      PrintLog(global.WARN_MODE, fmt.Sprintf("%s 不是有效的字段参数", item))
      continue
    }
    _cls  := map[string]uint8{"dm": 1, "sm": 2}
    field := cover.Field{ Name: getFieldName(nt[0]), Type: nt[1], Classify: _cls[classify], JsonName: nt[0]}
    field.SetTags()
    fields = append(fields, field)
  }
  return fields
}