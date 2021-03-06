package command_test

import (
	. "github.com/ForceCLI/force/command"
	"github.com/bmizerany/assert"
	"testing"
)

// test that all avialable commands come with at least a name and short usage information
func TestUsage(t *testing.T) {
	for _, cmd := range Commands {
		assert.NotEqual(t, cmd.Name(), "")
		assert.NotEqual(t, cmd.Short, "")
	}
}
