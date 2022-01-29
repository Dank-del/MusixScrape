# <img src="https://s.mxmcdn.net/site/images/logo_icon.svg" width="45px" align="left"></img>MusixScrape

### A Golang package to scrape [musixmatch.com](https://musixmatch.com)

## How to use ?

```go
package main

import (
	"github.com/Dank-del/MusixScrape/musixScrape"
	"log"
)

func main() {
	c := musixScrape.New(nil)
	res, err := c.GetLyricsFromLink("https://www.musixmatch.com/lyrics/BAND-MAID/Sense")
	if err != nil {
		log.Fatalln(err)
	}
	if res != nil {
		log.Println(res)
		/*
		   musixmatch_test.go:32: &{Sense Band-Maid 誰しも明日へ      向かう

		        見えない 知らない
		        自由も愛も与えて
		        さぁ 選べ舞い降りた天使達

		        感じる予感に
		        Shoot the arrow of fate

		        翼に意思を 身を委ねて
		        倦まず弛まず
		        ただ 生きてたいんだ
		        もう いっそ嘘でもいいから
		        正義を貫く 最後で
		        「幸せになりたい」
		        ただそれだけ

		        無かったことに出来ても
		        望みはしないだろう
		        守るものを見つけたなら
		        覚悟は出来てるだろ

		        感じる予感に
		        Shoot the arrow of fate

		        光に向かって 手を翳せば
		        眩しいほどに
		        あぁ 今 生きてるんだ
		        とってつけた神様なんか
		        誤魔化せないんだ 未来を
		        「幸せになりたい」
		        ただそれだけ

		        見極める瞳 祈り見つめ
		        エゴでも消せはしない命を
		        Think outside the box
		        This isn′t the end

		        誰のせいじゃない
		        絶望に伏した心さえ
		        繋がるringに救われて
		        創り創られて
		        騒々しい程変わって
		        そう 叫びを矢に

		        感じる予感に
		        Shoot the arrow of fate

		        翼に意思を 身を委ねて
		        倦まず弛まず
		        ただ 生きてたいんだ
		        もう いっそ嘘でもいいから
		        正義を貫く 最後を
		        「幸せになりたい」
		        ただそれだけ

		        It ain't over till it′s over (shoot!)
		        勇往邁進 GO}
		*/
	}
	restwo, err := c.Search("Yuu Miyashita - Koufukuron")
	if err != nil {
		log.Fatalln(err)
	}
	if restwo != nil {
		log.Println(restwo)
		/*

				  ほんの一歩さえも
				  踏み外せば堕ちる奈落
			 	  怖れる心は何故また
				  生きようとする？

				  醒めない悪夢の様なこの現実は
				  神の啓示なのか悪戯なのか
				  哀れだろう だけど今更 信じたいもの見つけてしまった
				  なら覆そう降伏論を
				  絶望の全てに解を出せれば
				  最後に残る結論は希望なのか 或いは

		*/
	}
}
```


## License
[![GPLv3](https://www.gnu.org/graphics/gplv3-127x51.png)](https://www.gnu.org/licenses/gpl-3.0.en.html)
<br>Licensed Under <a href="https://www.gnu.org/licenses/gpl-3.0.en.html">GNU General Public License v3</a>
