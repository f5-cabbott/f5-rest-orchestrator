package cmd

import (
	"github.com/spf13/cobra"
)

func newValidateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "validate",
		Short: "Schema-validate manifests; resolve secret references without reading values",
		Long: `Walks the manifest directory, parses YAML, validates each document
against its kind's JSON Schema, and confirms every secret:// reference exists
in the configured backend (without fetching the secret value). Exits non-zero
on any validation error.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errNotImplemented("validate")
		},
	}
}
