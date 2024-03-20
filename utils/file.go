package utils

import (
	"os"

	"github.com/grayxiaoxiao/lazy-struct/global"
)

func CreateFile(dir, filename string) (*os.File, error) {
  _, err := os.Stat(dir)
  if os.IsNotExist(err) {
    os.MkdirAll(dir, os.ModePerm)
  }
  _, err = os.Stat(dir + "/" + filename)
  if err == nil {
    PrintLog(global.WARN_MODE, "文件已存在")
    return nil, nil
  }
  file, err := os.Create(dir + "/" + filename)
  if err != nil {
    PrintLog(global.ERROR_MODE, err.Error())
    return nil, err
  }
  return file, nil
}