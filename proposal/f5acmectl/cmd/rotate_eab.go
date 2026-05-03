package cmd

import (
	"github.com/spf13/cobra"
)

func newRotateEABCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rotate-eab",
		Short: "Mint new EAB credentials, migrate ACME accounts, then revoke the old EAB",
		Long: `Per the configured EAB scope (per-device / per-partition / per-environment),
mints new (kid, hmacKey) pairs at the gateway, writes them to the secret
backend, reconfigures the BIG-IP native ACME client to bind a new account
under the new EAB, verifies a successful order, then revokes the previous
EAB. Idempotent and resumable.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errNotImplemented("rotate-eab")
		},
	}
	cmd.Flags().StringSlice("scope", nil, "Limit rotation to specific scopes (e.g. device=bigip-dc1-01,partition=tenant-a)")
	cmd.Flags().Bool("yes", false, "Skip confirmation")
	return cmd
}
