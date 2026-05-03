package cmd

import (
	"github.com/spf13/cobra"
)

func newPlanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan",
		Short: "Diff desired manifests against discovered state; emit a change set",
		Long: `Composes 'validate' + 'discover' and produces a human-readable diff
plus a machine-readable plan file (when --out is set) consumable by 'apply'.
Exit codes: 0 = no changes; 2 = changes pending; non-zero otherwise.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errNotImplemented("plan")
		},
	}
	cmd.Flags().String("out", "", "Write the machine-readable plan to this path")
	cmd.Flags().Bool("detailed-exitcode", false, "Use exit code 2 to signal pending changes")
	return cmd
}
