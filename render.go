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
	"os"
	"strconv"

	blt "bearlibterminal"
)

const (
	/* Constant values for layers. Their usage is optional,
	   but (for now, at leas) recommended, because default
	   rendering functions depends on these values.
	   They are important for proper clearing characters
	   that should not be displayed, as, for example,
	   bracelet under the monster. */
	_ = iota
	UILayer
	BoardLayer
	DeadLayer
	ObjectsLayer
	CreaturesLayer
	PlayerLayer
	LookLayer
	OverlayLayer
)

const (
	_ = iota
	HiddenInTunnel
	FinishedGameWithAllTreasures
	FinishedGameWithSomeTreasures
	FinishedGameWithoutTreasures
	FinishedGameWithoutExploringAllTiles
	PlayerDied
	PrintHighScores
)

func PrintOverlay(b Board, situation int, c *Creature) {
	blt.Layer(BoardLayer)
	blt.ClearArea(0, 0, MapSizeX*GameFontSpacingX, MapSizeY*GameFontSpacingY)
	blt.Layer(DeadLayer)
	blt.ClearArea(0, 0, MapSizeX*GameFontSpacingX, MapSizeY*GameFontSpacingY)
	blt.Layer(ObjectsLayer)
	blt.ClearArea(0, 0, MapSizeX*GameFontSpacingX, MapSizeY*GameFontSpacingY)
	blt.Layer(CreaturesLayer)
	blt.ClearArea(0, 0, MapSizeX*GameFontSpacingX, MapSizeY*GameFontSpacingY)
	blt.Layer(PlayerLayer)
	blt.ClearArea(0, 0, MapSizeX*GameFontSpacingX, MapSizeY*GameFontSpacingY)
	blt.Layer(OverlayLayer)
	blt.ClearArea(0, 0, MapSizeX*GameFontSpacingX, MapSizeY*GameFontSpacingY)
	blt.Refresh()
	blt.Layer(OverlayLayer)
	if situation == HiddenInTunnel {
		treasures := 0
		unexplored := 0
		for x := 0; x < MapSizeX; x++ {
			for y := 0; y < MapSizeY; y++ {
				if b[x][y].Treasure == true {
					treasures++
				}
				if b[x][y].Explored == false {
					unexplored++
				}
			}
		}
		msg1 := "[font=ui]You are about to hide in the old tunnel."
		msg2 := ""
		msg3 := ""
		msg4 := ""
		cc := -1
		if unexplored == 0 && treasures == 0 {
			msg2 = "[font=ui]You managed to stole back all the items\nrobbed by the invaders."
			msg3 = "[font=ui]Good job, kid.\nNow all you need to do is to wait\nuntil the vikings will leave this place."
			msg4 = "[font=ui](1) Wait..."
			cc = 0
		}
		if unexplored == 0 && treasures > 0 {
			msg2 = "[font=ui]The vikings still have the stolen goods..."
			msg3 = "[font=ui]Are you going to go back to the keep,\nor rather play it safe and wait\nuntil the invaders will leave this place?"
			msg4 = "[font=ui](1) Go back   (2) Wait"
			cc = 1
		}
		if unexplored > 0 {
			msg2 = "[font=ui]There are still places in this keep that you did not visit."
			msg3 = "[font=ui]Are you going to go back to the keep,\nor rather play it safe and wait\nuntil the invaders will leave this place?"
			msg4 = "[font=ui](1) Go back   (2) Wait"
			cc = 2
		}
		SmartPrint(6, 3, UIEntity, msg1)
		SmartPrint(6, 6, UIEntity, msg2)
		SmartPrint(6, 10, UIEntity, msg3)
		SmartPrint(6, 15, UIEntity, msg4)
		blt.Refresh()
		for {
			key := ReadInput()
			if key == blt.TK_1 {
				if cc == 0 {
					PrintOverlay(b, FinishedGameWithAllTreasures, c)
				}
				if cc == 1 || cc == 2 {
					break
				}
			} else if key == blt.TK_2 {
				if cc == 0 {
					continue
				} else if cc == 1 {
					if c.StoleAnything {
						PrintOverlay(b, FinishedGameWithSomeTreasures, c)
					} else {
						PrintOverlay(b, FinishedGameWithoutTreasures, c)
					}
				} else if cc == 2 {
					if c.StoleAnything {
						PrintOverlay(b, FinishedGameWithoutExploringAllTiles, c)
					} else {
						PrintOverlay(b, FinishedGameWithoutTreasures, c)
					}
				}
			}
		}
	}
	if situation == FinishedGameWithAllTreasures {
		Game.SpecialPoints = append(Game.SpecialPoints, AllTreasuresStolen)
		Game.Points = Game.CalculatePoints()
		UpdateScores()
		SaveScores(HighScores)
		msg := "[font=ui]Congratulations! You finished the game,\ncollecting all the stolen treasures.\n\nPress any key to exit..."
		SmartPrint(6, 8, UIEntity, msg)
		blt.Refresh()
		_ = ReadInput()
		PrintOverlay(b, PrintHighScores, c)
	}
	if situation == FinishedGameWithSomeTreasures {
		Game.Points = Game.CalculatePoints()
		UpdateScores()
		SaveScores(HighScores)
		msg := "[font=ui]Congratulations! You finished the game,\nstealing back some robbed treasures.\n\nPress any key to exit..."
		SmartPrint(6, 8, UIEntity, msg)
		blt.Refresh()
		_ = ReadInput()
		PrintOverlay(b, PrintHighScores, c)
	}
	if situation == FinishedGameWithoutTreasures {
		Game.SpecialPoints = append(Game.SpecialPoints, StoleNothing)
		Game.Points = Game.CalculatePoints()
		UpdateScores()
		SaveScores(HighScores)
		msg := "[font=ui]You finished the game.\nMaybe you didn't recover any valuables,\nbut you are alive, at least.\n\nPress any key to exit..."
		SmartPrint(6, 8, UIEntity, msg)
		blt.Refresh()
		_ = ReadInput()
		PrintOverlay(b, PrintHighScores, c)
	}
	if situation == FinishedGameWithoutExploringAllTiles {
		Game.Points = Game.CalculatePoints()
		UpdateScores()
		SaveScores(HighScores)
		msg := "[font=ui]You finished the game.\nYou didn't have a chance to explore every corner,\nbut you are alive, at least.\n\nPress any key to exit..."
		SmartPrint(6, 8, UIEntity, msg)
		blt.Refresh()
		_ = ReadInput()
		PrintOverlay(b, PrintHighScores, c)
	}
	if situation == PlayerDied {
		Game.SpecialPoints = append(Game.SpecialPoints, Dieded)
		Game.Points = Game.CalculatePoints()
		UpdateScores()
		SaveScores(HighScores)
		msg := "[font=ui]The vikings caught you unaware.\nYou couldn't fight back, and\nand the enemies had no mercy.\n\nPress any key to exit..."
		SmartPrint(6, 8, UIEntity, msg)
		blt.Refresh()
		_ = ReadInput()
		PrintOverlay(b, PrintHighScores, c)
	}
	if situation == PrintHighScores {
		SmartPrint(6, 3, UIEntity, "[font=ui]HIGH SCORES")
		for i, v := range HighScores.Entries {
			if i >= 10 {
				break
			}
			SmartPrint(6, 4+i, UIEntity,
					"[font=ui]" + strconv.Itoa(i+1) + ". " + strconv.Itoa(v.Points) + "  -  " + v.PlayerName)
		}
		blt.Refresh()
		_ = ReadInput()
		blt.Close()
		os.Exit(0)
	}
}

