package cmd

import "fmt"

func errNotImplemented(name string) error {
	return fmt.Errorf("subcommand %q is a stub: implementation pending", name)
}
