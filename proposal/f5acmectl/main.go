// Command f5acmectl declaratively manages BIG-IP native ACMEv2 client
// configuration against a Keyfactor AnyCA Gateway REST ACME endpoint.
//
// Manifests in YAML describe GatewayConfig / BigIPDevice / ManagedCertificate
// resources; subcommands plan and apply changes idempotently with a locked
// state backend, then a reconciler watches for drift.
package main

import (
	"fmt"
	"os"

	"github.com/example/f5acmectl/cmd"
)

func main() {
	if err := cmd.NewRoot().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