func PrintBoard(b Board, c Creatures) {
	/* Function PrintBoard is used in RenderAll function.
	   Takes level map and list of monsters as arguments
	   and iterates through Board.
	   It has to check for "]" and "[" characters, because
	   BearLibTerminal uses these symbols for config.
	   Instead of checking it here, one could just remember to
	   always pass "]]" instead of "]".
	   Prints every tile on its coords if certain conditions are met:
	   is Explored already, and:
	   - is in player's field of view (prints "normal" color) or
	   - is AlwaysVisible (prints dark color). */
	fov := FOVLength
	if b[c[0].X][c[0].Y].Hides == true {
		fov = FOVLengthShort
	}
	for x := 0; x < MapSizeX; x++ {
		for y := 0; y < MapSizeY; y++ {
			// Technically, "t" is new variable with own memory address...
			t := b[x][y] // Should it be *b[x][y]?
			blt.Layer(t.Layer)
			if t.Explored == true {
				ch := t.Char
				if t.Char == "[" || t.Char == "]" {
					ch = t.Char + t.Char
				}
				if t.Treasure == true {
					glyph := "[font=game][color=" + t.TreasureCol + "]" + t.TreasureChar
					SmartPrint(t.X, t.Y, MapEntity, glyph)
				} else {
					if IsInFOV(b, c[0].X, c[0].Y, t.X, t.Y, fov) == true {
						glyph := "[font=game][color=" + t.Color + "]" + ch
						SmartPrint(t.X, t.Y, MapEntity, glyph)
					} else {
						if t.AlwaysVisible == true {
							glyph := "[font=game][color=" + t.ColorDark + "]" + ch
							SmartPrint(t.X, t.Y, MapEntity, glyph)
						}
					}
				}
			}
		}
	}
}

