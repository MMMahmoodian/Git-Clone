package cmd

import (
	"strings"
	"testing"

	"github.com/MMMahmoodian/Simple-VCS/internal/testutil"
)

func TestInitCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"no args", []string{"init"}, "Running init command"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := testutil.RunSubCommand(t, initCommand, tt.args)
			if err != nil {
				t.Fatalf("command failed: %v", err)
			}
			if !strings.Contains(output, tt.expected) {
				t.Errorf("got %q, want to contain %q", output, tt.expected)
			}
		})
	}
}
