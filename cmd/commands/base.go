package commands

//base will contain functionality for much of the cli to function and execute commands within this repository

var baseConfig baseFlags

type baseFlags struct {
	format           string
	kubeConfig       string
	context          string
	manifest         string
	namespace        string
	minSeverity      string
	exitCode         int
	includeGenerated bool
	noColor          bool
}
