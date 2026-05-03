package cmd

import (
	"github.com/spf13/cobra"
)

func newDiscoverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "discover",
		Short: "Inspect current state of gateway EABs and BIG-IP ACME config; emit observed-state report",
		Long: `Read-only. Lists existing EABs at the AnyCA Gateway admin API and
queries each BigIPDevice via iControl REST for: TMOS version, ACME provider
configuration, configured ACME accounts, installed certs, and SSL profile
bindings. Output is a structured snapshot used by 'plan'.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errNotImplemented("discover")
		},
	}
	cmd.Flags().String("output", "yaml", "yaml|json")
	return cmd
}
