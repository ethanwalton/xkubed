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

//trying to get activity to show on github since it isn't for some reason
