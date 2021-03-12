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
	"errors"
	"fmt"
	"sort"
)

func (c *Creature) Look(b Board, o Objects, cs Creatures) {
	/* Look is method of Creature (that is supposed to be player).
	   It has to take Board, "global" Objects and Creatures as arguments,
	   because function PrintVector need to call RenderAll function.
	   At first, Look creates new para-vector, with player coords as
	   starting point, and dynamic end position.
	   Then ComputeVector checks what tiles are present
	   between Start and End, and adds their coords to vector values.
	   Line from Vector is drawn, then game waits for player input,
	   that will change position of "looking" cursors.
	   Loop breaks with Escape, Space or Enter input. */
	startX, startY := c.X, c.Y
	targetX, targetY := startX, startY
	fov := FOVLength
	if b[cs[0].X][cs[0].Y].Hides == true {
		fov = FOVLengthShort
	}
	for {
		var mon = []string{}
		var hps = []int{}
		var obj = []string{}
		var til = []string{}
		vec, err := NewVector(startX, startY, targetX, targetY)
		if err != nil {
			fmt.Println(err)
		}
		_ = ComputeVector(vec)
		_, _, _, _ = ValidateVector(vec, b, cs, o, StrLook)
		PrintVector(vec, VectorWhyInspect, VectorColorNeutral, VectorColorNeutral, b, o, cs)
		if b[targetX][targetY].Explored == true {
			tt, cc, oo := GetAllThingsFromTile(targetX, targetY, b, cs, o)
			if IsInFOV(b, c.X, c.Y, targetX, targetY, fov) == true {
				for _, v := range cc {
					s := "[color=" + v.Color + "]" + v.Char + "[/color] " + v.Name + " "
					mon = append(mon, s)
					hp := CalcHPPercent(v.HPCurrent, v.HPMax)
					hps = append(hps, hp)
				}
			}
			for _, v := range oo {
				s := "[color=" + v.Color + "]" + v.Char + "[/color] " + v.Name + " "
				obj = append(obj, s)
			}
			if tt != nil {
				s := "[color=" + tt.Color + "]" + tt.Char + "[/color] " + tt.Name + " "
				til = append(til, s)
			}
		}
		PrintLookingMessage(mon, obj, til, hps)
		key := ReadInput()
		if key == blt.TK_ESCAPE || key == blt.TK_ENTER || key == blt.TK_SPACE {
			break
		}
		CursorMovement(&targetX, &targetY, key)
	}
}

func PrintLookingMessage(monstersSlice, objectsSlice, tilesSlice []string, hpSlice []int) {
	/* */
	hpSymbol := " "
	y := 0
	for i, v := range monstersSlice {
		if hpSlice[i] >= 80 {
			hpSymbol = "[color=dark green]Ⅲ[/color]"
		} else if hpSlice[i] >= 40 {
			hpSymbol = "[color=dark yellow]Ⅱ[/color]"
		} else {
			hpSymbol = "[color=dark red]Ⅰ[/color]"
		}
		blt.Print(UIPosX, (UIPosY+9+y)*UIFontSpacingY, "[font=ui]"+v+
			" "+hpSymbol+"[/font]")
		y++
	}
	for _, v := range objectsSlice {
		blt.Print(UIPosX, (UIPosY+9+y)*UIFontSpacingY, "[font=ui]"+v+"[/font]")
		y++
	}
	for _, v := range tilesSlice {
		blt.Print(UIPosX, (UIPosY+9+y)*UIFontSpacingY, "[font=ui]"+v+"[/font]")
		y++
	}
	blt.Refresh()
}

