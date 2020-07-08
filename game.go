// Package mahjong is a (riichi) mahjong game client.
package main

import (
	"fmt"
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

func Set() []Tile {
	// Define a 136 length slice where each slice is equal to its position.
	set := make([]Tile, 136)

	for i := 0; i < len(set); i++ {
		set[i] = Tile{value: i}
	}

	fmt.Printf("Set: %v\n", set)
	return set
}

func GetUnicodeTile(tile Tile) rune {
	tileValue := tile.value / 4
	tileCode := SetTiles[tileValue]
	fmt.Printf("tileValue: %d and code: %q\n", tileValue, tileCode)
	return tileCode
}

func main() {
	Set()
	GetUnicodeTile(Tile{value: 0})
	GetUnicodeTile(Tile{value: 0})
	GetUnicodeTile(Tile{value: 120})
}
