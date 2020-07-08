// Package mahjong is a (riichi) mahjong game client.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Tile constants
const (
	// The # of tiles in each suit
	Winds         int = 16
	Dragons       int = 12
	Man, Sou, Pin int = 36, 36, 36

	// The start and End of each suit in the slice. - 1 used for indice purposes.
	WindsStart   = 0
	WindsEnd     = WindsStart + Winds - 1
	DragonsStart = WindsEnd + 1
	DragonsEnd   = DragonsStart + Dragons - 1
	ManStart     = DragonsEnd + 1
	ManEnd       = ManStart + Man - 1
	SouStart     = ManEnd + 1
	SouEnd       = SouStart + Sou - 1
	PinStart     = SouEnd + 1
	PinEnd       = PinStart + Pin - 1

	// Define Unicode mahjong tiles
	// Winds
	East  = '🀀'
	South = '🀁'
	West  = '🀂'
	North = '🀃'
	// Dragons
	RedDragon   = '🀄'
	GreenDragon = '🀅'
	WhiteDragon = '🀆'
	// Man
	OneMan   = '🀇'
	TwoMan   = '🀈'
	ThreeMan = '🀉'
	FourMan  = '🀊'
	FiveMan  = '🀋'
	SixMan   = '🀌'
	SevenMan = '🀍'
	EightMan = '🀎'
	NineMan  = '🀏'
	// Sou
	OneSou   = '🀐'
	TwoSou   = '🀑'
	ThreeSou = '🀒'
	FourSou  = '🀓'
	FiveSou  = '🀔'
	SixSou   = '🀕'
	SevenSou = '🀖'
	EightSou = '🀗'
	NineSou  = '🀘'
	// Pin
	OnePin   = '🀙'
	TwoPin   = '🀚'
	ThreePin = '🀛'
	FourPin  = '🀜'
	FivePin  = '🀝'
	SixPin   = '🀞'
	SevenPin = '🀟'
	EightPin = '🀠'
	NinePin  = '🀡'

	// Red Fives (First five in each suit)
	MAN_RED_FIVE = 44
	SOU_RED_FIVE = 80
	PIN_RED_FIVE = 116
)

type Tile struct {
	value int
}

type Hand struct {
	tiles []int
}

type DiscardPile struct {
	ownersName string
	tiles      map[int]Tile
}

type Player struct {
	hand    []int
	discard DiscardPile
	seat    string
}

// SimpleSet compiles slices of Mahjong tiles into a basic 'one of each tile' set.
// It returns the set, which allows an unmutated set to be worked with where necessary.
func SimpleSet() []rune {
	var (
		// Build a set of unicode tiles
		WindTiles   = append((make([]rune, 0)), East, South, West, North)
		DragonTiles = append((make([]rune, 0)), RedDragon, GreenDragon, WhiteDragon)
		ManTiles    = append((make([]rune, 0)), OneMan, TwoMan, ThreeMan, FourMan, FiveMan, SixMan, SevenMan, EightMan, NineMan)
		SouTiles    = append((make([]rune, 0)), OneSou, TwoSou, ThreeSou, FourSou, FiveSou, SixSou, SevenSou, EightSou, NineSou)
		PinTiles    = append((make([]rune, 0)), OnePin, TwoPin, ThreePin, FourPin, FivePin, SixPin, SevenPin, EightPin, NinePin)
		Partial1    = append(WindTiles, DragonTiles...)
		Partial2    = append(Partial1, ManTiles...)
		Partial3    = append(Partial2, SouTiles...)
		SetTiles    = append(Partial3, PinTiles...)
	)
	return SetTiles
}

// FullSet compiles slices of Mahjong tiles into a complete 'four of each tile' set.
// It returns the set, which allows an unmutated set to be worked with where necessary.
func FullSet() []int {
	FullSet := make([]int, 136)
	for i := 0; i < len(FullSet)/4; i++ {
		for j := 0; j <= 3; j++ {
			i := i * 4
			FullSet[i+j] = i + j
		}
	}
	return FullSet
}

// GetUnicodeTile takes an integer from a full set (0-135) and returns the associated tile rune (a unicode mahjong tile).
func GetUnicodeTile(tile int) rune {
	SetTiles := SimpleSet()
	tileValue := tile / 4
	tileCode := SetTiles[tileValue]
	return tileCode
}

// ShuffleSet takes set (a slice of integers) and shuffles it.
// It returns the ShuffledSet.
func ShuffleSet(set []int) []int {
	ShuffledSet := set
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ShuffledSet), func(i, j int) { ShuffledSet[i], ShuffledSet[j] = ShuffledSet[j], ShuffledSet[i] })
	return ShuffledSet
}

// AssignSeats takes set (a slice of strings) and shuffles it.
// It returns the Seats.
func AssignSeats(set []string) []string {
	Seats := set
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(Seats), func(i, j int) { Seats[i], Seats[j] = Seats[j], Seats[i] })
	return Seats
}

