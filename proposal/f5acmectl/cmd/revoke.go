package cmd

import (
	"github.com/spf13/cobra"
)

func newRevokeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke <managedcertificate-name>",
		Short: "Revoke an issued certificate via ACME revokeCert and remove from BIG-IP",
		Long: `Submits an ACME revokeCert request through the gateway, then unbinds
the cert from any SSL profiles and removes it from the BIG-IP filestore.
Reason codes follow RFC 5280.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errNotImplemented("revoke")
		},
	}
	cmd.Flags().Int("reason", 0, "RFC 5280 revocation reason code")
	cmd.Flags().Bool("yes", false, "Skip confirmation")
	return cmd
}
