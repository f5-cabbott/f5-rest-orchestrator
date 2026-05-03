package cmd

import (
	"github.com/spf13/cobra"
)

func newVerifyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "verify",
		Short: "Post-apply checks: certs installed, chains valid, profiles point at the right cert",
		Long: `For each ManagedCertificate in scope: confirm the cert exists on the
target BIG-IP, that notAfter is in the future and beyond the renewal window,
that the chain validates against the trusted CA bundle, and that the
configured client/serverSSL profiles reference the issued cert and key.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errNotImplemented("verify")
		},
	}
}
