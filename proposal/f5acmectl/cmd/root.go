package cmd

import (
	"github.com/spf13/cobra"
)

type globalOpts struct {
	manifestDir string
	stateURI    string
	logLevel    string
	logFormat   string
}

var opts globalOpts

func NewRoot() *cobra.Command {
	root := &cobra.Command{
		Use:           "f5acmectl",
		Short:         "Declaratively manage BIG-IP ACMEv2 cert lifecycle against a Keyfactor AnyCA Gateway",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.PersistentFlags().StringVarP(&opts.manifestDir, "manifests", "m", "./environments", "Path to manifest directory tree")
	root.PersistentFlags().StringVar(&opts.stateURI, "state", "", "State backend URI (s3://, gs://, az://, file://)")
	root.PersistentFlags().StringVar(&opts.logLevel, "log-level", "info", "trace|debug|info|warn|error")
	root.PersistentFlags().StringVar(&opts.logFormat, "log-format", "json", "json|text")

	root.AddCommand(
		newValidateCmd(),
		newDiscoverCmd(),
		newPlanCmd(),
		newApplyCmd(),
		newVerifyCmd(),
		newReconcileCmd(),
		newRevokeCmd(),
		newRotateEABCmd(),
		newStateCmd(),
		newVersionCmd(),
	)

	return root
}
