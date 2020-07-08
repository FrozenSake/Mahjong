// Package mahjong is a (riichi) mahjong game client.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
)

type Tile struct {
	value int
}

type Hand struct {
	tiles map[int]Tile
}

type Discard struct {
	tiles map[int]Tile
}

type Player struct {
	hand    Hand
	discard Discard
}

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

func GetUnicodeTile(tile int) rune {
	SetTiles := SimpleSet()
	tileValue := tile / 4
	tileCode := SetTiles[tileValue]
	return tileCode
}

func ShuffleSet(set []int) []int {
	ShuffledSet := set
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ShuffledSet), func(i, j int) { ShuffledSet[i], ShuffledSet[j] = ShuffledSet[j], ShuffledSet[i] })
	return ShuffledSet
}

func main() {
	SimpleSet()
	fmt.Println("SimpleSet")
	for _, c := range SimpleSet() {
		fmt.Printf("%q ", c)
	}
	fmt.Println("\nEnd Simpleset")
	fmt.Println("FullSet")
	for _, c := range FullSet() {
		fmt.Printf("%q ", GetUnicodeTile(c))
	}
	fmt.Println("\nEnd FullSet")
	fmt.Println("ShuffledSet")
	for _, c := range ShuffleSet(FullSet()) {
		fmt.Printf("%q ", GetUnicodeTile(c))
	}
	fmt.Println("\nEnd ShuffledSet")
}
