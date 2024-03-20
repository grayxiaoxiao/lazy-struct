package main

import (
	"os"

	"github.com/grayxiaoxiao/lazy-struct/cmd"
	"github.com/grayxiaoxiao/lazy-struct/global"
	"github.com/grayxiaoxiao/lazy-struct/utils"
)

func main() {
  utils.PrintLog(global.INFO_MODE, "Welcome to use lazy-struct........")
  if err := cmd.RootCmd.Execute(); err != nil {
    utils.PrintLog(global.ERROR_MODE, err.Error())
    os.Exit(1);
  }
}