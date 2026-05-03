package cmd

import (
	"github.com/spf13/cobra"
)

func newApplyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Execute a plan: mint EABs, configure ACME on BIG-IPs, bind certs to profiles",
		Long: `Idempotent. Acquires a state-backend lock, executes the plan in
dependency order (gateway -> EAB -> ACME account -> cert order -> binding),
verifies post-conditions, and writes the resulting state snapshot. Requires
either --plan-file from a prior 'plan' or --yes for non-interactive use.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errNotImplemented("apply")
		},
	}
	cmd.Flags().String("plan-file", "", "Plan file produced by 'f5acmectl plan --out'")
	cmd.Flags().Bool("yes", false, "Skip the interactive confirmation prompt")
	cmd.Flags().Bool("allow-destroy", false, "Permit destructive changes (EAB revocation, profile unbinding)")
	cmd.Flags().Duration("timeout", 0, "Per-order timeout (e.g. 10m). 0 = use built-in default.")
	return cmd
}