func (c *Creature) Target(b Board, o *Objects, cs *Creatures, dist string, fovLength int) bool {
	/* Target is method of Creature, that takes game map, objects, and
	   creatures as arguments. Returns bool that serves as indicator if
	   action took some time or not.
	   This method is "the big one", general, for handling targeting.
	   In short, player starts targetting, line is drawn from player
	   to monster, then function waits for input (confirmation - "fire",
	   breaking the loop, or continuing).
	   Explicitly:
	   - creates list of all potential targets in fov
	    * tries to automatically last target, but
	    * if fails, it targets the nearest enemy
	   - draws line between source (receiver) and target (coords)
	    * creates new vector
	    * checks if it is valid - monsterHit should not be nil
	    * prints brensenham's line (ie so-called "vector")
	   - waits for player input
	    * if player cancels, function ends
	    * if player confirms, valley is shoot (in target, or empty space)
	    * if valley is shot in empty space, vector is extrapolated to check
	      if it will hit any target
	    * player can switch between targets as well; it targets
	      next target automatically; at first, only monsters that are
	      valid target (ie clean shot is possible), then monsters that
	      are in range and fov, but line of shot is not clear
	    * in other cases, game will try to move cursor; invalid input
	      is ignored */
	finderLength := fovLength
	if dist == StrMelee {
		finderLength = 1
	}
	fov := FOVLength
	if b[c.X][c.Y].Hides == true {
		fov = FOVLengthShort
	}
	turnSpent := false
	var target *Creature
	targets := c.FindTargets(finderLength, b, *cs, *o)
	if dist == StrRanged || dist == StrThrowable {
		if Game.LastTarget != nil && Game.LastTarget != c &&
			IsInFOV(b, c.X, c.Y, Game.LastTarget.X, Game.LastTarget.Y, fov) == true {
			target = Game.LastTarget
		} else {
			var err error
			target, err = c.FindTarget(targets)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		if Game.LastTarget != nil && Game.LastTarget != c &&
			IsInFOV(b, c.X, c.Y, Game.LastTarget.X, Game.LastTarget.Y, fov) == true &&
			(c.X == Game.LastTarget.X || c.Y == Game.LastTarget.Y) {
			target = Game.LastTarget
		} else {
			var err error
			target, err = c.FindTarget(targets)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	targetX, targetY := target.X, target.Y
	for {
		vec, err := NewVector(c.X, c.Y, targetX, targetY)
		if err != nil {
			fmt.Println(err)
		}
		_ = ComputeVector(vec)
		valid, _, monsterHit, _ := ValidateVector(vec, b, targets, *o, dist)
		PrintVector(vec, VectorWhyTarget, VectorColorGood, VectorColorBad, b, *o, *cs)
		if monsterHit != nil {
			hp := CalcHPPercent(monsterHit.HPCurrent, monsterHit.HPMax)
			s := "[color=" + monsterHit.Color + "]" + monsterHit.Char +
				"[/color] " + monsterHit.Name + " "
			PrintLookingMessage([]string{s}, nil, nil, []int{hp})
		} else if targetX == (*cs)[0].X && targetY == (*cs)[0].Y {
			hp := CalcHPPercent((*cs)[0].HPCurrent, (*cs)[0].HPMax)
			s := "[color=" + (*cs)[0].Color + "]" + (*cs)[0].Char + "[/color] " +
				(*cs)[0].Name + " "
			PrintLookingMessage([]string{s}, nil, nil, []int{hp})
		}
		key := ReadInput()
		if key == blt.TK_ESCAPE {
			break
		}
		if (key == blt.TK_F && dist == StrRanged) ||
			(key == blt.TK_D && dist == StrMelee) ||
			(key == blt.TK_F && dist == StrThrowable) ||
			(key == blt.TK_ENTER && dist == StrThrowable) {
			monsterAimed := FindMonsterByXY(targetX, targetY, *cs)
			if dist == StrRanged {
				if c.RangedCurAmmo <= 0 {
					AddMessage("[color=dark yellow]You have no ammo. Reload!")
					break
				}
			} else if dist == StrThrowable {
				if c.ThrowablesCur <= 0 {
					AddMessage("You don't have anything to throw!")
					break
				}
			}
			if monsterAimed != nil && monsterAimed != c && monsterAimed.HPCurrent > 0 && valid == true {
				Game.LastTarget = monsterAimed
				monsterAimed.Staggered = RandRange(monsterAimed.StaggeredMin, monsterAimed.StaggeredMax)
			} else {
				if monsterAimed == c {
					break // Do not hurt yourself.
				}
				if monsterHit != nil {
					if monsterHit.HPCurrent > 0 {
						Game.LastTarget = monsterHit
						monsterHit.Staggered = RandRange(monsterHit.StaggeredMin, monsterHit.StaggeredMax)
					}
				} else {
					vx, vy := FindVectorDirection(vec)
					v := ExtrapolateVector(vec, vx, vy)
					_, _, monsterHitIndirectly, _ := ValidateVector(v, b, targets, *o, dist)
					if monsterHitIndirectly != nil {
						monsterHitIndirectly.Staggered = RandRange(monsterHitIndirectly.StaggeredMin, monsterHitIndirectly.StaggeredMax)
					}
				}
			}
			turnSpent = true
			break
		} else if key == blt.TK_TAB {
			monster := FindMonsterByXY(targetX, targetY, *cs)
			if monster != nil {
				target = NextTarget(monster, targets)
			} else {
				target = NextTarget(target, targets)
			}
			targetX, targetY = target.X, target.Y
			continue // Switch target
		}
		if dist == StrRanged || dist == StrThrowable {
			CursorMovement(&targetX, &targetY, key)
		} else {
			// TODO: behaviour like below should be default one,
			// with customizable distance between target and source.
			CursorMovement(&targetX, &targetY, key)
			if targetX >= c.X+2 {
				targetX = c.X + 1
			} else if targetX <= c.X-2 {
				targetX = c.X - 1
			}
			if targetY >= c.Y+2 {
				targetY = c.Y + 1
			} else if targetY <= c.Y-2 {
				targetY = c.Y - 1
			}
		}
	}
	if turnSpent == true {
		if dist == StrRanged {
			c.RangedCurAmmo--
		}
		if dist == StrThrowable {
			c.ThrowablesCur--
		}
	}
	return turnSpent
}

func CursorMovement(x, y *int, key int) {
	/* CursorMovement is function that takes pointers to coords, and
	   int-based user input. It uses MoveCursor function to
	   modify original values. */
	switch key {
	case blt.TK_UP:
		MoveCursor(x, y, 0, -1)
	case blt.TK_RIGHT:
		MoveCursor(x, y, 1, 0)
	case blt.TK_DOWN:
		MoveCursor(x, y, 0, 1)
	case blt.TK_LEFT:
		MoveCursor(x, y, -1, 0)
	}
}

func MoveCursor(x, y *int, dx, dy int) {
	/* Function MoveCursor takes pointers to coords, and
	   two other ints as direction indicators.
	   It adds direction to coordinate, checks if it is in
	   map bounds, and modifies original values accordingly.
	   This function is called by CursorMovement. */
	newX, newY := *x+dx, *y+dy
	if newX < 0 || newX >= MapSizeX {
		newX = *x
	}
	if newY < 0 || newY >= MapSizeY {
		newY = *y
	}
	*x, *y = newX, newY
}

func (c *Creature) FindTargets(length int, b Board, cs Creatures, o Objects) Creatures {
	/* FindTargets is method of Creature that takes several arguments:
	   length (that is supposed to be max range of attack), and: map, creatures,
	   objects. Returns list of creatures.
	   At first, method creates list of all monsters im c's field of view.
	   Then, this list is divided to two: first, with all "valid" targets
	   (clean (without obstacles) line between c and target) and second,
	   with all other monsters that remains in fov.
	   Both slices are sorted by distance from receiver, then merged.
	   It is necessary for autotarget feature - switching between targets
	   player will start from the nearest valid target, to the farthest valid target;
	   THEN, it will start to target "invalid" targets - again,
	   from nearest to farthest one. */
	targets := c.MonstersInFov(b, cs)
	targetable, unreachable := c.MonstersInRange(b, targets, o, length)
	sort.Slice(targetable, func(i, j int) bool {
		return targetable[i].DistanceBetweenCreatures(c) <
			targetable[j].DistanceBetweenCreatures(c)
	})
	sort.Slice(unreachable, func(i, j int) bool {
		return unreachable[i].DistanceBetweenCreatures(c) <
			unreachable[j].DistanceBetweenCreatures(c)
	})
	targets = nil
	targets = append(targets, targetable...)
	targets = append(targets, unreachable...)
	return targets
}

func (c *Creature) FindTarget(targets Creatures) (*Creature, error) {
	/* FindTarget is method of Creature that takes Creatures as arguments.
	   It returns specific Creature and error.
	   "targets" is supposed to be slice of Creature in player's fov,
	   sorted as explained in FindTargets docstring.
	   If this slice is empty, the target is set to receiver. If not,
	   it tries to target lastly targeted Creature. If it is not possible,
	   it targets first element of slice, and marks it as LastTarget.
	   This method throws an error if it can not find any target,
	   even including receiver. */
	var target *Creature
	if len(targets) == 0 {
		target = c
	} else {
		if Game.LastTarget != nil && CreatureIsInSlice(Game.LastTarget, targets) {
			target = Game.LastTarget
		} else {
			target = targets[0]
			Game.LastTarget = target
		}
	}
	var err error
	if target == nil {
		txt := TargetNilError(c, targets)
		err = errors.New("Could not find target, even the 'self' one." + txt)
	}
	return target, err
}

func NextTarget(target *Creature, targets Creatures) *Creature {
	/* Function NextTarget takes specific creature (target) and slice of creatures
	   (targets) as arguments. It tries to find the *next* target (used
	   with switching between targets, for example using Tab key).
	   At the end, it returns the next creature. */
	i, _ := FindCreatureIndex(target, targets)
	var t *Creature
	length := len(targets)
	if length > i+1 {
		t = targets[i+1]
	} else if length == 0 {
		t = target
	} else {
		t = targets[0]
	}
	return t
}

func (c *Creature) MonstersInRange(b Board, cs Creatures, o Objects,
	length int) (Creatures, Creatures) {
	/* MonstersInRange is method of Creature. It takes global map, Creatures
	   and Objects, and length (range indicator) as its arguments. It returns
	   two slices - one with monsters that are in range, and one with
	   monsters out of range.
	   At first, two empty slices are created, then function starts iterating
	   through Creatures from argument. It creates new vector from source (c)
	   to target, adds monster to proper slice. It also validates vector
	   (ie, won't add monster hidden behind wall) and skips all dead monsters. */
	var inRange = Creatures{}
	var outOfRange = Creatures{}
	for i, v := range cs {
		vec, err := NewVector(c.X, c.Y, v.X, v.Y)
		if err != nil {
			fmt.Println(err)
		}
		if ComputeVector(vec) <= length+1 { // "+1" is necessary due Vector values.
			dist := StrRanged // or StrThrowable - should it be passed as arg?
			if length <= 1 {
				dist = StrMelee
			}
			valid, _, _, _ := ValidateVector(vec, b, cs, o, dist)
			if cs[i].HPCurrent <= 0 {
				continue
			}
			if valid == true {
				inRange = append(inRange, cs[i])
			} else {
				outOfRange = append(outOfRange, cs[i])
			}
		}
	}
	return inRange, outOfRange
}

func ZeroLastTarget(c *Creature) {
	/* LastTarget is global variable (will be incorporated into
	   player struct in future). Function ZeroLastTarget changes
	   last target to nil, is last target matches creature
	   passed as argument. */
	if Game.LastTarget == c {
		Game.LastTarget = nil
	}
}
