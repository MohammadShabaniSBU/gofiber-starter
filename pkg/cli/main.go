package cli

import (
	"os"

	"bitbucket.org/dyfrag-internal/mass-media-core/pkg/cli/serve"
	"github.com/spf13/cobra"
)

// ExitFailure status code.
const ExitFailure = 1

func Execute() {
	//nolint: exhaustruct
	root := &cobra.Command{
		Use:   "monitor",
		Short: "final project of snapp course",
	}

	root.AddCommand(serve.New())

	if err := root.Execute(); err != nil {
		os.Exit(ExitFailure)
	}
}
