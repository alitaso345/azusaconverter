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
			input:    "部活動辞めたいって言ったの？",
			expected: "部活動辞めたいって言ったん？",
		},
		{
			name:     "来たのか/来たんか",
			input:    "...ついにここまで来たのかあ",
			expected: "...ついにここまで来たんかあ",
		},
		{
			name:     "聞こえるの/聞こえんねん",
			input:    "『トゥ』の勢いが強いのに『ク』が不明瞭だから、変なバランスに聞こえるの",
			expected: "『トゥ』の勢いが強いのに『ク』が不明瞭やから、変なバランスに聞こえんねん",
		},
		{
			name:     "だから/やから",
			input:    "あみかの場合、初心者とか関係なく天然だから心配されてるんでしょ",
			expected: "あみかの場合、初心者とか関係なく天然やから心配されてるんでしょ",
		},
		{
			name:     "すごい/めっちゃ、おいしそう/うまそう",
			input:    "あー、すごいおいしそう",
			expected: "あー、めっちゃうまそう",
		},
		{
			name:     "してなかった/してへんかった",
			input:    "あー、そういや、その説明してなかったね",
			expected: "あー、そういや、その説明してへんかったね",
		},
	}
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, Convert(ts.input), ts.expected)
		})
	}
}
