package main

import (
	"github.com/ethanwalton/xkubed/cmd"
)

func cli() {
	cmd.Execute()
}

func init() {
	BaseCmd.PersistentFlags().StringVarP(&baseConfig.kubeConfig)
}
