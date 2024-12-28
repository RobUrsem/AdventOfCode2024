package puzzle

import (
	"fmt"
	"testing"
)

func TestPatterns(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		possible int
	}{
		{
			name: "1",
			input: []string{
				"r, wr, b, g, bwu, rb, gb, br",
				"",
				"brwrr",
				"bggr",
				"gbbr",
				"rrbgbr",
				"ubwu",
				"bwurrg",
				"brgr",
				"bbrgwb",
			},
			possible: 6,
		},
		{
			name: "2",
			input: []string{
				"bgbuuw, brwbg, gbrw, uuugwrw, buurwrrg, gwbgw, wgrgb, rru, ubgu, uww, wubbg, bww, gruu, wurggb, rbwb, ruubb, urbbrrgb, ugg, brr, rbrrbrbb, bur, urwg, gwwu, rbwg, bgg, uguuuwu, ruuu, wwgwub, rbrg, grrgbrb, wbwbww, buruug, rrgg, gbgwwwrb, wgwwbg, wbggb, wbgrg, b, ubbuww, grw, ru, rgb, uuggwbr, rwwwrbu, rwrwwug, wubgw, gw, gwg, buwwb, uwubub, wgw, ggw, bru, bw, rwbbur, brbw, rgrwuwb, grr, gubrw, grgw, buw, uurg, uur, bbug, wbug, buug, gbrbu, u, bgr, brbbwrg, uwwbb, bwbbwr, bwru, wugwgg, rbbrug, bgbgu, wrbb, brbbwwb, gwbwgu, bub, gwbbbr, uuw, rbrrgw, wuug, bguuug, uurgwuu, uwb, bubg, urrb, grrg, wubr, rbgurww, ugw, wbub, urw, ggbwgg, rr, ggbbub, wbrr, rwrr, wub, gu, rgrub, wgwbub, gguu, rubbrg, wr, bggwb, uwwbggr, gwu, rwbb, wbwur, wuuur, rwg, bbrgb, guww, rurbgrr, ggb, bbwgggb, rur, wwuwu, gbb, grug, grugr, rbguwwwu, ggrrr, wbr, bbg, bwrwwu, wrb, guruwu, uuru, gbw, wwrwggug, gwww, bbw, uwbrg, wuwb, bbgb, wrr, rg, rgr, buubgbrr, rbuu, wbwww, guuwr, gbbg, gurgr, wrgwuu, gurwbb, bgrbw, rbw, rbwwww, rgrb, gub, guw, wgg, gbrbbu, bb, rggg, wb, rurrbrrw, wwubg, gbwgbu, uwuurr, wgbgu, bg, buwu, ugugrw, wwrg, wgurgbu, rrggwg, ugrrg, wrbu, ugrbubgg, gbg, brruwuu, ruuw, wgb, wgubw, urb, wwb, bwrwug, brbb, gbbbwr, grwbbr, wbwgbu, ugbwb, rrgu, ruub, rrrb, wug, rwu, rubbbwu, bbwwgu, wgbrr, guwbr, ggr, rbrw, ggu, uurw, ub, uguru, uuu, bgugu, uru, gr, gur, gwb, wur, grwwugbr, gwrugu, rgwbr, bbgwu, rwbbbrg, bbuu, gg, ruw, ubgb, bubrbb, ruwwr, wwg, gbr, bu, rgg, wwr, wugbrwu, uu, brgwg, bbrbu, ubgruu, uubu, wgr, bbru, uwurr, wubur, bwgu, rbuw, gbwb, rug, ugrguwrg, wwrrg, bwu, grgub, uwu, rgguwbg, gwrbg, urbbbr, wbu, wubb, rww, wuww, gwr, rwrwwg, bbr, ugwuu, wurr, guuwurgb, urugb, bwrbu, ruwbrbu, bwurg, guu, gwbw, wrrbrrw, wuugw, gbwr, ggurg, ubu, uug, gugu, rguuwg, gug, brg, rgu, wrgggu, urruu, uwbgb, wbbbwg, rbbug, wgwg, bbrg, bbggrwb, brwr, wbgrwur, rwwugbu, uwbwgbw, buwg, ubr, guguub, rrubr, rrr, wgrwr, rgruu, rrwr, uggw, wwwrw, uw, ubggwg, bgur, gbu, rrw, rbgrgbw, guugrrr, gugw, uub, bwb, wbrrbr, ugu, rb, bbu, guuggbub, uwr, bgw, rwuruu, wbru, gbrrr, uguu, wbrw, ugbb, wuggr, wbb, ugrww, grg, bgb, rugguub, ug, urg, gwwur, ugwb, gbguu, grbgwbgu, brgggw, gwrru, urur, ruu, bbbwurw, w, grb, brb, rrb, ruur, rgw, ggwg, rruwwur, ugr, wuwwrwg, gwrr, bbgru, ubb, wbbbb, wuu, rrg, wggb, uwww, bug, rrww, ggbguru, wuw, gru, ggrwr, wrbbgru, uwrwb, ugrw, rrgwr, brwb, gww, wubw, wggr, wrww, bbwrr, uwg, wgru, ur, wwu, bugwru, gugguw, wbw, rrwrgr, rurbbu, rwb, bbb, wrgugr, buu, bugrru, rw, r, bwwr, www, gbgwbu, brggbu, ww, gwrrb, gwgw, rwwbgbr, bwr, rggwb, wwrrrr, uwwrb, rbb, bgbrwbbw, rrgr, rbr, bgbbgbb, wrbrb, rubr, guwr, urr, wrw, wurubguw, rwr, ubg, uwwrrgrb, uugwwwgu, grrbg, wru, ugb, br, bgurw, bggb, rrwruw, bwrwu, wwbr, bbww, bggr, rbg, wgu, bruu, gwbugu, gb, rrgrgbb, wrg, grbwwb, bbgwgbru, bgu, bubb, wrgru, rrub, ggwgu, wuwubub, brwrw, ubw, wbg, gugubbu, ugwwurgr, gggb, wgwrggr",
				"",
				"gbbwbruuubbrwuwburggrrgrbrrbrbbwbbwubruwbuuwuwbuwwwgrrbbwg",
				"wuubgwbggwbgwugwubgwrgwuwrbwwbbugrruwrbwuubgwgugubbuwg",
				"uwbwwrrwwwwbwgrggrurwuggwwbwubbbgwbbwwwuwguurwbrguwgbubrbr",
			},
			possible: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			onsen := ReadPatterns(tc.input)
			numPossible := onsen.FindPossible()
			if numPossible != tc.possible {
				t.Errorf("Expected %v possible but got %v", tc.possible, numPossible)
			}
		})
	}
}

