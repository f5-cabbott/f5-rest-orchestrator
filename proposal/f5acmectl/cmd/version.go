package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Populated at build time via -ldflags.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("f5acmectl %s (commit %s, built %s)\n", version, commit, date)
		},
	}
}
