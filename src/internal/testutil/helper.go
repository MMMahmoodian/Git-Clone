package testutil

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func RunSubCommand(t *testing.T, sub *cobra.Command, args []string) (string, error) {
	t.Helper()

	var buf bytes.Buffer

	root := &cobra.Command{Use: "test"}
	cmd := *sub
	root.AddCommand(&cmd)

	root.SetOut(&buf)
	root.SetErr(&buf)
	root.SetArgs(args)

	err := root.Execute()
	return buf.String(), err
}
