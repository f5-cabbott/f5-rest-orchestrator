package cmd

import (
	"github.com/spf13/cobra"
)

func newStateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "state",
		Short: "Inspect or manipulate the f5acmectl state backend",
	}
	cmd.AddCommand(
		&cobra.Command{
			Use:   "show",
			Short: "Print the current state snapshot",
			RunE:  func(cmd *cobra.Command, args []string) error { return errNotImplemented("state show") },
		},
		&cobra.Command{
			Use:   "list",
			Short: "List versioned state snapshots",
			RunE:  func(cmd *cobra.Command, args []string) error { return errNotImplemented("state list") },
		},
		&cobra.Command{
			Use:   "lock",
			Short: "Acquire the state lock (advisory; for emergency holds)",
			RunE:  func(cmd *cobra.Command, args []string) error { return errNotImplemented("state lock") },
		},
		&cobra.Command{
			Use:   "unlock",
			Short: "Force-release a stuck state lock",
			RunE:  func(cmd *cobra.Command, args []string) error { return errNotImplemented("state unlock") },
		},
	)
	return cmd
}
