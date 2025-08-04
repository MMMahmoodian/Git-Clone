package cmd

import (
	"github.com/MMMahmoodian/Simple-VCS/internal/testutil"
	"strings"
	"testing"
)

func TestAddCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"no args", []string{"add"}, "Running add command"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := testutil.RunSubCommand(t, addCommand, tt.args)
			if err != nil {
				t.Fatalf("command failed: %v", err)
			}
			if !strings.Contains(output, tt.expected) {
				t.Errorf("got %q, want to contain %q", output, tt.expected)
			}
		})
	}
}
