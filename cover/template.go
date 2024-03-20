package cover

import (
	"bytes"
	"text/template"
)

type Template struct {
  PackageName   string `json:"packageName"`
  StructureName string `json:"sttructureName"`
  Fields      []Field  `json:"fields"`
}


func(t *Template) Parse() (string, error) {
  var _templateStr = `
  package {{.PackageName}}

  type {{.StructureName}} struct {
    {{- range .Fields }}
      {{.Name}} {{.Type}} {{.Tags}}
    {{- end}}
  }
  `
  tmpl, err := template.New("structure").Parse(_templateStr)
  if err != nil {
    return "", err
  }
  buffer := new(bytes.Buffer)
  err     = tmpl.Execute(buffer, t)
  if err != nil {
    return "", err
  }
  return buffer.String(), nil
}