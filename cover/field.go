package cover

import "fmt"

type Field struct {
  Name     string
  Type     string
  Tags     string
  JsonName string
  Classify uint8 // 1:gorm数据映射结构体 2:viper配置解析结构体
}

func(f *Field) SetTags() {
  if f.Classify == 1 {
    f.Tags = fmt.Sprintf("`json:\"%s\" gorm:\"comment:%s\"`", f.JsonName, f.JsonName)
  } else if f.Classify == 2 {
    f.Tags = fmt.Sprintf("`mapstructure:\"%s\" json:\"%s\" yaml:\"%s\"`", f.JsonName, f.JsonName, f.JsonName)
  }
}