func PrintObjects(b Board, o Objects, c Creatures) {
	/* Function PrintObjects is used in RenderAll function.
	   Takes map of level, slice of objects, and all monsters
	   as arguments.
	   Iterates through Objects.
	   It has to check for "]" and "[" characters, because
	   BearLibTerminal uses these symbols for config.
	   Instead of checking it here, one could just remember to
	   always pass "]]" instead of "]".
	   Prints every object on its coords if certain conditions are met:
	   AlwaysVisible bool is set to true, or is in player fov. */
	fov := FOVLength
	if b[c[0].X][c[0].Y].Hides == true {
		fov = FOVLengthShort
	}
	for _, v := range o {
		if (IsInFOV(b, c[0].X, c[0].Y, v.X, v.Y, fov) == true) ||
			((v.AlwaysVisible == true) && (b[v.X][v.Y].Explored == true)) {
			blt.Layer(v.Layer)
			ch := v.Char
			if v.Char == "]" || v.Char == "[" {
				ch = v.Char + v.Char
			}
			glyph := "[font=game][color=" + v.Color + "]" + ch
			SmartPrint(v.X, v.Y, ObjectEntity, glyph)
			for i := 0; i < v.Layer; i++ {
				blt.Layer(i)
				SmartClear(v.X, v.Y, ObjectEntity)
			}
		}
	}
}

