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
	"errors"
	"fmt"
	"io/ioutil"
	"unicode/utf8"

	blt "bearlibterminal"
)

const (
	// Special characters.
	CorpseChar = "%"
)

const (
	// Speed values
	SpeedVerySlow = iota - 2
	SpeedSlow
	SpeedNormal
	SpeedFast
	SpeedVeryFast
)

var SpeedValues = map[int]int{
	SpeedVerySlow: 2,
	SpeedSlow:     4,
	SpeedNormal:   1,
	SpeedFast:     4,
	SpeedVeryFast: 2,
}

type Creature struct {
	/* Creatures are living objects that
	   moves, attacks, dies, etc. */
	BasicProperties
	VisibilityProperties
	CollisionProperties
	FighterProperties
	EquipmentComponent
}

// Creatures holds every creature on map.
type Creatures []*Creature

func NewCreature(x, y int, b Board, monsterFile string) (*Creature, error) {
	/* NewCreature is function that returns new Creature from
	   json file passed as argument. It replaced old code that
	   was encouraging hardcoding data in go files.
	   Errors returned by json package are not very helpful, and
	   hard to work with, so there is lazy panic for them. */
	var monster = &Creature{}
	err := CreatureFromJson(CreaturesPathJson+monsterFile, monster)
	if err != nil {
		fmt.Println(err)
		panic(-1)
	}
	monster.X, monster.Y = x, y
	var err2 error
	if monster.Layer < 0 {
		txt := LayerError(monster.Layer)
		err2 = errors.New("Creature layer is smaller than 0." + txt)
	}
	if monster.Layer != CreaturesLayer {
		txt := LayerWarning(monster.Layer, CreaturesLayer)
		err2 = errors.New("Creature layer is not equal to CreaturesLayer constant." + txt)
	}
	if monster.X < 0 || monster.X >= MapSizeX || monster.Y < 0 || monster.Y >= MapSizeY {
		txt := CoordsError(monster.X, monster.Y)
		err2 = errors.New("Creature coords is out of window range." + txt)
	}
	if utf8.RuneCountInString(monster.Char) != 1 {
		txt := CharacterLengthError(monster.Char)
		err2 = errors.New("Creature character string length is not equal to 1." + txt)
	}
	if monster.HPMax < 0 {
		txt := InitialHPError(monster.HPMax)
		err2 = errors.New("Creature HPMax is smaller than 0." + txt)
	}
	if monster.Attack < 0 {
		txt := InitialAttackError(monster.Attack)
		err2 = errors.New("Creature attack value is smaller than 0." + txt)
	}
	if monster.Defense < 0 {
		txt := InitialDefenseError(monster.Defense)
		err2 = errors.New("Creature defense value is smaller than 0." + txt)
	}
	if monster.ChallengeLevel < 0 {
		txt := MonsterChallengeLevelError(monster.ChallengeLevel)
		err2 = errors.New("Creature challenge level is smaller than 0." + txt)
	}
	if monster.Equipment == nil {
		monster.Equipment = Objects{}
	}
	if monster.Inventory == nil {
		monster.Inventory = Objects{}
	}
	if monster.PatrolPoints == nil {
		monster.PatrolPoints = [][]int{{monster.X, monster.Y}}
		steps := RandRange(2, 3)
		for i := 0; i < steps-1; i++ {
			var x, y int
			for {
				tx := RandRange(MinDistanceBetweenPatrolPoints, MaxDistanceBetweenPatrolPoints)
				if RandInt(100) < 50 {
					tx *= (-1)
				}
				x = monster.PatrolPoints[i][0] + tx
				ty := RandRange(MinDistanceBetweenPatrolPoints, MaxDistanceBetweenPatrolPoints)
				if RandInt(100) < 50 {
					ty *= (-1)
				}
				y = monster.PatrolPoints[i][1] + ty
				if x > 0 && x < MapSizeX && y > 0 && y < MapSizeY && b[x][y].Blocked == false {
					break
				}
			}
			monster.PatrolPoints = append(monster.PatrolPoints, []int{x, y})
		}
	}
	if monster.NextPoint <= 0 {
		monster.NextPoint = 1
	}
	if monster.LastPosition == nil {
		monster.LastPosition = []int{monster.X, monster.Y}
	}
	monster.MaxOutOfFOV = 10
	return monster, err2
}

