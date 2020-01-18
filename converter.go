package azusaconverter

import (
	"regexp"
	"strings"
)

var dictionary = map[string]string{
	"言ったの":        "言ったん",
	"来たのか":        "来たんか",
	"聞こえるの":       "聞こえんねん",
	"だから":         "やから",
	"すごいおいしそう":    "めっちゃうまそう",
	"してなかった":      "してへんかった",
	"着てらしたのって":    "着てはったんって",
	"でしょ":         "やろ",
	"でしょ？":        "やろ？",
	"使わない":        "使わん",
	"狭いんだよな":      "狭いねんな",
	"ここだと":        "ここやと",
	"こんなもんだって":    "こんなもんやって",
	"んだって":        "んやって",
	"やめちゃうの":      "やめちゃうん",
	"伝えるね":        "伝えるな",
	"くださった":       "くれはった",
	"わたし":         "うち",
	"私が":          "うちが",
	"行ったんだと":      "行ったんやと",
	"になったよ":       "になったわ",
	"聞くの？":        "聞くん？",
	"大丈夫かな":       "大丈夫なんやろか",
	"ほんと":         "ほんま",
	"だったから":       "やったから",
	"だったよね":       "やったよな",
	"いいと思う":       "ええと思う",
	"できてないよ":      "できてへんよ",
	"だめだよ？":       "あかんで？",
	"考えてるんだなあ":    "考えてんねんなあ",
	"だめなの":        "あかんねんて",
	"いないと":        "おらんと",
	"してほしくないの":    "してほしくないねんなあ",
	"しちゃってるの":     "しちゃってんねん",
	"買うの？":        "買うん？",
	"なの？":         "なん？",
	"やってたの":       "やってたん",
	"いないで":        "おらんで",
	"いいの？":        "ええの？",
	"てなかったの":      "てなかったん",
	"美味しいのかな":     "美味しいんかな",
	"負けなんだよね":     "負けやねんなあ",
	"なるよ":         "なんで",
	"あんまり":        "あんま",
	"湧かないなぁ":      "湧かへんなぁ",
	"好きじゃなく":      "好きやなく",
	"どうしたの":       "どうしたん",
	"したよ":         "したわ",
	"いまさらだな":      "いまさらやな",
	"わからないし":      "わからへんし",
	"そういうんじゃない":   "そういうんやない",
	"あるよ":         "あんねんで",
	"思い出したんだけど":   "思い出してんけど",
	"言ってたんだ":      "言っててん",
	"だけどね":        "やけどね",
	"なんだけど":       "やねんけど",
	"言ってるから":      "言ってんねんから",
	"わからないじゃん":    "わからんやん",
	"ないんだし":       "ないんやし",
	"やろうよ":        "やろうや",
	"褒めてない":       "褒めてへん",
	"言われたくないよ":    "言われたくないわ",
	"じゃないよ":       "ちゃうよ",
	"言われたんだけど":    "言われてんけど",
	"取っておいたの":     "取っておいてん",
	"なにも":         "なんも",
	"言えなかった":      "言えへんかった",
	"話してたの？":      "話してたん？",
	"そうなの":        "そうやねん",
	"行ったの":        "行ったん",
	"前だよ":         "前やわ",
	"だけどな":        "やけどな",
	"いいじゃん":       "ええやん",
	"できるのは":       "できんのは",
	"うっとうしそうだな":   "うっとうしそうやな",
	"思ってたの":       "思っててん",
	"言わないでね？":     "言わんといてな？",
	"褒めないで":       "褒めんといて",
	"吐かないでよ":      "吐かんといてよ",
	"照れるな":        "照れるわ",
	"やめといたほうがいい":  "やめといたほうがええ",
	"食べたの":        "食べたん",
	"あげられないよ":     "あげれへんで",
	"帰ってこない":      "帰ってこおへん",
	"入ったの":        "入ってん",
	"遊べないんだ":      "遊べへんねん",
	"信じてない":       "信じてへん",
	"いなかった":       "おらんかった",
	"大事なんだよ":      "大事やねんて",
	"わけじゃないから":    "わけちゃうから",
	"決めたのかと":      "決めたんかと",
	"いるんだもん":      "おるんやもん",
	"なれない":        "なれへん",
	"話してない":       "話してへん",
	"決めてないけど":     "決めてへんけど",
	"いけないの":       "あかんの",
	"思ったの？":       "思ったん？",
	"どうしたらいいの":    "どうしたらいいんさ",
	"しょうがない":      "しゃあない",
	"伸びるんだもん":     "伸びんねんもん",
	"にいれば":        "におれば",
	"すぎない？":       "すぎひん？",
	"嫌いだった":       "嫌いやった",
	"見れない":        "見られへん",
	"そうじゃなくて":     "そうやなくて",
	"よく覚えて":       "よう覚えて",
	"そうだった":       "そうやった",
	"そうだったら":      "そうやったら",
	"できない自分":      "できひん自分",
	"許せない":        "許せへん",
	"そうだね":        "そうやな",
	"そうだよね":       "そうやんなあ",
	"そうなんだ":       "そうなんや",
	"うれしいよ":       "うれしいわ",
	"うれしいね":       "うれしいわ",
	"あったんじゃないの":   "あったんちゃうの",
	"嫌だなあ":        "嫌やなあ",
	"やってやろう":      "やったろう",
	"そうだけど":       "そうやけど",
	"聞いたの？":       "聞いたん？",
	"ダメ":          "あかん",
	"思わない":        "思わへん",
	"心配なの":        "心配やねん",
	"やってみたらいいよ":   "やってみたらええよ",
	"だったんだよね":     "やってんか",
	"それなのに":       "それやのに",
	"なんだろう":       "なんやろう",
	"それならよかった":    "それやったらよかったわ",
	"当たり前じゃん":     "当たり前やん",
	"大丈夫なの？":      "大丈夫なん？",
	"だろうけど":       "やろうけど",
	"あるんじゃないか":    "あるんやないか",
	"あったんじゃない？":   "あったんちゃうん？",
	"そんなこと言って":    "そんなこと言うて",
	"手を抜いて":       "手ぇ抜いて",
	"かもしれないよ":     "かもしれんで",
	"そんな性格だと":     "そんな性格やと",
	"ならないでしょ":     "ならんやろ",
	"されてるんでしょ":    "されてるんやろ",
	"できないでしょ":     "できひんやろ",
	"ありえないよ":      "ありえへんよ",
	"いいんだよ":       "ええねんで",
	"気にしないで":      "気にせんと",
	"かなあ":         "やなあ",
	"なるんだもん":      "なるんやもん",
	"したじゃん":       "したやん",
	"いない":         "おらん",
	"だし":          "やし",
	"じゃないし":       "とちゃうし",
	"だめだ、":        "あかん、",
	"言わなきゃだめ":     "言わなあかん",
	"待ってよ":        "待ってや",
	"ことはないよ":      "ことはないで",
	"ようよ":         "ようや",
	"感じだね":        "感じやな",
	"できないこと":      "できないこと",
	"しないのは":       "しいひんのは",
	"わからなくて":      "わからんくて",
	"やるもの":        "やるもん",
	"なってきたよ":      "なってきたわ",
	"じゃないんだ？":     "と違うんや？",
	"気がする":        "気ぃする",
	"どうもしないよ":     "どうもしいひんよ",
	"行ってたの？":      "行ってたん？",
	"呼び名だよ":       "呼び名やな",
	"するんだろ":       "するんやろ",
	"なにもない":       "なんもない",
	"なりそうだもん":     "なりそうやもん",
	"聞いてたよ":       "聞いてたで",
	"なっちゃうんだよなあ":  "なっちゃうねんなあ",
	"なんなの":        "なんやの",
	"ねえねえ":        "なあなあ",
	"なかったら":       "へんかったら",
	"してるよ":        "してるわ",
	"すごい期待してる":    "めっちゃ期待してる",
	"はやく":         "はよ",
	"いかないと":       "いかんと",
	"終わるよ":        "終わんで",
	"てたよね":        "てたやん",
	"決めたもの":       "決めたもん",
	"ほかのもの":       "ほかのもん",
	"入らなくなる":      "入らんくなる",
	"ではない":        "ちゃう",
	"気にしないけど":     "気にせえへんけど",
	"べつにいいじゃんか":   "べつにええやんか",
	"べつにいいって":     "べつにええって",
	"思ってない":       "思ってへん",
	"聞かないで":       "聞かんで",
	"でいいよ":        "でいいねん",
	"そこから":        "そっから",
	"やらなきゃいけなくなる": "やらなあかんくなる",
	"なるじゃんか":      "なるやんか",
	"なんでしょ":       "なんやろ",
	"だけだから":       "だけやから",
	"かもしれないし":     "かもしれへんし",
	"やるじゃんか":      "やるやんか",
	"言ってたじゃん":     "言うてたやん",
	"あったじゃんか":     "あったやんか",
	"やりなよ":        "やりーな",
	"いけないじゃん":     "あかんやん",
	"ごめんね":        "ごめんな",
	"よかったと思う":     "よかった思う",
	"言ってない":       "言ってへん",
	"辞めるの":        "辞めるん",
	"聞かないな":       "聞かへんな",
	"じゃないけど":      "やないけど",
	"怪しいんだけど":     "怪しいねんけど",
	"だよねー":        "やんなー",
	"あげないとね":      "あげへんとね",
	"だろうね":        "やろうな",
	"だと思う":        "やと思う",
	"カッコよかったよね":   "カッコよかったやんな",
	"ことだよ":        "ことやな",
	"アレだけど":       "アレやけど",
	"上手いじゃん":      "上手いやん",
	"じゃないかな":      "ちゃうかな",
	"取れてないから":     "取れてへんから",
	"直さないと":       "直さんと",
	"厳しいと思うよ":     "厳しい思うわ",
	"思ってなかった":     "思ってへんかった",
	"思わなかった":      "思わんかった",
	"ことばっか言って":    "ことばっか言うて",
	"じゃないんだから":    "やないねんから",
	"気づかないと":      "気づかんと",
	"だめだよ":        "あかんで",
	"ヤバくない":       "ヤバない",
	"帰るの？":        "帰るん？",
	"しれないね":       "しれへんね",
	"やってない":       "やってへん",
	"なんだから":       "やねんから",
	"やっぱり、":       "やっぱ、",
	"すごいんじゃない？":   "すごいんちゃう？",
	"できるん":        "できんの",
	"だろうか":        "やろか",
	"よくわからない":     "ようわからん",
	"合わせてね":       "合わせてな",
	"わかってない":      "わかってへん",
	"そんなの":        "そんなん",
	"頑張ってない":      "頑張ってへん",
	"好きだよ":        "好きやで",
	"飲みたいな":       "飲みたいねんな",
	"嫌じゃない":       "嫌やない",
	"寝れなかったんだけど":  "寝れへんかってんけど",
	"やりたいの":       "やりたいねん",
	"だけだし":        "だけやし",
	"呼んでる":        "呼んでん",
	"ずるいじゃん":      "ずっこいやん",
	"思ってないんだから":   "思ってへんねんから",
	"思ってるんだけど":    "思ってんねんけど",
	"できないんだよなあ":   "できひんねんなあ",
	"しないといけない":    "しなあかん",
	"べつにいいけど":     "べつにええけど",
	"説明するの":       "説明すんの",
	"だったけど":       "やったけど",
	"言ってなかった":     "言ってへんかった",
	"だったので":       "やったんで",
	"寂しかったの":      "寂しかってん",
	"だと":          "やと",
	"わからないんだよね":   "わからんねんな",
	"話したの？":       "話したん？",
	"言ってんの":       "言うてんの",
	"ばか":          "アホ",
	"だったの":        "やったん",
	"寂しくない":       "寂しない",
	"かもしれないじゃん":   "かもしれんやん",
	"行かなかったの":     "行かんかったん",
	"考えてるんだね":     "考えてんねんな",
	"足りない":        "足りひん",
	"嫌だ":          "嫌や",
	"怖かったの":       "怖かってん",
	"バカ":          "アホ",
	"求めてない":       "求めてへん",
	"いいかな":        "ええかな",
	"なんだー":        "なんやー",
	"たのかな":        "たんやろか",
	"なったんだから":     "なってんから",
	"らしいよ":        "らしいやん",
	"合ってないじゃん":    "合ってないねん",
	"やってられない":     "やってられへん",
	//"なんだ": "やねん",
	//"できなかった": "できひんかった",
	//"いいから": "ええから",
	//"だったのは": "やったのは",
	//"だけじゃない": "だけちゃう",
	//"なんだと": "なんやと",
	//"関わらない": "関わらん",
	//"言ってるの": "言ってんの",
	//"違うの？": "違うん？",
	//"だったんだ": "やってんな",
	//"いいんじゃない": "いいんちゃう",
	//"褒めてくれたの": "褒めてくれてん",
	//"なるんだから": "なるんやから",
	//"ばっかりだし": "ばっかやし",
	//"しれないよ": "しれへんで",
	//"わからないもん": "わからへんもん",
	//"〜だから": "〜やから",
	//"突き放したのは": "突き放したんは",
	//"思われてる": "思ってはる",
	//"どこにいるの": "どこにおんの",
	//"くださる": "くれはる",
	//"だったなんて": "やったなんて",
	//"信じられなくて": "信じられへんくて",
	//"どうでもいいよね": "どうでもええねんなあ",
	//"大変だね": "大変やんな",
	//"だろうし": "やろうし",
	//"変わってない": "変わってへん",
	//"してくださった": "してくれてはった",
	//"だよー": "やなー",
	//"謝らないで": "謝らんで",
	//"ないし": "へんし",
	//"怒ってない": "怒ってへん",
	//"だろうな": "やろうな",
	//"〜だな": "〜やな",
	//"大丈夫か": "大丈夫やろうか",
	//"大丈夫だって": "大丈夫やって",
	//"いたの": "いてん",
	//"あるんだけど": "あんねんけど",
	//"考えとくね": "考えとくわ",
	//"考えとくよ": "考えとくわ",
	//"言いすぎだって": "言いすぎやって",
	//"あるのは": "あんのは",
	//"ほうだと": "ほうやと",
	//"ないじゃん": "ないやん",
	//"逃げるの？": "逃げるん？",
	//"いる": "おる",
	//"変わらない": "変わらん",
	//"じゃない？": "ちゃう？",
	//"嫌いなの？": "嫌いなん？",
	//"なのに": "やのに",
	//"しないといけない": "せなあかん",
	//"違うの": "ちゃうの",
	"だったら": "やったら",
	//"負けない": "負けへん",
	//"やってるんだから": "やってんから",
	//"なんだろ": "なんやろ",
	//"言ってるん": "言うてるん",
	//"カッコいいよ": "カッコいいわ",
	"本当": "ほんま",
	//"よくない": "あかん",
	//"話しかけないでおいて": "話しかけんといて",
	//"時間ないし": "時間ないねんし",
	//"おかしいの": "おかしいねん",
	//"わかるんだし": "わかんのやし",
	//"くれないの": "くれへんの",
	//"いるの": "いんの",
	//"好きなの": "好きやねん",
	//"裏切ってない": "裏切ってへん",
	//"だと、": "やと、",
	//"してるんだ": "してるんや",
	//"なの、": "やねん、",
	//"知ってるの": "知ってんの",
	//"上手だな": "上手やな",
	//"友達でしょ？": "友達やろ？",
	"すごい好き": "めっちゃ好き",
	//"そうそう": "そうやねん",
}

var dictionary2 = map[string]string{
	"じゃないんやから": "やないねんから", // "だから": "やから" との競合対策
	"なんやから":    "やねんから",
	"へんんやから":   "へんねんから",
	"なったんやから":  "なってんから",
	"できないやろ":   "できひんやろ", // "でしょ" : "やろ" との競合対策
	"ならないやろ":   "ならんやろ",
	"だよ$":      "やで",   // 末尾の"やで"対策
	"なんやの$":    "なんなん", // 末尾の"なんなの": "なんやの" との競合対策
	"じゃんか$":    "やんか",
	"だね$":      "やな",
}

func Convert(text string) string {
	for capital, kyoto := range dictionary {
		text = strings.ReplaceAll(text, capital, kyoto)
	}
	for mismatch, correct := range dictionary2 {
		rep := regexp.MustCompile(mismatch)
		text = rep.ReplaceAllString(text, correct)
	}
	return text
}
