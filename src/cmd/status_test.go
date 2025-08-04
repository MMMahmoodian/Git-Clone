package cmd

import (
	"strings"
	"testing"

	"github.com/MMMahmoodian/Simple-VCS/internal/testutil"
)

func TestStatusCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"no args", []string{"status"}, "Running status command"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := testutil.RunSubCommand(t, statusCommand, tt.args)
			if err != nil {
				t.Fatalf("command failed: %v", err)
			}
			if !strings.Contains(output, tt.expected) {
				t.Errorf("got %q, want to contain %q", output, tt.expected)
			}
		})
	}
}
