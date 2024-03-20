package tmpl

type Gorm struct {
  Name string
  Type string
  Tags string
}

var GoTMPL = `
package {{.packageName}}

type {{.structName}} struct {
  {{- range .fields }}
    {{.Name}} {{.Type}} {{.Tags}}
  {{- end}}
}
`