func (c *Creature) MoveOrAttack(tx, ty int, b Board, o *Objects, all *Creatures) bool {
	/* Method MoveOrAttack decides if Creature will move or attack other Creature;
	   It has *Creature receiver, and takes tx, ty (coords) integers as arguments,
	   and map of current level, and list of all Creatures.
	   Starts by target that is nil, then iterates through Creatures. If there is
	   Creature on targeted tile, that Creature becomes new target for attack.
	   Otherwise, Creature moves to specified Tile.
	   It's supposed to take player as receiver (attack / moving enemies is
	   handled differently - check ai.go and combat.go).
	   For this game, the bump-attack mechanics is disabled. */
	var target *Creature
	turnSpent := false
	for i, _ := range *all {
		if (*all)[i].X == c.X+tx && (*all)[i].Y == c.Y+ty {
			if (*all)[i].HPCurrent > 0 {
				target = (*all)[i]
				break
			}
		}
	}
	if target != nil {
		if c != (*all)[0] { // Disable bump-attack mechanics.
			c.AttackTarget(target, o, &b, all, "")
			turnSpent = true
		} else {
			turnSpent = false
		}
	} else {
		turnSpent = c.Move(tx, ty, b, *all)
	}
	return turnSpent
}

func (c *Creature) Move(tx, ty int, b Board, cs Creatures) bool {
	/* Move is method of Creature; it takes target x, y as arguments;
	   check if next move won't put Creature off the screen, then updates
	   Creature coords. */
	turnSpent := false
	newX, newY := c.X+tx, c.Y+ty
	c.Hidden = false
	if newX >= 0 &&
		newX <= MapSizeX-1 &&
		newY >= 0 &&
		newY <= MapSizeY-1 {
		if b[newX][newY].Blocked == false {
			if c.Stuck == false {
				c.X = newX
				c.Y = newY
				if b[newX][newY].Slows == true {
					c.Stuck = true
				}
			} else {
				c.Stuck = false
			}
			if c.AIType == PlayerAI && b[newX][newY].Hides == true {
				hidden := true
				for i, v := range cs {
					if i == 0 {
						continue
					}
					if IsInFOV(b, v.X, v.Y, c.X, c.Y, FOVLength) && v.AITriggered == AITriggered {
						hidden = false
					}
				}
				c.Hidden = hidden
			}
			turnSpent = true
		} else {
			if c.AIType == PlayerAI && b[newX][newY].Hides == true {
				c.X = newX
				c.Y = newY
				hidden := true
				for i, v := range cs {
					if i == 0 {
						continue
					}
					if IsInFOV(b, v.X, v.Y, c.X, c.Y, FOVLength) && v.AITriggered == AITriggered {
						hidden = false
					}
				}
				c.Hidden = hidden
			}
		}
	}
	return turnSpent
}

func (c *Creature) UseEnvironment(b *Board) bool {
	turnSpent := false
	if c.AIType != PlayerAI {
		return turnSpent
	}
	if (*b)[c.X][c.Y].Name == "Hatch" {
		if c.LightItem1 {
			c.StoleAnything = true
			Game.SmallStolen++
		}
		if c.LightItem2 {
			c.StoleAnything = true
			Game.SmallStolen++
		}
		if c.LightItem3 {
			c.StoleAnything = true
			Game.SmallStolen++
		}
		if c.MediumItem1 {
			c.StoleAnything = true
			Game.MediumStolen++
		}
		if c.MediumItem2 {
			c.StoleAnything = true
			Game.MediumStolen++
		}
		if c.HeavyItem1 {
			c.StoleAnything = true
			Game.HeavyStolen++
		}
		PrintOverlay(*b, HiddenInTunnel, c)
		Game.Breaks++
		turnSpent = true
	}
	return turnSpent
}