func TestSingleCase(t *testing.T) {
	input := []string{
		"bgbuuw, brwbg, gbrw, uuugwrw, buurwrrg, gwbgw, wgrgb, rru, ubgu, uww, wubbg, bww, gruu, wurggb, rbwb, ruubb, urbbrrgb, ugg, brr, rbrrbrbb, bur, urwg, gwwu, rbwg, bgg, uguuuwu, ruuu, wwgwub, rbrg, grrgbrb, wbwbww, buruug, rrgg, gbgwwwrb, wgwwbg, wbggb, wbgrg, b, ubbuww, grw, ru, rgb, uuggwbr, rwwwrbu, rwrwwug, wubgw, gw, gwg, buwwb, uwubub, wgw, ggw, bru, bw, rwbbur, brbw, rgrwuwb, grr, gubrw, grgw, buw, uurg, uur, bbug, wbug, buug, gbrbu, u, bgr, brbbwrg, uwwbb, bwbbwr, bwru, wugwgg, rbbrug, bgbgu, wrbb, brbbwwb, gwbwgu, bub, gwbbbr, uuw, rbrrgw, wuug, bguuug, uurgwuu, uwb, bubg, urrb, grrg, wubr, rbgurww, ugw, wbub, urw, ggbwgg, rr, ggbbub, wbrr, rwrr, wub, gu, rgrub, wgwbub, gguu, rubbrg, wr, bggwb, uwwbggr, gwu, rwbb, wbwur, wuuur, rwg, bbrgb, guww, rurbgrr, ggb, bbwgggb, rur, wwuwu, gbb, grug, grugr, rbguwwwu, ggrrr, wbr, bbg, bwrwwu, wrb, guruwu, uuru, gbw, wwrwggug, gwww, bbw, uwbrg, wuwb, bbgb, wrr, rg, rgr, buubgbrr, rbuu, wbwww, guuwr, gbbg, gurgr, wrgwuu, gurwbb, bgrbw, rbw, rbwwww, rgrb, gub, guw, wgg, gbrbbu, bb, rggg, wb, rurrbrrw, wwubg, gbwgbu, uwuurr, wgbgu, bg, buwu, ugugrw, wwrg, wgurgbu, rrggwg, ugrrg, wrbu, ugrbubgg, gbg, brruwuu, ruuw, wgb, wgubw, urb, wwb, bwrwug, brbb, gbbbwr, grwbbr, wbwgbu, ugbwb, rrgu, ruub, rrrb, wug, rwu, rubbbwu, bbwwgu, wgbrr, guwbr, ggr, rbrw, ggu, uurw, ub, uguru, uuu, bgugu, uru, gr, gur, gwb, wur, grwwugbr, gwrugu, rgwbr, bbgwu, rwbbbrg, bbuu, gg, ruw, ubgb, bubrbb, ruwwr, wwg, gbr, bu, rgg, wwr, wugbrwu, uu, brgwg, bbrbu, ubgruu, uubu, wgr, bbru, uwurr, wubur, bwgu, rbuw, gbwb, rug, ugrguwrg, wwrrg, bwu, grgub, uwu, rgguwbg, gwrbg, urbbbr, wbu, wubb, rww, wuww, gwr, rwrwwg, bbr, ugwuu, wurr, guuwurgb, urugb, bwrbu, ruwbrbu, bwurg, guu, gwbw, wrrbrrw, wuugw, gbwr, ggurg, ubu, uug, gugu, rguuwg, gug, brg, rgu, wrgggu, urruu, uwbgb, wbbbwg, rbbug, wgwg, bbrg, bbggrwb, brwr, wbgrwur, rwwugbu, uwbwgbw, buwg, ubr, guguub, rrubr, rrr, wgrwr, rgruu, rrwr, uggw, wwwrw, uw, ubggwg, bgur, gbu, rrw, rbgrgbw, guugrrr, gugw, uub, bwb, wbrrbr, ugu, rb, bbu, guuggbub, uwr, bgw, rwuruu, wbru, gbrrr, uguu, wbrw, ugbb, wuggr, wbb, ugrww, grg, bgb, rugguub, ug, urg, gwwur, ugwb, gbguu, grbgwbgu, brgggw, gwrru, urur, ruu, bbbwurw, w, grb, brb, rrb, ruur, rgw, ggwg, rruwwur, ugr, wuwwrwg, gwrr, bbgru, ubb, wbbbb, wuu, rrg, wggb, uwww, bug, rrww, ggbguru, wuw, gru, ggrwr, wrbbgru, uwrwb, ugrw, rrgwr, brwb, gww, wubw, wggr, wrww, bbwrr, uwg, wgru, ur, wwu, bugwru, gugguw, wbw, rrwrgr, rurbbu, rwb, bbb, wrgugr, buu, bugrru, rw, r, bwwr, www, gbgwbu, brggbu, ww, gwrrb, gwgw, rwwbgbr, bwr, rggwb, wwrrrr, uwwrb, rbb, bgbrwbbw, rrgr, rbr, bgbbgbb, wrbrb, rubr, guwr, urr, wrw, wurubguw, rwr, ubg, uwwrrgrb, uugwwwgu, grrbg, wru, ugb, br, bgurw, bggb, rrwruw, bwrwu, wwbr, bbww, bggr, rbg, wgu, bruu, gwbugu, gb, rrgrgbb, wrg, grbwwb, bbgwgbru, bgu, bubb, wrgru, rrub, ggwgu, wuwubub, brwrw, ubw, wbg, gugubbu, ugwwurgr, gggb, wgwrggr",
		"",
		"ugbrrggwbwwububbbuwbugrruuugugwbwwrgbrwbrwgg",
		"brurgwgbuwuwwgbuwgruuwgbgwwbubbbrbrbwrwugwuuggwg",
		"wbgwbuwbwrggggrbwguwggwgrwwgrbwbwgwwwwgbggurggrggbwbbwbgwg",
		"gwggbbbwuwbwurugubruubwburgwrbrwwbgruurbbgrbrgbwgugg",
		"wubbbrgbwuwburrrgrubbgwrrugrgrgrgwbwrggggwuurbwggrgwggg",
	}
	onsen := ReadPatterns(input)
	numPossible := onsen.FindPossible()
	if numPossible == 5 {
		fmt.Println("Success")
	} else {
		fmt.Println("Failed")
	}
}