func PrintCreatures(b Board, c Creatures) {
	/* Function PrintCreatures is used in RenderAll function.
	   Takes map of level and slice of Creatures as arguments.
	   Iterates through Creatures.
	   It has to check for "]" and "[" characters, because
	   BearLibTerminal uses these symbols for config.
	   Instead of checking it here, one could just remember to
	   always pass "]]" instead of "]".
	   Checks for every creature on its coords if certain conditions are met:
	   AlwaysVisible bool is set to true, or is in player fov.
	   Now, it prints three times. First time, to print corpses.
	   Then, to print living creatures (except the player).
	   Last time, it prints player.
	   This way, it is ensured that corpses will always be hidden
	   under the living monsters. */
	// Print corpses
	fov := FOVLength
	if b[c[0].X][c[0].Y].Hides == true {
		fov = FOVLengthShort
	}
	for i, v := range c {
		if i == 0 {
			continue // Player will be drawn separately
		}
		if v.Layer != DeadLayer {
			continue
		}
		if (IsInFOV(b, c[0].X, c[0].Y, v.X, v.Y, fov) == true) ||
			(v.AlwaysVisible == true) {
			blt.Layer(v.Layer)
			ch := v.Char
			if v.Char == "]" || v.Char == "[" {
				ch = v.Char + v.Char
			}
			glyph := "[font=game][color=" + v.Color + "]" + ch
			SmartPrint(v.X, v.Y, MonsterEntity, glyph)
			for j := 0; j < v.Layer; j++ {
				blt.Layer(j)
				SmartClear(v.X, v.Y, MonsterEntity)
			}
		}
	}
	// Print living creatures
	for i, v := range c {
		if i == 0 {
			continue // Player will be drawn separately
		}
		if v.Layer == DeadLayer {
			continue
		}
		if (IsInFOV(b, c[0].X, c[0].Y, v.X, v.Y, fov) == true) ||
			(v.AlwaysVisible == true) {
			blt.Layer(v.Layer)
			ch := v.Char
			if v.Char == "]" || v.Char == "[" {
				ch = v.Char + v.Char
			}
			glyph := "[font=game][color=" + v.Color + "]" + ch
			SmartPrint(v.X, v.Y, MonsterEntity, glyph)
			for j := 0; j < v.Layer; j++ {
				blt.Layer(j)
				SmartClear(v.X, v.Y, MonsterEntity)
			}
		}
		if IsInFOV(b, c[0].X, c[0].Y, v.X, v.Y, fov) == false {
			if v.Char == "‼" {
				blt.Layer(v.Layer)
				glyph := "[font=game][color=" + v.Color + "]" + v.Char
				SmartPrint(v.X, v.Y, MonsterEntity, glyph)
				for j := 0; j < v.Layer; j++ {
					blt.Layer(j)
					SmartClear(v.X, v.Y, MonsterEntity)
				}
			}
		}
	}
	// Print player
	blt.Layer(PlayerLayer)
	playerColor := c[0].Color
	if b[c[0].X][c[0].Y].Hides == true {
		if c[0].Hidden {
			playerColor = "darkest gray"
		} else {
			playerColor = "dark gray"
		}
	}
	SmartPrint(c[0].X, c[0].Y,
		MonsterEntity, "[font=game][color="+playerColor+"]"+c[0].Char)
	for i := 0; i < PlayerLayer; i++ {
		blt.Layer(i)
		SmartClear(c[0].X, c[0].Y, MonsterEntity)
	}
}

