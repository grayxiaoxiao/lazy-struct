package cmd

import (
	"github.com/grayxiaoxiao/lazy-struct/cmd/internal"
	"github.com/spf13/cobra"
)

type Root struct {
  Name        string
  Params      string
  Version     string
  Abstract    string
  Description string
}

func NewRootCommand() {
}

var (
  RootCmd = &cobra.Command{
    Use:     "lazy-struct",
    Short:   "",
    Long:    "",
    Version: "1.0",
  }
)

func init() {
  RootCmd.AddCommand(internal.GenCmd)
  RootCmd.AddCommand(internal.ImpCmd)
}