func Deal() ([]int, []int, []int, []int, []int, []int, []int) {
	Tiles := ShuffleSet(FullSet())
	var DeadWall []int
	var EastHand, SouthHand, WestHand, NorthHand []int
	var DrawWall, DoraIndicators, UraDoraIndicators []int
	var EastStart, SouthStart, WestStart, NorthStart = 0, 34, 68, 102

	DiceRoll := RollDice(2)
	WallChoice := DiceRoll % 4
	WallBreak := 0

	if WallChoice == 1 {
		WallBreak = EastStart + (DiceRoll * 2)
	} else if WallChoice == 2 {
		WallBreak = SouthStart + (DiceRoll * 2)
	} else if WallChoice == 3 {
		WallBreak = WestStart + (DiceRoll * 2)
	} else if WallChoice == 4 {
		WallBreak = NorthStart + (DiceRoll * 2)
	}

	DrawWall = append(Tiles[WallBreak:], Tiles[:WallBreak]...)

	for i := 1; i <= 3; i++ {
		for w := 1; w <= 4; w++ {
			StartDraw := ((i - 1) * w * 4)
			EndDraw := StartDraw + 4
			Dealt := Tiles[StartDraw:EndDraw]
			if w == 1 {
				EastHand = append(EastHand, Dealt...)
			} else if w == 2 {
				SouthHand = append(SouthHand, Dealt...)
			} else if w == 3 {
				WestHand = append(WestHand, Dealt...)
			} else if w == 4 {
				NorthHand = append(NorthHand, Dealt...)
			}
		}
	}

	DrawWall = Tiles[47:]

	if WallBreak >= 14 {
		DeadWall = Tiles[WallBreak-14 : WallBreak]
	} else {
		remainder := 14 - WallBreak
		DeadWall = append(Tiles[:WallBreak], Tiles[(len(Tiles)-remainder):]...)
	}

	DoraIndicators = append(DoraIndicators, DeadWall[4], DeadWall[6], DeadWall[8], DeadWall[10], DeadWall[12])
	UraDoraIndicators = append(UraDoraIndicators, DeadWall[5], DeadWall[7], DeadWall[9], DeadWall[11], DeadWall[13])

	return DrawWall, DoraIndicators, UraDoraIndicators, EastHand, SouthHand, WestHand, NorthHand

	//Deal Pseudocode
	//Four Walls: 17x2 Tiles
	//Roll Dice: 1:East 2:South 3:West 4:North
	//Roll Dice: Count to dice count.
	//Begin Deal: 2x2 to East, 2x2 to South, 2x2 to West, 2x2 to North, repeat 3 times (12 tiles)
	//Dead wall: 7x2, starting from where the deal began.
	//Dora Indicator: top tile, 3 from the right of the dead wall. (ooooXoo)
}

func RollDice(num int) int {
	var dice = []int{1, 2, 3, 4, 5, 6}
	value := 0

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < num; i++ {
		value += dice[rand.Intn(len(dice))]
	}

	return value
}

func Draw(tiles []int) ([]int, []int) {
	drawn := tiles[:1]
	tiles = tiles[1:]

	return drawn, tiles
}

func Discard() {

}

func CallChi() {

}

func CallPon() {

}

func CallKan() {

}

func ClosedKan() {

}

func CallRon() {

}

func CallTsumo() {

}

func main() {
	SimpleSet()

	east := Player{seat: "East"}
	south := Player{seat: "South"}
	west := Player{seat: "West"}
	north := Player{seat: "North"}
	var drawWall, doraIndicators, uraDoraIndicators []int

	//DrawWall, DoraIndicators, UraDoraIndicators, EastHand, SouthHand, WestHand, NorthHand
	drawWall, doraIndicators, uraDoraIndicators, east.hand, south.hand, west.hand, north.hand = Deal()

	// DEEBUG / OUTPUT FOR VERIFICATION

	fmt.Println("Dealt Game")
	fmt.Printf("drawWall: ")
	for _, c := range drawWall {
		fmt.Printf("%q", GetUnicodeTile(c))
	}
	fmt.Printf("\n")
	fmt.Printf("doraIndicators: ")
	for _, c := range doraIndicators {
		fmt.Printf("%q", GetUnicodeTile(c))
	}
	fmt.Printf("\n")
	fmt.Printf("uraDoraIndicators: ")
	for _, c := range uraDoraIndicators {
		fmt.Printf("%q", GetUnicodeTile(c))
	}
	fmt.Printf("\n")
	fmt.Printf("east: ")
	for _, c := range east.hand {
		fmt.Printf("%q", GetUnicodeTile(c))
	}
	fmt.Printf("\n")
	fmt.Printf("south: ")
	for _, c := range south.hand {
		fmt.Printf("%q", GetUnicodeTile(c))
	}
	fmt.Printf("\n")
	fmt.Printf("west: ")
	for _, c := range west.hand {
		fmt.Printf("%q", GetUnicodeTile(c))
	}
	fmt.Printf("\n")
	fmt.Printf("north: ")
	for _, c := range north.hand {
		fmt.Printf("%q", GetUnicodeTile(c))
	}
	fmt.Printf("\nEnd Dealt Game\n")

	fmt.Println("SimpleSet")
	for _, c := range SimpleSet() {
		fmt.Printf("%q", c)
	}
	fmt.Println("\nEnd Simpleset")

	fmt.Println("FullSet")
	for _, c := range FullSet() {
		fmt.Printf("%q", GetUnicodeTile(c))
	}
	fmt.Println("\nEnd FullSet")

	fmt.Println("ShuffledSet")
	for _, c := range ShuffleSet(FullSet()) {
		fmt.Printf("%q", GetUnicodeTile(c))
	}
	fmt.Println("\nEnd ShuffledSet")

	fmt.Printf("Man Red Five: %q ", GetUnicodeTile(MAN_RED_FIVE))
	fmt.Printf("Pin Red Five: %q ", GetUnicodeTile(PIN_RED_FIVE))
	fmt.Printf("Sou Red Five: %q\n", GetUnicodeTile(SOU_RED_FIVE))
}