func (c *Creature) Drop(b *Board) bool {
	turnSpent := false
	x, y := c.X, c.Y
	if (*b)[x][y].Char != "." || (*b)[x][y].Treasure {
		AddMessage("This tile is already occupied.")
	} else {
		msg := ""
		l, m, h := false, false, false
		if c.LightItem1 || c.LightItem2 || c.LightItem3 {
			l = true
			msg += "(L)ight  "
		}
		if c.MediumItem1 || c.MediumItem2 {
			m = true
			msg += "(M)edium  "
		}
		if c.HeavyItem1 {
			h = true
			msg += "(H)Heavy  "
		}
		if l == false && m == false && h == false {
			AddMessage("You have nothing to drop")
		} else {
			AddMessage("What item do you want to drop here?")
			AddMessage(msg)
			for {
				key := ReadInput()
				if key == blt.TK_L && l {
					if c.LightItem3 {
						c.LightItem3 = false
					} else if c.LightItem2 {
						c.LightItem2 = false
					} else if c.LightItem1 {
						c.LightItem1 = false
					}
					AddMessage("You tossed the pouch from your pocket.")
					(*b)[x][y].Treasure = true
					(*b)[x][y].TreasureCol = "yellow"
					(*b)[x][y].TreasureChar = TreasureCharLight
					turnSpent = true
					break
				} else if key == blt.TK_M && m {
					if c.MediumItem2 {
						c.MediumItem2 = false
					} else if c.MediumItem1 {
						c.MediumItem1 = false
					}
					AddMessage("You stripped the valuables off your belt.")
					(*b)[x][y].Treasure = true
					(*b)[x][y].TreasureCol = "yellow"
					(*b)[x][y].TreasureChar = TreasureCharMedium
					turnSpent = true
					break
				} else if key == blt.TK_H && h {
					c.HeavyItem1 = false
					AddMessage("You dropped the treasure bag on the ground.")
					(*b)[x][y].Treasure = true
					(*b)[x][y].TreasureCol = "yellow"
					(*b)[x][y].TreasureChar = TreasureCharHeavy
					turnSpent = true
					break
				} else {
					break
				}
			}
		}
	}
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
		c.Speed = SpeedFast
	case 1:
		c.Speed = SpeedNormal
	case 2:
		c.Speed = SpeedSlow
	case 3:
		c.Speed = SpeedVerySlow
	}
	return turnSpent
}

func (c *Creature) PickUp(b *Board) bool {
	/* PickUp is method that has *Creature as receiver
	   and slice of *Object as argument.
	   Creature tries to pick object up.
	   If creature stands on object that is possible to pick,
	   object is added to c's inventory, and removed
	   from "global" slice of objects.
	   Picking objects up takes turn only if it is
	   successful attempt. */
	turnSpent := false
	if c.AIType == PlayerAI {
		if (*b)[c.X][c.Y].Treasure == true {
			switch (*b)[c.X][c.Y].TreasureChar {
			case TreasureCharLight:
				if c.LightItem1 && c.LightItem2 && c.LightItem3 {
					AddMessage("Your pockets are full already.")
				} else {
					if c.LightItem1 == false {
						c.LightItem1 = true
					} else if c.LightItem2 == false {
						c.LightItem2 = true
					} else if c.LightItem3 == false {
						c.LightItem3 = true
					}
					(*b)[c.X][c.Y].Treasure = false
					AddMessage("You put the coins in your pocket.")
					turnSpent = true
				}
			case TreasureCharMedium:
				if c.MediumItem1 && c.MediumItem2 {
					AddMessage("You do not have free space on the belt anymore.")
				} else {
					if c.MediumItem1 == false {
						c.MediumItem1 = true
					} else if c.MediumItem2 == false {
						c.MediumItem2 = true
					}
					(*b)[c.X][c.Y].Treasure = false
					AddMessage("You put valuables down to the belt")
					turnSpent = true
				}
			case TreasureCharHeavy:
				if c.HeavyItem1 {
					AddMessage("You can not carry anything else on your back.")
				} else {
					c.HeavyItem1 = true
					(*b)[c.X][c.Y].Treasure = false
					AddMessage("You put the treasure in a bag and put it on your back.")
					turnSpent = true
				}
			}
		} else if (*b)[c.X][c.Y].Char == "…" {
			if c.ThrowablesCur < c.ThrowablesMax {
				c.ThrowablesCur++
				(*b)[c.X][c.Y].Char = "."
				(*b)[c.X][c.Y].Name = "floor"
				(*b)[c.X][c.Y].Color = "lighter grey"
				turnSpent = true
			} else {
				AddMessage("You don't have a place for more pebbles.")
			}
		}
	}
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
		c.Speed = SpeedFast
	case 1:
		c.Speed = SpeedNormal
	case 2:
		c.Speed = SpeedSlow
	case 3:
		c.Speed = SpeedVerySlow
	}
	return turnSpent
}

