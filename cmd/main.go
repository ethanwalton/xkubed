package main

import (
	"github.com/ethanwalton/xkubed/cmd/commands"
)

func main() {
	commands.Execute()
}

func init() {
	BaseCmd.PersistentFlags().StringVarP(&baseConfig.kubeConfig)
}
