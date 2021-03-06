# <img src="https://s.mxmcdn.net/site/images/logo_icon.svg" width="45px" align="left"></img>MusixScrape [![Go Reference](https://pkg.go.dev/badge/github.com/Dank-del/MusixScrape.svg)](https://pkg.go.dev/github.com/Dank-del/MusixScrape) [![Go-linux](https://github.com/Dank-del/MusixScrape/actions/workflows/go-linux.yml/badge.svg)](https://github.com/Dank-del/MusixScrape/actions/workflows/go-linux.yml) [![Go-macos](https://github.com/Dank-del/MusixScrape/actions/workflows/go-macos.yml/badge.svg)](https://github.com/Dank-del/MusixScrape/actions/workflows/go-macos.yml) [![Go-windows](https://github.com/Dank-del/MusixScrape/actions/workflows/go-windows.yml/badge.svg)](https://github.com/Dank-del/MusixScrape/actions/workflows/go-windows.yml)

### A Golang package to scrape [musixmatch.com](https://www.musixmatch.com)

## How to use?

```go
package main

import (
	"log"
	"net/url"

	"github.com/Dank-del/MusixScrape/musixScrape"
)

func main() {
	c := musixScrape.New()
	url, err := url.Parse("https://www.musixmatch.com/lyrics/BAND-MAID/Sense")
	if err != nil {
		log.Fatalln(err)
	}
	lyrics, err := c.GetLyrics(url)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(lyrics)
	/*
		{Band-Maid Sense 願ってもなにも出来ないと
			わかってはいるけど
			失う怖さが邪魔をしても誰しも明日へ向かう

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
	results, err := c.Search("Yuu Miyashita - Koufukuron")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(results)
	/*
		[{宮下 遊 Koufukuron TV Size Ver. https://www.musixmatch.com/lyrics/%E5%AE%AE%E4%B8%8B-%E9%81%8A/Koufukuron-TV-Size-Ver}]
	*/
}
```

## License
[![GPLv3](https://www.gnu.org/graphics/gplv3-127x51.png)](https://www.gnu.org/licenses/gpl-3.0.en.html)
<br>Licensed Under <a href="https://www.gnu.org/licenses/gpl-3.0.en.html">GNU General Public License v3</a>
