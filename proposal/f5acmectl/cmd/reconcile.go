package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

func newReconcileCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reconcile",
		Short: "Read-only drift / expiry / orphan detection; emits metrics and alerts",
		Long: `Runs on a schedule (CI cron, k8s CronJob, systemd timer). Diffs
each BIG-IP's installed certs against state and manifests; classifies findings
as drift, orphan-on-device, missing-on-device, near-expiry, or
recent-renewal-failure. Emits Prometheus metrics on the port given by
--metrics-addr and exits non-zero if any class triggers an alert threshold.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errNotImplemented("reconcile")
		},
	}
	cmd.Flags().String("metrics-addr", ":9090", "Listen address for /metrics; empty disables")
	cmd.Flags().Duration("expiry-warn", 14*24*time.Hour, "Warn when a cert's notAfter is within this duration")
	cmd.Flags().Bool("once", true, "Run a single reconcile pass and exit (set false to loop)")
	return cmd
}
