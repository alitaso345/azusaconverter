package azusaconverter

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "言ったの？/言ったん？",
			input:    "...だから、部活動辞めたいって言ったの？",
			expected: "...だから、部活動辞めたいって言ったん？",
		},
	}
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, Convert(ts.input), ts.expected)
		})
	}
}
