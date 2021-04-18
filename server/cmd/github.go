package cmd

import (
	"istio-redirector/pkg/github"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(githubCmd)
}

var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "Create PR on Github from file",
	RunE: func(cmd *cobra.Command, args []string) error {

		github.Create([]byte("test from cmd"), "test-from-cli")

		return nil
	},
}