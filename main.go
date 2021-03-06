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

import (
	blt "bearlibterminal"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var KeyboardLayout int
var CustomControls bool

var MsgBuf = []string{}
var Game = GameData{[]string{}, 0, []string{}, 0, 0, 0, 0, nil, 20, 1, 1, 10, 0, 0, 0}
var HighScores = Scores{[]Score{}}

var GlobalSeed int64

func main() {
	var cells = new(Board)
	var objs = new(Objects)
	var actors = new(Creatures)
	blt.Composition(blt.TK_ON)
	StartGame(cells, actors, objs)
	for {
		//		if (Game.TurnCounter == 0 || Game.TurnCounter%Game.SpawnRatio == 0) &&
		//			Game.BreakTime <= 0 {
		//			SpawnMonsters(*cells, actors)
		//		}
		if Game.WaveCur >= Game.WaveMax && Game.LivingMonsters <= 0 {
			AddMessage("You fought well. Now, there is some time to rest.")
			Game.BreakTime += 10
			Game.WaveNo++
			Game.WaveMax = Game.WaveNo * 10
			Game.WaveCur = 0
			if Game.WaveNo%3 == 0 {
				Game.SpawnAmount++
			}
			if Game.WaveNo%2 == 0 {
				Game.WaveMax += 4
			}
			(*actors)[0].Restore()
		}
		Game.Points = Game.CalculatePoints()
		RenderAll(*cells, *objs, *actors)
		if (*actors)[0].HPCurrent <= 0 {
			UpdateScores(*actors)
			SaveScores(HighScores)
			DeleteSaves()
			blt.Read()
			break
		}
		key := ReadInput()
		if (key == blt.TK_S && blt.Check(blt.TK_SHIFT) != 0) ||
			key == blt.TK_CLOSE {
			// It saves GameData, too
			err := SaveGame(*cells, *actors, *objs)
			if err != nil {
				fmt.Println(err)
			}
			break
		} else if key == blt.TK_Q && blt.Check(blt.TK_SHIFT) != 0 {
			AddMessage("Do you really want to quit? It'll delete your saves. Y/N")
			end := false
			for {
				decision := ReadInput()
				if decision == blt.TK_Y {
					DeleteSaves()
					end = true
					break
				} else if decision == blt.TK_N {
					AddMessage("Okay, then.")
					break
				}
			}
			if end == true {
				break
			}
		} else {
			turnSpent := Controls(key, (*actors)[0], cells, actors, objs)
			if turnSpent == true {
				Game.TurnCounter++
				if Game.BreakTime > 0 {
					Game.BreakTime--
				}
				CreaturesTakeTurn(*cells, actors, *objs)
			}
		}
	}
	blt.Close()
}

func NewGame(b *Board, c *Creatures, o *Objects) {
	/* Function NewGame initializes game state - creates player, monsters, and game map.
	   This implementation is generic-placeholder, for testing purposes. */
	// ATTENTION!
	// The commented code below is perfectly valid example
	// of creating game data during runtime.
	/*
		enemy, err := NewCreature(MapSizeX-2, MapSizeY-2, "patherRanged.json")
		if err != nil {
			fmt.Println(err)
		}
		w1, err := NewObject(0, 0, "weapon1.json")
		if err != nil {
			fmt.Println(err)
		}
		w2, err := NewObject(0, 0, "weapon2.json")
		if err != nil {
			fmt.Println(err)
		}
		wm, err := NewObject(0, 0, "melee.json")
		if err != nil {
			fmt.Println(err)
		}
		var enemyEq = EquipmentComponent{Objects{w1, w2, wm}, Objects{}}
		enemy.EquipmentComponent = enemyEq
		*c = Creatures{player, enemy}
		obj, err := NewObject(24, 15, "heal.json")
		*o = Objects{obj}
		if err != nil {
			fmt.Println(err)
		}
	*/
	var err error
	*b = InitializeEmptyMap()
	tries := 1000
	for {
		MakeRoomsMap(b)
		validPass := TestMapTilesConnections(b)
		validTreasures := TestMapTreasureDistribution(b)
		if validPass && validTreasures {
			break
		}
		ZeroMap(b)
		tries--
		if tries < 0 {
			break
		}
	}
	playerX, playerY := RandInt(len(*b)-1), RandInt(len((*b)[0])-1)
	for {
		if (*b)[playerX][playerY].Blocked == false &&
			(*b)[playerX][playerY].Treasure == false &&
			(*b)[playerX][playerY].Hides == false {
			break
		}
		playerX, playerY = RandInt(len(*b)-1), RandInt(len((*b)[0])-1)
	}
	player, err := NewPlayer(playerX, playerY)
	if err != nil {
		fmt.Println(err)
	}
	(*b)[playerX][playerY].Char = ">"
	(*b)[playerX][playerY].Name = "Hatch"
	(*b)[playerX][playerY].Color = "lightest gray"
	(*b)[playerX][playerY].ColorDark = "grey"
	*c = Creatures{player}
	Game.MonstersList = ReadMonsterFiles()
}

func StartGame(b *Board, c *Creatures, o *Objects) {
	/* Function StartGame determines if game save is present (and valid), then
	   loads data, or initializes new game.
	   Panics if some-but-not-all save files are missing. */
	_, errHighScores := os.Stat(ScoresPathGob)
	_, errBoard := os.Stat(MapPathGob)
	_, errCreatures := os.Stat(CreaturesPathGob)
	_, errObjects := os.Stat(ObjectsPathGob)
	_, errGame := os.Stat(GamePathGob)
	if errHighScores == nil {
		_ = LoadScores(&HighScores)
	} else {
		_ = SaveScores(HighScores)
	}
	if errBoard == nil && errCreatures == nil && errObjects == nil && errGame == nil {
		LoadGame(b, c, o)
	} else if errBoard != nil && errCreatures != nil && errObjects != nil && errGame != nil {
		NewGame(b, c, o)
	} else {
		txt := CorruptedSaveError(errBoard, errCreatures, errObjects, errGame)
		fmt.Println("Error: save files are corrupted: " + txt)
		panic(-1)
	}
}

func (g GameData) CalculatePoints() int {
	points := 0
	points += g.TurnCounter
	points -= g.TotalHPLost * 5
	points += g.TotalDMGDealt * 5
	points += (g.WaveNo - 1) * 100
	points += g.KillPoints
	return points
}

func init() {
	GlobalSeed = time.Now().UTC().UnixNano()
	rand.Seed(GlobalSeed)
	fmt.Println(GlobalSeed)
	InitializeFOVTables()
	InitializeBLT()
	InitializeKeyboardLayouts()
	ReadOptionsControls()
	ChooseKeyboardLayout()
}