func (c *Creature) DropFromInventory(objects *Objects, index int) bool {
	/* Drop is method that has Creature as receiver and takes
	   "global" list of objects as main argument, and additional
	   integer that is index of item to be dropped from c's Inventory.
	   At first, turnSpent is set to false, to make it true
	   at the end of function. It may be considered as obsolete WET,
	   because 'return true' would be sufficient, but it is
	   a bit more readable now.
	   Objs is dereferenced objects and it is absolutely necessary
	   to do any actions on these objects.
	   Drop do two things:
	   at first, it adds specific item to the game map,
	   then it removes this item from its owner Inventory. */
	turnSpent := false
	objs := *objects
	if c.AIType == PlayerAI {
		AddMessage("You dropped " + c.Inventory[index].Name + ".")
	}
	// Add item to the map.
	object := c.Inventory[index]
	object.X, object.Y = c.X, c.Y
	objs = append(objs, object)
	*objects = objs
	// Then remove item from inventory.
	copy(c.Inventory[index:], c.Inventory[index+1:])
	c.Inventory[len(c.Inventory)-1] = nil
	c.Inventory = c.Inventory[:len(c.Inventory)-1]
	turnSpent = true
	return turnSpent
}

func (c *Creature) DropFromEquipment(objects *Objects, slot int) bool {
	/* DropFromEquipment is method of *Creature that takes "global" objects,
	   and int (as index) as arguments, and returns bool (result depends if
	   action was successful, therefore if took a turn).
	   This function is very similar to DropFromInventory, but is kept
	   due to explicitness.
	   The difference is that Equipment checks Equipment index, not
	   specific object, so additionally checks for nils, and instead of
	   removing item from slice, makes it nil.
	   This behavior is important, because while Inventory is "dynamic"
	   slice, Equipment is supposed to be "fixed size" - slots are present
	   all the time, but the can be empty (ie nil) or occupied (ie object). */
	turnSpent := false
	objs := *objects
	object := c.Equipment[slot]
	if object == nil {
		return turnSpent // turn is not spent because there is no object to drop
	}
	// else {
	if c.AIType == PlayerAI {
		AddMessage("You removed and dropped " + object.Name + ".")
	}
	// add item to map
	object.X, object.Y = c.X, c.Y
	objs = append(objs, object)
	*objects = objs
	// then remove from slot
	c.Equipment[slot] = nil
	turnSpent = true
	return turnSpent
}

func (c *Creature) EquipItem(o *Object, slot int) (bool, error) {
	/* EquipItem is method of *Creature that takes *Object and int (that is
	   indicator to index of Equipment slot) as arguments; it returns
	   bool and error.
	   At first, EquipItem checks for errors:
	    - if object to equip exists
	    - if this equipment slot is not occupied
	   then equips item and removes it from inventory. */
	var err error
	if o == nil {
		txt := EquipNilError(c)
		err = errors.New("Creature tried to equip *Object that was nil." + txt)
	}
	if c.Equipment[slot] != nil {
		txt := EquipSlotNotNilError(c, slot)
		err = errors.New("Creature tried to equip item into already occupied slot." + txt)
	}
	if o.Slot != slot {
		txt := EquipWrongSlotError(o.Slot, slot)
		err = errors.New("Creature tried to equip item into wrong slot." + txt)
	}
	turnSpent := false
	// Equip item...
	c.Equipment[slot] = o
	// ...then remove it from inventory.
	index, err := FindObjectIndex(o, c.Inventory)
	if err != nil {
		fmt.Println(err)
	}
	copy(c.Inventory[index:], c.Inventory[index+1:])
	c.Inventory[len(c.Inventory)-1] = nil
	c.Inventory = c.Inventory[:len(c.Inventory)-1]
	if c.AIType == PlayerAI {
		AddMessage("You equipped " + o.Name + ".")
	}
	turnSpent = true
	return turnSpent, err
}