func PrintUI(c *Creature) {
	/* Function PrintUI takes *Creature (it's supposed to be player) as argument.
	   It prints UI infos on the right side of screen.
	   For now its functionality is very modest, but it will expand when
	   new elements of game mechanics will be introduced. So, for now, it
	   provides only one basic, yet essential information: player's HP. */
	//REMEMBER: [offset=x,y]
	blt.Layer(UILayer)
	hp := ""
	for i := 0; i < c.HPCurrent; i++ {
		hp += "[color=red]♥"
	}
	for i := c.HPCurrent; i < c.HPMax; i++ {
		hp += "[/color][color=darker red]♡"
	}
	blt.Print(UIPosX, UIPosY+UIFontSpacingY, "[font=game]"+hp)
	enc := ""
	if c.LightItem1 == true {
		enc += "[color=yellow]" + TreasureCharLight
	} else {
		enc += "[color=gray]" + TreasureCharLight
	}
	if c.LightItem2 == true {
		enc += "[color=yellow]" + TreasureCharLight
	} else {
		enc += "[color=gray]" + TreasureCharLight
	}
	if c.LightItem3 == true {
		enc += "[color=yellow]" + TreasureCharLight
	} else {
		enc += "[color=gray]" + TreasureCharLight
	}
	if c.MediumItem1 == true {
		enc += "[color=yellow]" + TreasureCharMedium
	} else {
		enc += "[color=gray]" + TreasureCharMedium
	}
	if c.MediumItem2 == true {
		enc += "[color=yellow]" + TreasureCharMedium
	} else {
		enc += "[color=gray]" + TreasureCharMedium
	}
	if c.HeavyItem1 == true {
		enc += "[color=yellow]" + TreasureCharHeavy
	} else {
		enc += "[color=gray]" + TreasureCharHeavy
	}
	enc += "[color=lighter gray] → "
	encm := 0
	if c.LightItem1 && c.LightItem2 && c.LightItem3 {
		encm += 1
	}
	if c.MediumItem1 && c.MediumItem2 {
		encm += 1
	}
	if c.HeavyItem1 {
		encm += 1
	}
	switch encm {
	case 0:
		enc += "[color=dark green]▏"
	case 1:
		enc += "[color=dark yellow]▍"
	case 2:
		enc += "[color=dark orange]▋"
	case 3:
		enc += "[color=dark red]█"
	}
	blt.Print(UIPosX, UIPosY+(2*UIFontSpacingY), "[font=game]"+enc)

	/*
		name := "Player"
		blt.Print(UIPosX, UIPosY * UIFontSpacingY, "[font=ui]" + name)
		hp := "HP: " + strconv.Itoa(c.HPCurrent) + "\\" + strconv.Itoa(c.HPMax)
		blt.Print(UIPosX, (UIPosY+1) * UIFontSpacingY, "[font=ui][color=red]" + hp)
		turn := "Turns: " + strconv.Itoa(Game.TurnCounter)
		if Game.BreakTime > 0 {
			turn = turn + " [color=darker yellow]BREAK[/color]"
		}
		blt.Print(UIPosX, (UIPosY+2) * UIFontSpacingY, "[font=ui]" + turn)
		monsters := "Killed: " + strconv.Itoa(len(Game.MonstersKilled))
		blt.Print(UIPosX, (UIPosY+3) * UIFontSpacingY, "[font=ui]" + monsters)
		wave := "Wave:" + strconv.Itoa(Game.WaveNo)
		blt.Print(UIPosX, (UIPosY+4) * UIFontSpacingY, "[font=ui]" + wave)
		ranged := c.CreateRangedString()
		blt.Print(UIPosX, (UIPosY+5) * UIFontSpacingY + (UIFontSpacingY / 2), "[font=ui]" + ranged)
		throwables := c.CreateThrowablesString()
		blt.Print(UIPosX, (UIPosY+7) * UIFontSpacingY, "[font=ui]" + throwables)
		status := c.CreateStatusString()
		blt.Print(UIPosX, (UIPosY+9) * UIFontSpacingY, "[font=ui]" + status)
	*/
}

func PrintLog() {
	/* Function PrintLog prints game messages at the bottom of screen. */
	blt.Layer(UILayer)
	PrintMessages(LogPosX, LogPosY)
}

func RenderAll(b Board, o Objects, c Creatures) {
	/* Function RenderAll prints every tile and character on game screen.
	   Takes board slice (ie level map), slice of objects, and slice of creatures
	   as arguments.
	   At first, it clears whole terminal window, then uses arguments:
	   CastRays (for raycasting FOV) of first object (assuming that it is player),
	   then calls functions for printing map, objects and creatures.
	   Calls PrintLog that writes message log.
	   At the end, RenderAll calls blt.Refresh() that makes
	   changes to the game window visible. */
	blt.Clear()
	CastRays(b, c[0].X, c[0].Y)
	PrintBoard(b, c)
	PrintObjects(b, o, c)
	PrintCreatures(b, c)
	PrintUI((c)[0])
	PrintLog()
	blt.Refresh()
}
