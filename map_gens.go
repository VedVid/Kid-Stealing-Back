/*
Copyright (c) 2018, Tomasz "VedVid" Nowakowski
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package main


import "math/rand"


const (
	MonstersMin = 4
	MonstersMax = 6
)

var (
	TreasureMin        = 6
	TreasureMax        = 12
	TreasureLightMin   = 3
	TreasureLightMax   = 6
	TreasureMediumMin  = 2
	TreasureMediumMax  = 4
	TreasureHeavyMin   = 1
	TreasureHeavyMax   = 2
	TreasureCharLight  = "☼"
	TreasureCharMedium = "☥"
	TreasureCharHeavy  = "⚱"
)

func MakeRoomsMap(b *Board) {
	roomSizeX := MapSizeX / 5
	roomSizeY := MapSizeY / 5
	// Create room grid.
	borderIndexX := ((MapSizeX - 1) / (roomSizeX + 1)) * (roomSizeX + 1)
	borderIndexY := ((MapSizeY - 1) / (roomSizeY + 1)) * (roomSizeY + 1)
	crumblinWallChance := 12
	for x := 0; x < MapSizeX; x++ {
		for y := 0; y < MapSizeY; y++ {
			if x > borderIndexX || y > borderIndexY {
				(*b)[x][y].Char = "."
				(*b)[x][y].Name = "floor"
				(*b)[x][y].Color = "light grey"
				(*b)[x][y].ColorDark = "grey"
				(*b)[x][y].Layer = BoardLayer
				(*b)[x][y].Explored = true
				(*b)[x][y].Slows = false
				(*b)[x][y].Hides = false
				(*b)[x][y].Blocked = false
				(*b)[x][y].BlocksSight = false
			} else if x%(roomSizeX+1) == 0 || y%(roomSizeY+1) == 0 {
				if crumblinWallChance > RandInt(100) && x > 0 && x < MapSizeX-1 && y > 0 && y < MapSizeY-1 {
					(*b)[x][y].Char = "%"
					(*b)[x][y].Name = "crumbling wall"
					(*b)[x][y].Color = "#404b59"
					(*b)[x][y].ColorDark = "grey"
					(*b)[x][y].Layer = BoardLayer
					(*b)[x][y].Explored = true
					(*b)[x][y].Slows = false
					(*b)[x][y].Hides = true
					(*b)[x][y].Blocked = true
					(*b)[x][y].BlocksSight = true
				} else {
					(*b)[x][y].Char = "#"
					(*b)[x][y].Name = "wall"
					(*b)[x][y].Color = "light grey"
					(*b)[x][y].ColorDark = "grey"
					(*b)[x][y].Layer = BoardLayer
					(*b)[x][y].Explored = true
					(*b)[x][y].Slows = false
					(*b)[x][y].Hides = false
					(*b)[x][y].Blocked = true
					(*b)[x][y].BlocksSight = true
				}
			}
		}
	}
	// Add some decorations
	var rooms = [][]int{ // StartX, StartY, EndX, EndY
		// first row
		[]int{1, 1, 5, 5},
		[]int{7, 1, 11, 5},
		[]int{13, 1, 17, 5},
		[]int{19, 1, 23, 5},
		// second row
		[]int{1, 7, 5, 11},
		[]int{7, 7, 11, 11},
		[]int{13, 7, 17, 11},
		[]int{19, 7, 23, 11},
		// third row
		[]int{1, 13, 5, 17},
		[]int{7, 13, 11, 17},
		[]int{13, 13, 17, 17},
		[]int{19, 13, 23, 17},
		// fourht row
		[]int{1, 19, 5, 23},
		[]int{7, 19, 11, 23},
		[]int{13, 19, 17, 23},
		[]int{19, 19, 23, 23},
	}
	decorationChance := 30
	for row := 0; row < len(rooms); row++ {
		room := rooms[row]
		if decorationChance < RandInt(100) {
			layoutRoom := HardcodedRooms[rand.Intn(len(HardcodedRooms))]
			for xx := 0; xx < 5; xx++ {
				for yy := 0; yy < 5; yy++ {
					ch := layoutRoom[xx][yy]
					cx := room[0] + xx
					cy := room[1] + yy
					switch ch {
					case ".":
						(*b)[cx][cy].Char = "."
						(*b)[cx][cy].Name = "floor"
						(*b)[cx][cy].Color = "light grey"
						(*b)[cx][cy].ColorDark = "grey"
						(*b)[cx][cy].Layer = BoardLayer
						(*b)[cx][cy].Explored = true
						(*b)[cx][cy].Slows = false
						(*b)[cx][cy].Hides = false
						(*b)[cx][cy].Blocked = false
						(*b)[cx][cy].BlocksSight = false
					case "o":
						(*b)[cx][cy].Char = "○"
						(*b)[cx][cy].Name = "pillar"
						(*b)[cx][cy].Color = "grey"
						(*b)[cx][cy].ColorDark = "grey"
						(*b)[cx][cy].Layer = BoardLayer
						(*b)[cx][cy].Explored = true
						(*b)[cx][cy].Slows = false
						(*b)[cx][cy].Hides = false
						(*b)[cx][cy].Blocked = true
						(*b)[cx][cy].BlocksSight = true
					case "h":
						(*b)[cx][cy].Char = "h"
						(*b)[cx][cy].Name = "chair"
						(*b)[cx][cy].Color = "dark orange"
						(*b)[cx][cy].ColorDark = "grey"
						(*b)[cx][cy].Layer = BoardLayer
						(*b)[cx][cy].Explored = true
						(*b)[cx][cy].Slows = false
						(*b)[cx][cy].Hides = false
						(*b)[cx][cy].Blocked = false
						(*b)[cx][cy].BlocksSight = false
					case "T":
						(*b)[cx][cy].Char = "T"
						(*b)[cx][cy].Name = "table"
						(*b)[cx][cy].Color = "dark orange"
						(*b)[cx][cy].ColorDark = "grey"
						(*b)[cx][cy].Layer = BoardLayer
						(*b)[cx][cy].Explored = true
						(*b)[cx][cy].Slows = false
						(*b)[cx][cy].Hides = true
						(*b)[cx][cy].Blocked = false
						(*b)[cx][cy].BlocksSight = false
					}
				}
			}
		}
	}
	// Create doors
	xRooms := ((MapSizeX - 1) / (roomSizeX + 1))
	yRooms := ((MapSizeY - 1) / (roomSizeY + 1))
	for i := 0; i < yRooms; i++ {
		for j := 0; j < xRooms; j++ {
			if i == 1 {
				y := i*(roomSizeY+1) - RandRange(1, roomSizeY)
				x := j * (roomSizeX + 1)
				(*b)[x][y].Char = "+"
				(*b)[x][y].Name = "closed doors"
				(*b)[x][y].Color = "darker orange"
				(*b)[x][y].ColorDark = "grey"
				(*b)[x][y].Layer = BoardLayer
				(*b)[x][y].Explored = true
				(*b)[x][y].Slows = false
				(*b)[x][y].Hides = false
				(*b)[x][y].Blocked = false
				(*b)[x][y].BlocksSight = true
			}
			if j == 1 {
				y := i * (roomSizeY + 1)
				x := j*(roomSizeX+1) - RandRange(1, roomSizeX)
				(*b)[x][y].Char = "+"
				(*b)[x][y].Name = "closed doors"
				(*b)[x][y].Color = "darker orange"
				(*b)[x][y].ColorDark = "grey"
				(*b)[x][y].Layer = BoardLayer
				(*b)[x][y].Explored = true
				(*b)[x][y].Slows = false
				(*b)[x][y].Hides = false
				(*b)[x][y].Blocked = false
				(*b)[x][y].BlocksSight = true
			}
			y := i * (roomSizeY + 1)
			x := j*(roomSizeX+1) + RandRange(1, roomSizeX)
			(*b)[x][y].Char = "+"
			(*b)[x][y].Name = "closed doors"
			(*b)[x][y].Color = "darker orange"
			(*b)[x][y].ColorDark = "grey"
			(*b)[x][y].Layer = BoardLayer
			(*b)[x][y].Explored = true
			(*b)[x][y].Slows = false
			(*b)[x][y].Hides = false
			(*b)[x][y].Blocked = false
			(*b)[x][y].BlocksSight = true
			y = i*(roomSizeY+1) + RandRange(1, roomSizeY)
			x = j * (roomSizeX + 1)
			(*b)[x][y].Char = "+"
			(*b)[x][y].Name = "closed doors"
			(*b)[x][y].Color = "darker orange"
			(*b)[x][y].ColorDark = "grey"
			(*b)[x][y].Layer = BoardLayer
			(*b)[x][y].Explored = true
			(*b)[x][y].Slows = false
			(*b)[x][y].Hides = false
			(*b)[x][y].Blocked = false
			(*b)[x][y].BlocksSight = true
		}
	}
	for x := 0; x < MapSizeX; x++ {
		for y := 0; y < MapSizeY; y++ {
			if (*b)[x][y].Char == "+" {
				if x > 0 {
					if (*b)[x-1][y].Char != "." && (*b)[x-1][y].Char != "#" {
						(*b)[x][y].Char = "#"
						(*b)[x][y].Name = "wall"
						(*b)[x][y].Color = "light grey"
						(*b)[x][y].ColorDark = "grey"
						(*b)[x][y].Layer = BoardLayer
						(*b)[x][y].Explored = true
						(*b)[x][y].Slows = false
						(*b)[x][y].Hides = false
						(*b)[x][y].Blocked = true
						(*b)[x][y].BlocksSight = true
					}
				}
				if x < MapSizeX-1 {
					if (*b)[x+1][y].Char != "." && (*b)[x+1][y].Char != "#" {
						(*b)[x][y].Char = "#"
						(*b)[x][y].Name = "wall"
						(*b)[x][y].Color = "light grey"
						(*b)[x][y].ColorDark = "grey"
						(*b)[x][y].Layer = BoardLayer
						(*b)[x][y].Explored = true
						(*b)[x][y].Slows = false
						(*b)[x][y].Hides = false
						(*b)[x][y].Blocked = true
						(*b)[x][y].BlocksSight = true
					}
				}
				if y > 0 {
					if (*b)[x][y-1].Char != "." && (*b)[x][y-1].Char != "#" {
						(*b)[x][y].Char = "#"
						(*b)[x][y].Name = "wall"
						(*b)[x][y].Color = "light grey"
						(*b)[x][y].ColorDark = "grey"
						(*b)[x][y].Layer = BoardLayer
						(*b)[x][y].Explored = true
						(*b)[x][y].Slows = false
						(*b)[x][y].Hides = false
						(*b)[x][y].Blocked = true
						(*b)[x][y].BlocksSight = true
					}
				}
				if y < MapSizeY-1 {
					if (*b)[x][y+1].Char != "." && (*b)[x][y+1].Char != "#" {
						(*b)[x][y].Char = "#"
						(*b)[x][y].Name = "wall"
						(*b)[x][y].Color = "light grey"
						(*b)[x][y].ColorDark = "grey"
						(*b)[x][y].Layer = BoardLayer
						(*b)[x][y].Explored = true
						(*b)[x][y].Slows = false
						(*b)[x][y].Hides = false
						(*b)[x][y].Blocked = true
						(*b)[x][y].BlocksSight = true
					}
				}
				if x == 0 || x >= MapSizeX-1 || y == 0 || y >= MapSizeY-1 {
					(*b)[x][y].Char = "#"
					(*b)[x][y].Name = "wall"
					(*b)[x][y].Color = "light grey"
					(*b)[x][y].ColorDark = "grey"
					(*b)[x][y].Layer = BoardLayer
					(*b)[x][y].Explored = true
					(*b)[x][y].Slows = false
					(*b)[x][y].Hides = false
					(*b)[x][y].Blocked = true
					(*b)[x][y].BlocksSight = true
				}
			} else if (*b)[x][y].Char == "%" {
				if x > 0 {
					if (*b)[x-1][y].Char == "%" {
						(*b)[x][y].Char = "#"
						(*b)[x][y].Name = "wall"
						(*b)[x][y].Color = "light grey"
						(*b)[x][y].ColorDark = "grey"
						(*b)[x][y].Layer = BoardLayer
						(*b)[x][y].Explored = true
						(*b)[x][y].Slows = false
						(*b)[x][y].Hides = false
						(*b)[x][y].Blocked = true
						(*b)[x][y].BlocksSight = true
					}
				}
				if x < MapSizeX-1 {
					if (*b)[x+1][y].Char == "%" {
						(*b)[x][y].Char = "#"
						(*b)[x][y].Name = "wall"
						(*b)[x][y].Color = "light grey"
						(*b)[x][y].ColorDark = "grey"
						(*b)[x][y].Layer = BoardLayer
						(*b)[x][y].Explored = true
						(*b)[x][y].Slows = false
						(*b)[x][y].Hides = false
						(*b)[x][y].Blocked = true
						(*b)[x][y].BlocksSight = true
					}
				}
				if y > 0 {
					if (*b)[x][y-1].Char == "%" {
						(*b)[x][y].Char = "#"
						(*b)[x][y].Name = "wall"
						(*b)[x][y].Color = "light grey"
						(*b)[x][y].ColorDark = "grey"
						(*b)[x][y].Layer = BoardLayer
						(*b)[x][y].Explored = true
						(*b)[x][y].Slows = false
						(*b)[x][y].Hides = false
						(*b)[x][y].Blocked = true
						(*b)[x][y].BlocksSight = true
					}
				}
				if y < MapSizeY-1 {
					if (*b)[x][y+1].Char == "%" {
						(*b)[x][y].Char = "#"
						(*b)[x][y].Name = "wall"
						(*b)[x][y].Color = "light grey"
						(*b)[x][y].ColorDark = "grey"
						(*b)[x][y].Layer = BoardLayer
						(*b)[x][y].Explored = true
						(*b)[x][y].Slows = false
						(*b)[x][y].Hides = false
						(*b)[x][y].Blocked = true
						(*b)[x][y].BlocksSight = true
					}
				}
			}
		}
	}
	// Crush walls
	xRooms = ((MapSizeX - 1) / (roomSizeX + 1))
	yRooms = ((MapSizeY - 1) / (roomSizeY + 1))
	deleteChance := 33
	for i := 1; i < yRooms; i++ {
		for j := 1; j < xRooms; j++ {
			y := i * (roomSizeX + 1)
			x := j*(roomSizeY+1) + 1
			if RandInt(100) < deleteChance {
				for xx := 0; xx < roomSizeX; xx++ {
					(*b)[x][y].Char = "."
					(*b)[x][y].Name = "floor"
					(*b)[x][y].Color = "light grey"
					(*b)[x][y].ColorDark = "grey"
					(*b)[x][y].Layer = BoardLayer
					(*b)[x][y].Explored = true
					(*b)[x][y].Slows = false
					(*b)[x][y].Hides = false
					(*b)[x][y].Blocked = false
					(*b)[x][y].BlocksSight = false
					x++
				}
			}
		}
	}
	for i := 1; i < yRooms; i++ {
		for j := 1; j < xRooms; j++ {
			y := i*(roomSizeY+1) + 1
			x := j * (roomSizeX + 1)
			if RandInt(100) < deleteChance {
				for yy := 0; yy < roomSizeY; yy++ {
					(*b)[x][y].Char = "."
					(*b)[x][y].Name = "floor"
					(*b)[x][y].Color = "light grey"
					(*b)[x][y].ColorDark = "grey"
					(*b)[x][y].Layer = BoardLayer
					(*b)[x][y].Explored = true
					(*b)[x][y].Slows = false
					(*b)[x][y].Hides = false
					(*b)[x][y].Blocked = false
					(*b)[x][y].BlocksSight = false
					y++
				}
			}
		}
	}
	// Put some treasures
	treasureL := RandRange(TreasureLightMin, TreasureLightMax)
	treasureM := RandRange(TreasureMediumMin, TreasureMediumMax)
	treasureH := RandRange(TreasureHeavyMin, TreasureHeavyMax)
	treasureLC := 0
	treasureMC := 0
	treasureHC := 0
	for {
		if treasureLC >= treasureL && treasureMC >= treasureM &&
			treasureHC >= treasureH {
			break
		}
		x := RandInt(MapSizeX - 1)
		y := RandInt(MapSizeY - 1)
		if (*b)[x][y].Char == "." {
			chances := RandInt(100)
			if 33 < chances && treasureLC < treasureL {
				(*b)[x][y].Treasure = true
				(*b)[x][y].TreasureCol = "yellow"
				(*b)[x][y].TreasureChar = TreasureCharLight
				treasureLC++
			} else if 66 < chances && treasureMC < treasureM {
				(*b)[x][y].Treasure = true
				(*b)[x][y].TreasureCol = "yellow"
				(*b)[x][y].TreasureChar = TreasureCharMedium
				treasureMC++
			} else if treasureHC < treasureH {
				(*b)[x][y].Treasure = true
				(*b)[x][y].TreasureCol = "yellow"
				(*b)[x][y].TreasureChar = TreasureCharHeavy
				treasureHC++
			}
		}
	}
	// Remove some doors
	wallChance := 40
	floorChance := 55
	for x := 0; x < MapSizeX; x++ {
		for y := 0; y < MapSizeY; y++ {
			if (*b)[x][y].Char == "+" {
				chance := RandInt(100)
				if chance < wallChance {
					(*b)[x][y].Char = "#"
					(*b)[x][y].Name = "wall"
					(*b)[x][y].Color = "light grey"
					(*b)[x][y].ColorDark = "grey"
					(*b)[x][y].Layer = BoardLayer
					(*b)[x][y].Explored = true
					(*b)[x][y].Slows = false
					(*b)[x][y].Blocked = true
					(*b)[x][y].BlocksSight = true
					if TestMapTilesConnections(b) == false {
						(*b)[x][y].Char = "+"
						(*b)[x][y].Name = "closed doors"
						(*b)[x][y].Color = "darker orange"
						(*b)[x][y].ColorDark = "grey"
						(*b)[x][y].Layer = BoardLayer
						(*b)[x][y].Explored = true
						(*b)[x][y].Slows = false
						(*b)[x][y].Blocked = false
						(*b)[x][y].BlocksSight = true
					}
				} else if chance < floorChance {
					(*b)[x][y].Char = "."
					(*b)[x][y].Name = "floor"
					(*b)[x][y].Color = "light grey"
					(*b)[x][y].ColorDark = "grey"
					(*b)[x][y].Layer = BoardLayer
					(*b)[x][y].Explored = true
					(*b)[x][y].Slows = false
					(*b)[x][y].Blocked = false
					(*b)[x][y].BlocksSight = false
				}
			}
		}
	}
}

func MakeRandomMap(b *Board) {
	for x := 0; x < MapSizeX; x++ {
		for y := 0; y < MapSizeY; y++ {
			ch := RandInt(100)
			if ch <= 20 {
				(*b)[x][y].Char = "#"
				(*b)[x][y].Name = "wall"
				(*b)[x][y].Color = "lightest grey"
				(*b)[x][y].ColorDark = "grey"
				(*b)[x][y].Layer = BoardLayer
				(*b)[x][y].Explored = true
				(*b)[x][y].Slows = false
				(*b)[x][y].Blocked = true
				(*b)[x][y].BlocksSight = true
			} else {
				(*b)[x][y].Char = ","
				(*b)[x][y].Name = "grass"
				(*b)[x][y].Color = "#D2B48C"
				(*b)[x][y].ColorDark = "grey"
				(*b)[x][y].Layer = BoardLayer
				(*b)[x][y].Explored = true
				(*b)[x][y].Slows = false
				(*b)[x][y].Blocked = false
				(*b)[x][y].BlocksSight = false
			}
		}
	}
}