func (c *Creature) DequipItem(slot int) (bool, error) {
	/* DequipItem is method of Creature. It is called when receiver is about
	   to dequip weapon from "ready" equipment slot.
	   At first, weapon is added to Inventory, then Equipment slot is set to nil. */
	var err error
	if c.Equipment[slot] == nil {
		txt := DequipNilError(c, slot)
		err = errors.New("Creature tried to DequipItem that was nil." + txt)
	}
	if c.AIType == PlayerAI {
		AddMessage("You dequipped " + c.Equipment[slot].Name + ".")
	}
	turnSpent := false
	c.Inventory = append(c.Inventory, c.Equipment[slot]) //adding items to inventory should have own function, that will check "bounds" of inventory
	c.Equipment[slot] = nil
	turnSpent = true
	return turnSpent, err
}

func (c *Creature) Die(o *Objects, b *Board, cs *Creatures) {
	/* Method Die is called when Creature's HP drops below zero.
	   Die() has *Creature as receiver.
	   Receiver properties changes to fit better to corpse. */
	randCol := RandInt(100)
	if randCol <= 10 {
		(*b)[c.X][c.Y].Color = "red"
		(*b)[c.X][c.Y].ColorDark = "red"
	} else if randCol <= 35 {
		(*b)[c.X][c.Y].Color = "darker red"
		(*b)[c.X][c.Y].ColorDark = "darker red"
	} else {
		(*b)[c.X][c.Y].Color = "dark red"
		(*b)[c.X][c.Y].Color = "dark red"
	}
	if c != (*cs)[0] {
		ZeroLastTarget(c)
		AddMonsterKilled(c)
		AddPointsForMonsterKilled(c)
		i, _ := FindCreatureIndex(c, *cs)
		var temp = Creatures{}
		temp = append(temp, (*cs)[:i]...)
		temp = append(temp, (*cs)[i+1:]...)
		*cs = temp
		Game.LivingMonsters--
	}
}

func FindMonsterByXY(x, y int, c Creatures) *Creature {
	/* Function FindMonsterByXY takes desired coords and list
	   of all available creatures. It iterates through this list,
	   and returns nil or creature that occupies specified coords. */
	var monster *Creature
	for i := 0; i < len(c); i++ {
		if x == c[i].X && y == c[i].Y {
			monster = c[i]
			break
		}
	}
	return monster
}

func AddMonsterKilled(c *Creature) {
	/* Function AddMonsterKilled takes single argument of *Creature.
	   It's called when the monster die, and adds its name
	   to the list of the killed monsters stored in the Game variable. */
	Game.MonstersKilled = append(Game.MonstersKilled, c.Name)
}

func AddPointsForMonsterKilled(c *Creature) {
	Game.KillPoints += c.ChallengeLevel * 10
}

func ReadMonsterFiles() []string {
	var c = []string{}
	fs, err := ioutil.ReadDir("./data/monsters/")
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range fs {
		c = append(c, f.Name())
	}
	return c
}

