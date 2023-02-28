package commands

//base will contain functionality for much of the cli to function and execute commands within this repository
import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

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

// RootCmd defines the shell command usage for kubeaudit.
var BaseCmd = &cobra.Command{
	Use:   "kubeaudit",
	Short: "A Kubernetes security auditor",
	Long: `Kubeaudit audits Kubernetes clusters for common security controls.
kubeaudit has three modes:
  1. Manifest mode: If a Kubernetes manifest file is provided using the -f/--manifest flag, kubeaudit will audit the manifest file. Kubeaudit also supports autofixing in manifest mode using the 'autofix' command. This will fix the manifest in-place. The fixed manifest can be written to a different file using the -o/--out flag.
  2. Cluster mode: If kubeaudit detects it is running in a cluster, it will audit the other resources in the cluster.
  3. Local mode: kubeaudit will try to connect to a cluster using the local kubeconfig file ($HOME/.kube/config). A different kubeconfig location can be specified using the -c/--kubeconfig flag
`,
}

// Execute is a wrapper that accounts for errors in Commands run
func Execute() {
	if err := BaseCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
