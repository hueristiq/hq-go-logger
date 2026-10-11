package colorizer

import (
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	hqgologgerlevels "github.com/hueristiq/hq-lib-logger-go/levels"
)

func TestFatihColorizerColoredLevels(t *testing.T) {
	t.Parallel()

	c := NewFatihColorizer()
	require.NotNil(t, c)

	// Sprint consults the package-global color.NoColor, which is true when
	// standard output is not a terminal (as in test runs). Enable the
	// per-instance override so the ANSI styling is deterministic.
	for _, colored := range []*color.Color{c.fatal, c.errColor, c.info, c.warn, c.debug} {
		colored.EnableColor()
	}

	// LevelFatal and LevelError deliberately share one styling; every colored
	// level must still produce non-empty, colorized output.
	tests := []struct {
		name  string
		level hqgologgerlevels.Level
	}{
		{"fatal", hqgologgerlevels.LevelFatal},
		{"error", hqgologgerlevels.LevelError},
		{"info", hqgologgerlevels.LevelInfo},
		{"warn", hqgologgerlevels.LevelWarn},
		{"debug", hqgologgerlevels.LevelDebug},
	}

	outputs := make(map[string]struct{}, len(tests))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := c.Colorize("INF", tt.level)

			assert.NotEmpty(t, got)
			assert.Contains(t, got, "INF")
			assert.NotEqual(t, "INF", got, "colored levels must not pass through unchanged")
		})
	}

	for _, tt := range tests {
		outputs[c.Colorize("INF", tt.level)] = struct{}{}
	}

	assert.Greater(t, len(outputs), 1, "colored levels must not all share a single output")
}

func TestFatihColorizerSilentAndInvalidPassThrough(t *testing.T) {
	t.Parallel()

	c := NewFatihColorizer()
	require.NotNil(t, c)

	for _, level := range []hqgologgerlevels.Level{
		hqgologgerlevels.LevelSilent,
		hqgologgerlevels.Level(-1),
		hqgologgerlevels.Level(99),
	} {
		assert.Equal(t, "INF", c.Colorize("INF", level))
	}
}
