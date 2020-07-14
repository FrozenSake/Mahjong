// Package mahjong is a (riichi) mahjong game client.
package main

//riichi

//seven pairs (chiitoitsu)

//nagashi mangan

//luck yaku
//tsumo
//ippatsu (riichi addon)
//haitei (last tile from the wall)
//houtei (last discard)
//rinshan (Dead wall draw [Off a Kan])
//chankan (Robbing a Kan)
//double riichi

//sequences
//pinfu
//iipeikou (identical sequences in one suit)
//sanshoku (three colour straight, same chi in all suits)
//ittsuu (123 456 789 in single suit)
//ryanpeikou (two sets of indentical sequences)

//triple/quad
//all pon
//three hidden pon
//sanshoku (same pon three suits)
//three kans

//terminal/honour
//tanyao (all simples)
//yakuhai (dragons, your wind, prevailing wind)
//chanta (all sets include a terminal or honour)
//junchan (pure terminals, no honours)
//honrou (only terminals/honours)
//little three dragons

//suits
//half-flush
//flush

//yakuman
//Orphans
//Four concealed pon
//Big three dragons
//little four winds
//big four winds
// all honours
// all terminals
// all green
// nine gates
// four kan

//opening yakuman
//heaven, earth, man

func CalculateYaku(hand Hand) {

}

func HasYaku(hand Hand) bool {
	yaku := hand
	if yaku.open {
		return true
	} else {
		return false
	}
}
