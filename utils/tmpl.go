package utils

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/grayxiaoxiao/lazy-struct/global"
	"github.com/grayxiaoxiao/lazy-struct/tmpl"
)

func ParseTmpl(pkg, name string, fields []tmpl.Gorm) string {
  tmplStr, err := template.New("struct").Parse(tmpl.GoTMPL)
  if err != nil {
    PrintLog(global.ERROR_MODE, err.Error())
    return ""
  }
  _map := map[string]interface{}{
    "packageName": pkg,
    "structName":  strings.Title(name),
    "fields":      fields,
  }
  _buf := new(bytes.Buffer)
  err = tmplStr.Execute(_buf, _map)
  if err != nil {
    PrintLog(global.ERROR_MODE, err.Error())
    return ""
  }
  return _buf.String()
}

func CreateStructure(dir, pkg, name string, fields []tmpl.Gorm) error {
  filename  := fmt.Sprintf("%s.go", strings.ToLower(name))
  file, err := CreateFile(dir, filename)
  if err != nil {
    PrintLog(global.ERROR_MODE, err.Error())
    return err
  }
  buff  := ParseTmpl(pkg, name, fields)
  _, err = file.WriteString(strings.Replace(buff, "\n\n", "\n", -1))
  return err
}

func GenerateFields(args []string) []tmpl.Gorm {
  _len := len(args)
  if _len == 0 {
    return []tmpl.Gorm{}
  } else {
    gorms  := make([]tmpl.Gorm, _len)
    for _, item := range args {
      sp := strings.Split(item, ":")
      if len(sp) != 2 {
        PrintLog(global.WARN_MODE, fmt.Sprintf("%s 不是有效的字段参数", item))
        continue
      }
      field := tmpl.Gorm{
        Type: sp[1],
        Tags: fmt.Sprintf("`json:\"%s\" gorm:\"comment:%s\"`", sp[0], sp[0]),
      }
      name := ""
      _sp := strings.Split(sp[0], "_")
      for _, es := range _sp {
        name += strings.Title(es)
      }
      field.Name = name
      gorms = append(gorms, field)
    }
    return gorms
  }
}