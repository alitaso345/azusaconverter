package azusaconverter

import (
	"strings"
)

var dictionary = map[string]string{
	"言ったの":   "言ったん",
	"来たのか":   "来たんか",
	"聞こえるの":  "聞こえんねん",
	"だから":    "やから",
	"すごい":    "めっちゃ",
	"おいしそう":  "うまそう",
	"してなかった": "してへんかった",
}

func Convert(text string) string {
	for capital, kyoto := range dictionary {
		text = strings.ReplaceAll(text, capital, kyoto)
	}
	return text
}