func SpawnMonsters(b Board, c *Creatures) {
	var x, y int
	monstersNum := RandRange(MonstersMin, MonstersMax)
	//var monsters *Creatures
	var placementX = []int{}
	var placementY = []int{}
	place1 := 0
	place2 := 0
	place3 := 0
	place4 := 0
	for {
		placementX = nil
		placementY = nil
		validPlacement := true
		for i := 0; i < monstersNum; i++ {
			place := RandInt(100)
			if place <= 25 {
				for {
					// "+-4" helps to avoid spawning multiple
					// enemies next to the map center
					x = RandRange(0, (MapSizeX-4) / 2)
					y = RandRange(0, (MapSizeY-4) / 2)
					if b[x][y].Char == "." && IsInFOV(b, (*c)[0].X, (*c)[0].Y, x, y, FOVLength) == false {
						break
					}
				}
				placementX = append(placementX, x)
				placementY = append(placementY, y)
				place1++
			} else if place <= 50 {
				for {
					x = RandRange((MapSizeX+4) / 2, MapSizeX-1)
					y = RandRange(0, (MapSizeY-4) / 2)
					if b[x][y].Char == "." && IsInFOV(b, (*c)[0].X, (*c)[0].Y, x, y, FOVLength) == false {
						break
					}
				}
				placementX = append(placementX, x)
				placementY = append(placementY, y)
				place2++
			} else if place <= 75 {
				for {
					x = RandRange((MapSizeX+4) / 2, MapSizeX-1)
					y = RandRange((MapSizeY+4) / 2, MapSizeY-1)
					if b[x][y].Char == "." && IsInFOV(b, (*c)[0].X, (*c)[0].Y, x, y, FOVLength) == false {
						break
					}
				}
				placementX = append(placementX, x)
				placementY = append(placementY, y)
				place3++
			} else {
				for {
					x = RandRange(0, (MapSizeX-4) / 2)
					y = RandRange((MapSizeY+4) / 2, MapSizeY-1)
					if b[x][y].Char == "." && IsInFOV(b, (*c)[0].X, (*c)[0].Y, x, y, FOVLength) == false {
						break
					}
				}
				placementX = append(placementX, x)
				placementY = append(placementY, y)
				place4++
			}
		}
		if place1 == 0 || place2 == 0 || place3 == 0 || place4 == 0 {
			validPlacement = false
		}
		if validPlacement == false {
			continue
		} else {
			break
		}
	}
	for i := 0; i < monstersNum; i++ {
		m, _ := NewCreature(placementX[i], placementY[i], b, "viking_warrior.json")
		w1, _ := NewObject(0, 0, "weapon1.json")
		w2, _ := NewObject(0, 0, "weapon2.json")
		wm, _ := NewObject(0, 0, "melee.json")
		var monsterEq = EquipmentComponent{Objects{w1, w2, wm}, Objects{}}
		m.EquipmentComponent = monsterEq
		*c = append(*c, m)
	}

/*
	for i := 0; i <= Game.SpawnAmount; i++ {
		if Game.WaveCur < Game.WaveMax {
			for {
				place := RandInt(100)
				if place <= 25 {
					x = 1
					y = RandRange(1, MapSizeY-2)
				} else if place <= 50 {
					x = RandRange(1, MapSizeX-2)
					y = 1
				} else if place <= 75 {
					x = MapSizeX - 2
					y = RandRange(1, MapSizeY-2)
				} else {
					x = RandRange(1, MapSizeX-2)
					y = MapSizeY - 2
				}
				if IsInFOV(b, (*c)[0].X, (*c)[0].Y, x, y, FOVLength) == false {
					break
				}
			}
			var monster *Creature
			var err error
			for i := 0; i < 100; i++ {
				var temp *Creature
				m := Game.MonstersList[RandInt(len(Game.MonstersList)-1)]
				temp, err = NewCreature(x, y, b, m)
				if Game.TurnCounter >= temp.SpawnTimer {
					monster = temp
				}
			}
			if monster == nil {
				panic("It's not possible to spawn a monster! There is no monster" +
					" with such low SpawnTimer.")
			}
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
			var monsterEq = EquipmentComponent{Objects{w1, w2, wm}, Objects{}}
			monster.EquipmentComponent = monsterEq
			*c = append(*c, monster)
			Game.LivingMonsters++
			Game.WaveCur++
		}
	}
	*/
}
