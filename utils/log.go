package utils

import (
	"fmt"
	"os"

	"github.com/grayxiaoxiao/lazy-struct/global"
)

func PrintLog(mode, message string) {
  var format string
  switch mode {
  case global.ERROR_MODE:
    format = "\033[31m%s\033[m\n"
  case global.WARN_MODE:
    format = "\033[33m%s\033[m\n"
  case global.INFO_MODE:
    format =  "\033[32m%s\033[m\n"
  default:
    format = "%s"
  }
  fmt.Fprintf(os.Stdout, format, message)
}