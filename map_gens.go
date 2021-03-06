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


import "fmt"


func MakeRoomsMap(b *Board) {
	roomSizeX := MapSizeX / 5
	roomSizeY := MapSizeY / 5
	// Create room grid.
	borderIndexX := ((MapSizeX - 1) / (roomSizeX + 1)) * (roomSizeX + 1)
	borderIndexY := ((MapSizeY - 1) / (roomSizeY + 1)) * (roomSizeY + 1)
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
				(*b)[x][y].Blocked = false
				(*b)[x][y].BlocksSight = false
			} else if x % (roomSizeX + 1) == 0 || y % (roomSizeY + 1) == 0 {
				(*b)[x][y].Char = "#"
				(*b)[x][y].Name = "wall"
				(*b)[x][y].Color = "light grey"
				(*b)[x][y].ColorDark = "grey"
				(*b)[x][y].Layer = BoardLayer
				(*b)[x][y].Explored = true
				(*b)[x][y].Slows = false
				(*b)[x][y].Blocked = true
				(*b)[x][y].BlocksSight = true
			}
		}
	}
	// Create doors
	xRooms := ((MapSizeX - 1) / (roomSizeX + 1))
	yRooms := ((MapSizeY - 1) / (roomSizeY + 1))
	for i := 0; i < yRooms; i++ {
		for j := 0; j < xRooms; j++ {
			if i == 1 {
				y := i * (roomSizeY + 1) - RandRange(1, roomSizeY)
				x := j * (roomSizeX + 1)
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
			if j == 1 {
				y := i * (roomSizeY + 1)
				x := j * (roomSizeX + 1) - RandRange(1, roomSizeX)
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
			y := i * (roomSizeY + 1)
			x := j * (roomSizeX + 1) + RandRange(1, roomSizeX)
			(*b)[x][y].Char = "+"
			(*b)[x][y].Name = "closed doors"
			(*b)[x][y].Color = "darker orange"
			(*b)[x][y].ColorDark = "grey"
			(*b)[x][y].Layer = BoardLayer
			(*b)[x][y].Explored = true
			(*b)[x][y].Slows = false
			(*b)[x][y].Blocked = false
			(*b)[x][y].BlocksSight = true
			y = i * (roomSizeY + 1) + RandRange(1, roomSizeY)
			x = j * (roomSizeX + 1)
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
	}
	for x := 0; x < MapSizeX; x++ {
		for y := 0; y < MapSizeY; y++ {
			if (*b)[x][y].Char == "+" {
				if x > 0 {
					if (*b)[x-1][y].Char == "+" {
						(*b)[x][y].Char = "#"
						(*b)[x][y].Name = "wall"
						(*b)[x][y].Color = "light grey"
						(*b)[x][y].ColorDark = "grey"
						(*b)[x][y].Layer = BoardLayer
						(*b)[x][y].Explored = true
						(*b)[x][y].Slows = false
						(*b)[x][y].Blocked = true
						(*b)[x][y].BlocksSight = true
					}
				}
				if x < MapSizeX-1 {
					if (*b)[x+1][y].Char == "+" {
						(*b)[x][y].Char = "#"
						(*b)[x][y].Name = "wall"
						(*b)[x][y].Color = "light grey"
						(*b)[x][y].ColorDark = "grey"
						(*b)[x][y].Layer = BoardLayer
						(*b)[x][y].Explored = true
						(*b)[x][y].Slows = false
						(*b)[x][y].Blocked = true
						(*b)[x][y].BlocksSight = true
					}
				}
				if y > 0 {
					if (*b)[x][y-1].Char == "+" {
						(*b)[x][y].Char = "#"
						(*b)[x][y].Name = "wall"
						(*b)[x][y].Color = "light grey"
						(*b)[x][y].ColorDark = "grey"
						(*b)[x][y].Layer = BoardLayer
						(*b)[x][y].Explored = true
						(*b)[x][y].Slows = false
						(*b)[x][y].Blocked = true
						(*b)[x][y].BlocksSight = true
					}
				}
				if y < MapSizeY-1 {
					if (*b)[x][y+1].Char == "+" {
						(*b)[x][y].Char = "#"
						(*b)[x][y].Name = "wall"
						(*b)[x][y].Color = "light grey"
						(*b)[x][y].ColorDark = "grey"
						(*b)[x][y].Layer = BoardLayer
						(*b)[x][y].Explored = true
						(*b)[x][y].Slows = false
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
					(*b)[x][y].Blocked = true
					(*b)[x][y].BlocksSight = true
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
			x := j * (roomSizeY + 1) + 1
			if RandInt(100) < deleteChance {
				for xx := 0; xx < roomSizeX; xx++ {
					(*b)[x][y].Char = "."
					(*b)[x][y].Name = "floor"
					(*b)[x][y].Color = "light grey"
					(*b)[x][y].ColorDark = "grey"
					(*b)[x][y].Layer = BoardLayer
					(*b)[x][y].Explored = true
					(*b)[x][y].Slows = false
					(*b)[x][y].Blocked = false
					(*b)[x][y].BlocksSight = false
					x++
				}
			}
		}
	}
	for i := 1; i < yRooms; i++ {
		for j := 1; j < xRooms; j++ {
			y := i * (roomSizeY + 1) + 1
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
					(*b)[x][y].Blocked = false
					(*b)[x][y].BlocksSight = false
					y++
				}
			}
		}
	}
	// Remove some doors
	var doors = []*Tile{}
	for x := 0; x < MapSizeX; x++ {
		for y := 0; y < MapSizeY; y++ {
			if (*b)[x][y].Char == "+" {
				doors = append(doors, (*b)[x][y])
			}
		}
	}
	fmt.Println(doors)
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
