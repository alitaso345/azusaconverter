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
			expected: "あみかの場合、初心者とか関係なく天然やから心配されてるんやろ",
		},
		{
			name:     "すごいおいしそう/めっちゃうまそう",
			input:    "あー、すごいおいしそう",
			expected: "あー、めっちゃうまそう",
		},
		{
			name:     "してなかった/してへんかった",
			input:    "あー、そういや、その説明してなかったね",
			expected: "あー、そういや、その説明してへんかったね",
		},
		{
			name:     "着てらしたのって/着てはったんって",
			input:    "あぁ、他の先輩らが違う名前のジャージ着てらしたのって、そういう理由なんですね",
			expected: "あぁ、他の先輩らが違う名前のジャージ着てはったんって、そういう理由なんですね",
		},
		{
			name:     "でしょ？/やろ？",
			input:    "あー、中学の子らが来る三週間後のイベントって、シングやるんでしょ？",
			expected: "あー、中学の子らが来る三週間後のイベントって、シングやるんやろ？",
		},
		{
			name:     "使わない/使わん",
			input:    "あー、普段使わない筋肉使ったせいで身体がやばい",
			expected: "あー、普段使わん筋肉使ったせいで身体がやばい",
		},
		{
			name:     "狭いんだよな/狭いねんな",
			input:    "あそこの風呂、狭いんだよなあ",
			expected: "あそこの風呂、狭いねんなあ",
		},
		{
			name:     "ここだと/ここやと",
			input:    "あと、ここだと琉球の民族衣装着て写真撮れるんだって",
			expected: "あと、ここやと琉球の民族衣装着て写真撮れるんやって",
		},
		{
			name:     "こんなもんだって/こんなもんやって",
			input:    "どこも強いとこはこんなもんだって",
			expected: "どこも強いとこはこんなもんやって",
		},
		{
			name:     "やめちゃうの/やめちゃうん",
			input:    "あとちょっと頑張ればもっとよくなるのに、なんでそこでやめちゃうの？って思うこと結構あります",
			expected: "あとちょっと頑張ればもっとよくなるのに、なんでそこでやめちゃうん？って思うこと結構あります",
		},
		{
			name:     "伝えるね/伝えるな",
			input:    "あとで伝えるね",
			expected: "あとで伝えるな",
		},
		{
			name:     "くださった/くれはった",
			input:    "あの、いまの話、わたしだから教えてくださったんですか？",
			expected: "あの、いまの話、うちやから教えてくれはったんですか？",
		},
		{
			name:     "わたし/うち、行ったんだと/行ったんやと",
			input:    "わたし、てっきり久美子も南宇治に行ったんだと思ってた",
			expected: "うち、てっきり久美子も南宇治に行ったんやと思ってた",
		},
		{
			name:     "私/うち",
			input:    "私が心配性なんじゃなくて、相手があみかだから心配しちゃうの",
			expected: "うちが心配性なんじゃなくて、相手があみかやから心配しちゃうの",
		},
		{
			name:     "になったよ/になったわ",
			input:    "あのとき、先輩と目が合って、心臓が止まりそうになったよ",
			expected: "あのとき、先輩と目が合って、心臓が止まりそうになったわ",
		},
		{
			name:     "聞くの？/聞くん？",
			input:    "あみか、なんで的場に聞くの？",
			expected: "あみか、なんで的場に聞くん？",
		},
		{
			name:     "大丈夫かな/大丈夫なんやろか",
			input:    "あみか、ほんとに大丈夫かな",
			expected: "あみか、ほんまに大丈夫なんやろか",
		},
		{
			name:     "ほんと/ほんま、だったから/やったから",
			input:    "ステップやれるの、ほんと楽しみだったから",
			expected: "ステップやれるの、ほんま楽しみやったから",
		},
		{
			name:     "だったよね/やったよな",
			input:    "あみか、確かまだポジション位置不安定だったよね？",
			expected: "あみか、確かまだポジション位置不安定やったよな？",
		},
		{
			name:     "いいと思う/ええと思う",
			input:    "あみかが興味あるならやってもいいと思うけど",
			expected: "あみかが興味あるならやってもええと思うけど",
		},
		{
			name:     "できてないよ/できてへんよ",
			input:    "あみかさ、『ク』ができてないよ",
			expected: "あみかさ、『ク』ができてへんよ",
		},
		{
			name:     "だめだよ？/あかんで？",
			input:    "あみかさ、知らん人からお菓子もらっても食べたらだめだよ？",
			expected: "あみかさ、知らん人からお菓子もらっても食べたらあかんで？",
		},
		{
			name:     "考えてるんだなあ/考えてんねんなあ",
			input:    "あみかってば、ちゃんと考えてるんだなあ",
			expected: "あみかってば、ちゃんと考えてんねんなあ",
		},
		{
			name:     "だめなの/あかんねんて、いないと/おらんと",
			input:    "あみかには、わたしがいないとだめなの！",
			expected: "あみかには、うちがおらんとあかんねんて！",
		},
		{
			// Note: 原作では「あんまり」を使っているが、 あんまり -> あんま の変換パターンを優先する
			name:     "してほしくないの/してほしくないねんなあ",
			input:    "あみかにはあんまりそういう経験してほしくないの",
			expected: "あみかにはあんまそういう経験してほしくないねんなあ",
		},
		{
			name:     "しちゃってるの/しちゃってんねん",
			input:    "あみかの場合、フライングしちゃってるの",
			expected: "あみかの場合、フライングしちゃってんねん",
		},
		{
			name:     "買うの？/買うん？",
			input:    "あみかはいつ楽器買うの？",
			expected: "あみかはいつ楽器買うん？",
		},
		{
			name:     "なの？/なん？",
			input:    "あみかはさ、ここがお気に入りなの？",
			expected: "あみかはさ、ここがお気に入りなん？",
		},
		{
			name:     "やってたの/やってたん",
			input:    "あみかは中学時代、何やってたの？",
			expected: "あみかは中学時代、何やってたん？",
		},
		{
			//Note: 原作では「いいの？」となっているが、今回は いいの -> ええの の変換パターンを優先する
			name:     "いないで/おらんで",
			input:    "あれ、カラーガードの子らといないでいいの？",
			expected: "あれ、カラーガードの子らとおらんでええの？",
		},
		{
			name:     "いいの？/ええの？",
			input:    "あれ、そういえばシングの練習はいいの？",
			expected: "あれ、そういえばシングの練習はええの？",
		},
		{
			name:     "なかったの/なかったん",
			input:    "あれ、まだ人数集まってなかったの？",
			expected: "あれ、まだ人数集まってなかったん？",
		},
		{
			name:     "美味しいのかな/美味しいんかな",
			input:    "あれってほんとに美味しいのかな",
			expected: "あれってほんまに美味しいんかな",
		},
		{
			name:     "負けなんだよね/負けやねんなあ",
			input:    "あれなー、なんか冷静に考えちゃうと負けなんだよね",
			expected: "あれなー、なんか冷静に考えちゃうと負けやねんなあ",
		},
		{
			name:     "なるよ/なんで",
			input:    "あんまり急いで食べたら豚になるよ",
			expected: "あんま急いで食べたら豚になんで",
		},
		{
			name:     "あんまり/あんま",
			input:    "うち、あんまりそういうのに興味ないから",
			expected: "うち、あんまそういうのに興味ないから",
		},
		{
			name:     "湧かないなぁ/湧かへんなぁ",
			input:    "あんま実感湧かないなぁ",
			expected: "あんま実感湧かへんなぁ",
		},
		{
			// Note: 原作では「いいじゃん」を使う場面もあるが、いいじゃん -> ええやん の変換を優先させる
			name:     "いいじゃん/ええやん",
			input:    "いいじゃん",
			expected: "ええやん",
		},
		{
			name:     "好きじゃなく/好きやなく",
			input:    "いいじゃん、好きじゃなくたって",
			expected: "ええやん、好きやなくたって",
		},
		{
			name:     "どうしたの/どうしたん",
			input:    "いきなりどうしたの？",
			expected: "いきなりどうしたん？",
		},
		{
			name:     "したよ/したわ",
			input:    "いきなり変なこと言うからびっくりしたよ",
			expected: "いきなり変なこと言うからびっくりしたわ",
		},
		{
			name:     "かな/やな",
			input:    "イタリアの音楽家オリヴァドーティが作ったスクールバンド向けの序曲で、吹奏楽の定番だね",
			expected: "イタリアの音楽家オリヴァドーティが作ったスクールバンド向けの序曲で、吹奏楽の定番やな",
		},
		{
			name:     "わからないし/わからへんし",
			input:    "いつスライド落っことすかわからないし",
			expected: "いつスライド落っことすかわからへんし",
		},
		{
			name:     "そういうんじゃない/そういうんやない",
			input:    "いまのはそういうんじゃないって",
			expected: "いまのはそういうんやないって",
		},
		{
			name:     "あるよ/あんねんで",
			input:    "いま頑張ったら、的場だってAに行ける可能性はあるよ",
			expected: "いま頑張ったら、的場だってAに行ける可能性はあんねんで",
		},
		{
			name:     "思い出したんだけど/思い出してんけど、言ってたんだ/言っててん",
			input:    "いま思い出したんだけどさ、そういえば芹菜もガラス工芸館行きたいって言ってたんだ",
			expected: "いま思い出してんけどさ、そういえば芹菜もガラス工芸館行きたいって言っててん",
		},
		{
			name:     "だけどね/やけどね",
			input:    "いま北宇治に行ってるんだけどね",
			expected: "いま北宇治に行ってるんやけどね",
		},
		{
			name:     "なんだけど/やねんけど",
			input:    "いや、しょうもないことやねんけどさ",
			expected: "いや、しょうもないことやねんけどさ",
		},
		{
			name:     "言ってるから/言ってんねんから",
			input:    "いや、そうは言うけどさ、べつにわたしが問題ないって言ってるからよくない？",
			expected: "いや、そうは言うけどさ、べつにうちが問題ないって言ってんねんからよくない？",
		},
		{
			name:     "わからないじゃん/わからんやん",
			input:    "いや、わからないじゃん",
			expected: "いや、わからんやん",
		},
		{
			name:     "ないんだし/ないんやし、やろうよ/やろうや",
			input:    "いや、誰がAに行くとかまだ決まってないんだし、オーディションまでにできることはやろうよ",
			expected: "いや、誰がAに行くとかまだ決まってないんやし、オーディションまでにできることはやろうや",
		},
		{
			name:     "褒めてない/褒めてへん",
			input:    "いや、褒めてないけどな",
			expected: "いや、褒めてへんけどな",
		},
		{
			name:     "言われたくないよ/言われたくないわ",
			input:    "いやいや、あみかには言われたくないよ",
			expected: "いやいや、あみかには言われたくないわ",
		},
		{
			name:     "じゃないよ/ちゃうよ",
			input:    "いやいや、ゆっくりしてる場合じゃないよ",
			expected: "いやいや、ゆっくりしてる場合ちゃうよ",
		},
		{
			name:     "言われたんだけど/言われてんけど、取っておいたの/取っておいてん",
			input:    "いろんな子に私物くださいって言われたんだけど、ジャージだけはアンタのために取っておいたの",
			expected: "いろんな子に私物くださいって言われてんけど、ジャージだけはアンタのために取っておいてん",
		},
		{
			name:     "なにも/なんも、言えなかった/言えへんかった",
			input:    "わたし、なんも言えへんかったんです",
			expected: "うち、なんも言えへんかったんです",
		},
		{
			name:     "してたの/してたん",
			input:    "うちが黄檗で降りたあと、二人でそんな話してたの？",
			expected: "うちが黄檗で降りたあと、二人でそんな話してたん？",
		},
		{
			name:     "そうなの/そうやねん",
			input:    "うちさ、昔からそうなの",
			expected: "うちさ、昔からそうやねん",
		},
		{
			name:     "行ったの/行ったん、前だよ/前やわ",
			input:    "うちの学校も、コンクールのほうで最後に全国行ったの、めちゃくちゃ前だよ",
			expected: "うちの学校も、コンクールのほうで最後に全国行ったん、めちゃくちゃ前やわ",
		},
		{
			name:     "だけどな/やけどな",
			input:    "うちはめっちゃ楽しみだけどな",
			expected: "うちはめっちゃ楽しみやけどな",
		},
		{
			name:     "友達でしょ？/友達やろ？",
			input:    "うちら、友達でしょ？",
			expected: "うちら、友達やろ？",
		},
		{
			name:     "できるのは/できんのは",
			input:    "うちらにできるのは名工が弱くなってるように祈ること！",
			expected: "うちらにできんのは名工が弱くなってるように祈ること！",
		},
		{
			name:     "うっとうしそうだな/うっとうしそうやな、思ってたの/思っててん",
			input:    "うっとうしそうだなとは思ってたの",
			expected: "うっとうしそうやなとは思っててん",
		},
		{
			name:     "言わないでね？/言わんといてな？",
			input:    "うれしいけど、ほかの先輩の前ではあんまりそういうこと言わないでね？",
			expected: "うれしいけど、ほかの先輩の前ではあんまそういうこと言わんといてな？",
		},
		{
			name:     "すごい好き/めっちゃ好き",
			input:    "うん、すごい好き",
			expected: "うん、めっちゃ好き",
		},
		{
			name:     "そうそう/そうやねん",
			input:    "うん、そうなの、中学時代の友達",
			expected: "うん、そうやねん、中学時代の友達",
		},
		{
			name:     "褒めないで/褒めんといて",
			input:    "えー、あんまり褒めないで",
			expected: "えー、あんま褒めんといて",
		},
		{
			name:     "吐かないでよ/吐かんといてよ",
			input:    "えー、こんなとこで吐かないでよ？",
			expected: "えー、こんなとこで吐かんといてよ？",
		},
		{
			name:     "照れるな/照れるわ",
			input:    "えー、そう言われると、ほんと照れるな",
			expected: "えー、そう言われると、ほんま照れるわ",
		},
		{
			name:     "やめといたほうがいい/やめといたほうがええ",
			input:    "えー、やめといたほうがいいって",
			expected: "えー、やめといたほうがええって",
		},
		{
			name:     "食べたの/食べたん",
			input:    "えっ、もう食べたの？",
			expected: "えっ、もう食べたん？",
		},
		{
			name:     "いまさらだな/いまさらやな",
			input:    "おお、マジでいまさらだな",
			expected: "おお、マジでいまさらやな",
		},
		{
			name:     "あげられないよ/あげれへんで",
			input:    "お世辞言ったってこんなもんしかあげられないよ",
			expected: "お世辞言ったってこんなもんしかあげれへんで",
		},
		{
			name:     "帰ってこない/帰ってこおへん",
			input:    "お母さん、まだ帰ってこないのかな",
			expected: "お母さん、まだ帰ってこおへんのかな",
		},
		{
			name:     "入ったの/入ってん",
			input:    "お母さんが仕事で忙しいときにどこで時間を潰すかって話になって、それで金管バンドに入ったの",
			expected: "お母さんが仕事で忙しいときにどこで時間を潰すかって話になって、それで金管バンドに入ってん",
		},
		{
			name:     "遊べないんだ/遊べへんねん",
			input:    "お盆は用事あるから遊べないんだ",
			expected: "お盆は用事あるから遊べへんねん",
		},
		{
			name:     "信じてない/信じてへん、やってられない/やってられへん",
			input:    "けど、そう信じてないとこっちもやってられないじゃんか",
			expected: "けど、そう信じてへんとこっちもやってられへんやんか",
		},
		{
			name:     "いなかった/おらんかった",
			input:    "けど、卑怯な手を使ってまでポジションを取ろうとしてる人はいなかった",
			expected: "けど、卑怯な手を使ってまでポジションを取ろうとしてる人はおらんかった",
		},
		{
			name:     "大事なんだよ/大事やねんて",
			input:    "こういうの、以外に大事なんだよ",
			expected: "こういうの、以外に大事やねんて",
		},
		{
			name:     "じゃないから/ちゃうから",
			input:    "ここらへんは全国までの大会につながってるわけじゃないから、金、銀、銅の結果をもらって終わりってかんじになるかな",
			expected: "ここらへんは全国までの大会につながってるわけちゃうから、金、銀、銅の結果をもらって終わりってかんじになるかな",
		},
		{
			name:     "決めたのかと/決めたんかと",
			input:    "ごめんごめん、てっきりまた誰かに流されて高校も決めたのかと思ってたから",
			expected: "ごめんごめん、てっきりまた誰かに流されて高校も決めたんかと思ってたから",
		},
		{
			name:     "いるんだもん/おるんやもん、なれない/なれへん",
			input:    "こんなに人数いるんだもん、誰だって好きになれない子はいる",
			expected: "こんなに人数おるんやもん、誰だって好きになれへん子はいる",
		},
		{
			name:     "話してない/話してへん",
			input:    "こんな時間まで話してたのに、なんにも話してないの？",
			expected: "こんな時間まで話してたのに、なんにも話してへんの？",
		},
		{
			name:     "決めてないけど/決めてへんけど",
			input:    "さあ、決めてへんけど",
			expected: "さあ、決めてへんけど",
		},
		{
			name:     "いけないの/あかんの",
			input:    "さっきから上っ面だとかばっかり言ってるけど、それの何がいけないの",
			expected: "さっきから上っ面だとかばっかり言ってるけど、それの何があかんの",
		},
		{
			name:     "思ったの？/思ったん？",
			input:    "じゃあ、なんでうちの部に入部しようと思ったの？",
			expected: "じゃあ、なんでうちの部に入部しようと思ったん？",
		},
		{
			name:     "どうしたらいいの/どうしたらいいんさ",
			input:    "じゃあ、わたしはどうしたらいいの",
			expected: "じゃあ、うちはどうしたらいいんさ",
		},
		{
			name:     "しょうがない/しゃあない",
			input:    "しょうがない、交代で指導しよう",
			expected: "しゃあない、交代で指導しよう",
		},
		{
			name:     "伸びるんだもん/伸びんねんもん",
			input:    "すぐ伸びるんだもん",
			expected: "すぐ伸びんねんもん",
		},
		{
			name:     "にいれば/におれば",
			input:    "ずっと一緒にいればいいじゃんか",
			expected: "ずっと一緒におればええやんか",
		},
		{
			name:     "すぎない？/すぎひん？",
			input:    "せっかく音楽を好きになったのに、先輩のせいで嫌いになるって悲しすぎない？",
			expected: "せっかく音楽を好きになったのに、先輩のせいで嫌いになるって悲しすぎひん？",
		},
		{
			name:     "嫌いだった/嫌いやった",
			input:    "そういう卑怯な自分が惨めで、めっちゃ嫌いだった",
			expected: "そういう卑怯な自分が惨めで、めっちゃ嫌いやった",
		},
		{
			name:     "見れない/見られへん",
			input:    "そういう目で見れないもん",
			expected: "そういう目で見られへんもん",
		},
		{
			name:     "そうじゃなくて/そうやなくて",
			input:    "そうじゃなくて、歩きながら演奏するのがマーチングって感じ",
			expected: "そうやなくて、歩きながら演奏するのがマーチングって感じ",
		},
		{
			name:     "よく覚えて/よう覚えて",
			input:    "そうそう、よく覚えてたな",
			expected: "そうそう、よう覚えてたな",
		},
		{
			name:     "そうだった/そうやった",
			input:    "そうだったなあ、すっかり忘れてた",
			expected: "そうやったなあ、すっかり忘れてた",
		},
		{
			name:     "そうだったら/そうやったら、できない/できひん、許せない/許せへん",
			input:    "そうだったらいいんですけど、努力家なわけじゃなくて、ただできない自分が許せないだけなんですよ",
			expected: "そうやったらいいんですけど、努力家なわけじゃなくて、ただできひん自分が許せへんだけなんですよ",
		},
		{
			name:     "そうだね/そうやな",
			input:    "そうだねー",
			expected: "そうやなー",
		},
		{
			name:     "そうだよね/そうやんなあ",
			input:    "そうだよね",
			expected: "そうやんなあ",
		},
		{
			name:     "そうなんだ/そうなんや",
			input:    "そうなんだ",
			expected: "そうなんや",
		},
		{
			name:     "うれしいよ/うれしいわ、うれしいね/うれしいわ",
			input:    "そう言ってもらえるとなんかうれしいよ",
			expected: "そう言ってもらえるとなんかうれしいわ",
		},
		{
			name:     "あったんじゃないの/あったんちゃうの",
			input:    "そこでいろいろあったんじゃないの",
			expected: "そこでいろいろあったんちゃうの",
		},
		{
			name:     "嫌だなあ/嫌やなあ",
			input:    "そのせいであみかが吹奏楽嫌いになっちゃったら、ちょっと嫌だなあ",
			expected: "そのせいであみかが吹奏楽嫌いになっちゃったら、ちょっと嫌やなあ",
		},
		{
			name:     "やってやろう/やったろう",
			input:    "その心意気でやってやろうとは思ってます",
			expected: "その心意気でやったろうとは思ってます",
		},
		{
			name:     "そうだけど/そうやけど",
			input:    "そりゃあそうだけどさ",
			expected: "そりゃあそうやけどさ",
		},
		{
			name:     "聞いたの？/聞いたん？",
			input:    "それ、誰から聞いたん？",
			expected: "それ、誰から聞いたん？",
		},
		{
			name:     "ダメ/あかん",
			input:    "それがダメとはわたしは思わない",
			expected: "それがあかんとはうちは思わへん",
		},
		{
			name:     "心配なの/心配やねん",
			input:    "それがちょっと心配なの",
			expected: "それがちょっと心配やねん",
		},
		{
			name:     "やってみたらいい/やってみたらええよ",
			input:    "それがばっちりできるようになったら、楽譜のところでやってみたらいいよ",
			expected: "それがばっちりできるようになったら、楽譜のところでやってみたらええよ",
		},
		{
			name:     "だったんだよね/やってんか",
			input:    "それでさ、うちの中学ってまあまあ強かったから、オーディションとか結構シビアだったんだよね",
			expected: "それでさ、うちの中学ってまあまあ強かったから、オーディションとか結構シビアやってんか",
		},
		{
			name:     "それなのに/それやのに、なんだろう/なんやろう",
			input:    "それなのにたまに上手くいかないのはなんでなんだろう",
			expected: "それやのにたまに上手くいかないのはなんでなんやろう",
		},
		{
			name:     "それならよかった/それやったらよかったわ",
			input:    "それならよかった",
			expected: "それやったらよかったわ",
		},
		{
			name:     "当たり前じゃん/当たり前やん",
			input:    "それに、自分より上手い子に嫉妬するのは当たり前じゃん",
			expected: "それに、自分より上手い子に嫉妬するのは当たり前やん",
		},
		{
			name:     "いけないの/あかんの",
			input:    "それの何がいけないの？",
			expected: "それの何があかんの？",
		},
		{
			name:     "大丈夫なの？/大丈夫なん？",
			input:    "それは全然いいけど、大丈夫なの？",
			expected: "それは全然いいけど、大丈夫なん？",
		},
		{
			name:     "だろうけど/やろうけど、あるんじゃないか/あるんやないか",
			input:    "それもあるだろうけど、うちと同じ学年でめっちゃ上手かったトランペットの子が北宇治に行ったから、今年の北宇治はなんかあるんじゃないかって噂にはなってた",
			expected: "それもあるやろうけど、うちと同じ学年でめっちゃ上手かったトランペットの子が北宇治に行ったから、今年の北宇治はなんかあるんやないかって噂にはなってた",
		},
		{
			name:     "あったんじゃない？/あったんちゃうん？",
			input:    "そんなこと言って、ほんまは自信あったんじゃない？",
			expected: "そんなこと言うて、ほんまは自信あったんちゃうん？",
		},
		{
			name:     "手を抜いて/手ぇ抜いて、しれないよ/しれんで",
			input:    "そんなふうに手を抜いてたら、いつか後悔する日が来るかもしれないよ",
			expected: "そんなふうに手ぇ抜いてたら、いつか後悔する日が来るかもしれんで",
		},
		{
			name:     "できない/できひん",
			input:    "そんな性格だと友達できないでしょ",
			expected: "そんな性格やと友達できひんやろ",
		},
		{
			name:     "ならないでしょ/ならんやろ",
			input:    "だいたい、わたしに褒められたとか自慢にもならないでしょ",
			expected: "だいたい、うちに褒められたとか自慢にもならんやろ",
		},
		{
			name:     "ありえないよ/ありえへんよ",
			input:    "だいたい、部内恋愛はありえないよ",
			expected: "だいたい、部内恋愛はありえへんよ",
		},
		{
			name:     "いいんだよ/ええねんで",
			input:    "あみかはそのままでいいんだよ",
			expected: "あみかはそのままでええねんで",
		},
		{
			name:     "気にしないで/気にせんと",
			input:    "だから、志保は気にしないで自分の練習に集中してくれたらいいから",
			expected: "やから、志保は気にせんと自分の練習に集中してくれたらいいから",
		},
		{
			name:     "かなあ/やなあ",
			input:    "お互いウィンウィンな関係のイベントって感じかなあ",
			expected: "お互いウィンウィンな関係のイベントって感じやなあ",
		},
		{
			name:     "なるんだもん/なるんやもん",
			input:    "だって、なんか心配になるんだもん",
			expected: "だって、なんか心配になるんやもん",
		},
		{
			name:     "したじゃん/したやん",
			input:    "だって、北中の子らはだいたい南宇治に進学したじゃん",
			expected: "だって、北中の子らはだいたい南宇治に進学したやん",
		},
		{
			name:     "いない/おらん、いいじゃん/ええやん",
			input:    "だって、友達いないよりいたほうがいいじゃん",
			expected: "だって、友達おらんよりいたほうがええやん",
		},
		{
			name:     "だし/やし",
			input:    "だってあみかだし",
			expected: "だってあみかやし",
		},
		{
			name:     "じゃないし/とちゃうし",
			input:    "だってさ、人の性格なんてそんな一朝一夕に変わるようなもんじゃないし",
			expected: "だってさ、人の性格なんてそんな一朝一夕に変わるようなもんとちゃうし",
		},
		{
			name:     "だろうけど/やろうけど",
			input:    "たまたま受かってたってだけだろうけど、コンクールの舞台に出られるってのはほんとにうれしい",
			expected: "たまたま受かってたってだけやろうけど、コンクールの舞台に出られるってのはほんまにうれしい",
		},
		{
			name:     "だめだ/あかん",
			input:    "だめだ、言語中枢がやられとる",
			expected: "あかん、言語中枢がやられとる",
		},
		{
			name:     "言わなきゃだめ/言わなあかん",
			input:    "ちゃんと言わなきゃだめって思ったの",
			expected: "ちゃんと言わなあかんって思ったの",
		},
		{
			name:     "いいじゃん/ええやん",
			input:    "ちょっとぐらいはしゃいでもいいじゃんかー",
			expected: "ちょっとぐらいはしゃいでもええやんかー",
		},
		{
			name:     "待ってよ/待ってや",
			input:    "ちょっと待ってよ",
			expected: "ちょっと待ってや",
		},
		{
			name:     "ことはないよ/ことはないで",
			input:    "っていうか、マーチングする学校でもあそこまで動くことはないよ",
			expected: "っていうか、マーチングする学校でもあそこまで動くことはないで",
		},
		{
			name:     "ようよ/ようや",
			input:    "で、一緒に修学旅行の計画立てようよ",
			expected: "で、一緒に修学旅行の計画立てようや",
		},
		{
			name:     "感じだね/感じやな",
			input:    "で、二年生のときは三年生の補佐をして、三年生になったときにようやくドラムメジャーとして舞台に立つって感じだね",
			expected: "で、二年生のときは三年生の補佐をして、三年生になったときにようやくドラムメジャーとして舞台に立つって感じやな",
		},
		{
			name:     "できないこと/できないこと",
			input:    "できないことがあるなら、できるまで手伝ってあげる",
			expected: "できないことがあるなら、できるまで手伝ってあげる",
		},
		{
			name:     "しないのは/しいひんのは",
			input:    "でも、それをしないのは相手を傷つけたくないから",
			expected: "でも、それをしいひんのは相手を傷つけたくないから",
		},
		{
			name:     "思わない/思わへん",
			input:    "でも、それをちゃんと認められるなら、わたしはそいつを嫌なやつとは思わない",
			expected: "でも、それをちゃんと認められるなら、うちはそいつを嫌なやつとは思わへん",
		},
		{
			name:     "わからなくて/わからんくて",
			input:    "でも、やっぱりわからなくて",
			expected: "でも、やっぱりわからんくて",
		},
		{
			name:     "やるもの/やるもん",
			input:    "でも、吹奏楽ってみんなでやるものだから",
			expected: "でも、吹奏楽ってみんなでやるもんやから",
		},
		{
			name:     "なってきたよ/なってきたわ",
			input:    "でもまあ、久美子の顔見てたらちょっとマシになってきたよ",
			expected: "でもまあ、久美子の顔見てたらちょっとマシになってきたわ",
		},
		{
			name:     "じゃないんだ？/と違うんや？",
			input:    "でも今日はトンカツじゃないんだ？",
			expected: "でも今日はトンカツと違うんや？",
		},
		{
			name:     "どうしたの/どうしたん",
			input:    "どうしたの？",
			expected: "どうしたん？",
		},
		{
			name:     "気がする/気ぃする",
			input:    "どうにもコンクールのほうは流す感じになってる気がするなあ",
			expected: "どうにもコンクールのほうは流す感じになってる気ぃするなあ",
		},
		{
			name:     "どうもしないよ/どうもしいひんよ",
			input:    "どうもしないよ",
			expected: "どうもしいひんよ",
		},
		{
			name:     "行ってたの？/行ってたん？",
			input:    "どこ行ってたの？",
			expected: "どこ行ってたん？",
		},
		{
			name:     "呼び名だよ/呼び名やな",
			input:    "ドラムメジャーっていうのは、簡単に言うとマーチングバンドの指揮者の呼び名だよ",
			expected: "ドラムメジャーっていうのは、簡単に言うとマーチングバンドの指揮者の呼び名やな",
		},
		{
			name:     "するんだろ/するんやろ",
			input:    "どんなふうに練習するんだろ",
			expected: "どんなふうに練習するんやろ",
		},
		{
			name:     "なにもない/なんもない",
			input:    "なにもないです",
			expected: "なんもないです",
		},
		{
			name:     "なりそうだもん/なりそうやもん",
			input:    "なんか、あみかって見知らぬ人からもらった木の実とか食べて食中毒になりそうだもん",
			expected: "なんか、あみかって見知らぬ人からもらった木の実とか食べて食中毒になりそうやもん",
		},
		{
			name:     "聞いてたよ/聞いてたで",
			input:    "なんか、指導がすごい先生が来たって噂は聞いてたよ",
			expected: "なんか、指導がすごい先生が来たって噂は聞いてたで",
		},
		{
			name:     "なっちゃうんだよなあ/なっちゃうねんなあ",
			input:    "なんか危なっかしいっていうか、ぼけーっとしてるから心配になっちゃうんだよなあ",
			expected: "なんか危なっかしいっていうか、ぼけーっとしてるから心配になっちゃうねんなあ",
		},
		{
			name:     "だね/やな",
			input:    "なんていうか、うちが面倒みてやらないと！って他人に思わせるタイプだね",
			expected: "なんていうか、うちが面倒みてやらないと！って他人に思わせるタイプやな",
		},
		{
			name:     "なんなの/なんやの",
			input:    "なんなの自慢って",
			expected: "なんやの自慢って",
		},
		{
			name:     "ねえねえ/なあなあ",
			input:    "ねえねえ、アンタって呼ぶのはやめてよ",
			expected: "なあなあ、アンタって呼ぶのはやめてよ",
		},
		{
			name:     "なかったら/へんかったら",
			input:    "ノートなんて忘れなかったらよかった！",
			expected: "ノートなんて忘れへんかったらよかった！",
		},
		{
			name:     "してるよ/してるわ",
			input:    "ハイハイ、すごい期待してるよ",
			expected: "ハイハイ、めっちゃ期待してるわ",
		},
		{
			name:     "はやく/はよ、いかないと/いかんと、終わるよ/終わんで",
			input:    "はやくいかないと昼休みが終わるよ",
			expected: "はよいかんと昼休みが終わんで",
		},
		{
			name:     "たよね/たやん",
			input:    "パレードのときとか、南先輩が先頭歩いてたよね",
			expected: "パレードのときとか、南先輩が先頭歩いてたやん",
		},
		{
			name:     "決めたもの/決めたもん、ほかのもの/ほかのもん、入らなくなる/入らんくなる",
			input:    "ひとつやろうって決めたものがあると、ほかのものは全部目に入らなくなる",
			expected: "ひとつやろうって決めたもんがあると、ほかのもんは全部目に入らんくなる",
		},
		{
			name:     "ではない/ちゃう",
			input:    "べつに、いい子ではないよ",
			expected: "べつに、いい子ちゃうよ",
		},
		{
			name:     "いいんだよ/ええねんて",
			input:    "べつに、なんにも返さなくていいんだよ",
			expected: "べつに、なんにも返さなくてええねんで",
		},
		{
			name:     "気にしないけど/気にせえへんけど",
			input:    "べつに、わたしは気にしないけど",
			expected: "べつに、うちは気にせえへんけど",
		},
		{
			name:     "べつにいいじゃんか/べつにええやんか",
			input:    "べつにいいじゃんか、減るもんでもないし",
			expected: "べつにええやんか、減るもんでもないし",
		},
		{
			name:     "べつにいいって/べつにええって",
			input:    "べつにいいって",
			expected: "べつにええって",
		},
		{
			name:     "思ってない/思ってへん",
			input:    "べつに気まずいとか思ってないって",
			expected: "べつに気まずいとか思ってへんって",
		},
		{
			name:     "聞かない/聞かん、でいいよ/でいいねん",
			input:    "ほかの子の話なんて、なんも聞かないでいいよ",
			expected: "ほかの子の話なんて、なんも聞かんでいいねん",
		},
		{
			name:     "でしょ？/やろ？",
			input:    "ほら、サンフェスでも旗振ってたでしょ？",
			expected: "ほら、サンフェスでも旗振ってたやろ？",
		},
		{
			name:     "だったら/やったら、そこから/そっから、やらなきゃいけなるなる/やらなあかんくなる、じゃんか/やんか",
			input:    "ほら、トロンボーンのなかにいるときだったらわたしが面倒みれるけど、そこから出たら全部自分でやらなきゃいけなくなるじゃんか",
			expected: "ほら、トロンボーンのなかにいるときやったらうちが面倒みれるけど、そっから出たら全部自分でやらなあかんくなるやんか",
		},
		{
			name:     "なんでしょ/なんやろ",
			input:    "ほら、トロンボーンの先輩らって結構人間できてる人多いけど、花音らによれば桃花先輩ってかなり性格キツい人なんでしょ？",
			expected: "ほら、トロンボーンの先輩らって結構人間できてる人多いけど、花音らによれば桃花先輩ってかなり性格キツい人なんやろ？",
		},
		{
			name:     "だけだから/だけやから",
			input:    "ほら、ハピコンのシングってステージドリルやんの一年だけだから",
			expected: "ほら、ハピコンのシングってステージドリルやんの一年だけやから",
		},
		{
			name:     "かもしれないし/かもしれへんし",
			input:    "ほら、やっぱマンツーマンで練習したほうがいいかもしれないし",
			expected: "ほら、やっぱマンツーマンで練習したほうがいいかもしれへんし",
		},
		{
			name:     "やるじゃんか/やるやんか",
			input:    "ほら、金管楽器の場合はチューニング管を抜き差ししてやるじゃんか",
			expected: "ほら、金管楽器の場合はチューニング管を抜き差ししてやるやんか",
		},
		{
			name:     "言ってたじゃん/言うてたやん",
			input:    "ほら、志保も前言ってたじゃん",
			expected: "ほら、志保も前言うてたやん",
		},
		{
			name:     "あったじゃんか/あったやんか",
			input:    "ほら、指定の部分あったじゃんか",
			expected: "ほら、指定の部分あったやんか",
		},
	}
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, Convert(ts.input), ts.expected)
		})
	}
}
