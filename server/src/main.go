package main

import (
	//"log"
	"flag"
)

const (
	AppName = "xkubed"
)

type ScanConfig struct {
	Port                        int    `default:"8080"`
	RiskConfigFilePath          string `split_words:"true" required:"true"`
	RefreshStateIntervalMinutes int    `split_words:"true" default:"720"` // every 12 hours
}

func init() {
	flag.Parse() //parsing for flag
}

func main() {
	//fmt.Println("hello world")
	//var config ServiceConfig

}
