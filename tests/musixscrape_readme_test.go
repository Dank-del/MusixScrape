/*
   MusixScrape - A Golang library to scrape lyrics from musixmatch.com
   Copyright (C) 2021  Sayan Biswas, ALiwoto

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package tests_test

import (
	"net/url"
	"testing"

	"github.com/Dank-del/MusixScrape/musixScrape"
)

func TestReadme(t *testing.T) {
	c := musixScrape.New()
	url, err := url.Parse("https://www.musixmatch.com/lyrics/BAND-MAID/Sense")
	if err != nil {
		t.Error(err)
		return
	}
	lyrics, err := c.GetLyrics(url)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(lyrics)
	/*
		musixscrape_readme_test.go:40: {Band-Maid Sense 願ってもなにも出来ないと
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
		t.Error(err)
		return
	}
	t.Log(results)
	/*
		musixscrape_readme_test.go:108: [{宮下 遊 Koufukuron TV Size Ver. https://www.musixmatch.com/lyrics/%E5%AE%AE%E4%B8%8B-%E9%81%8A/Koufukuron-TV-Size-Ver}]
	*/
}
