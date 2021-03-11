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

type BasicProperties struct {
	/* BasicProperties is struct that aggregates
	   all widely used data, necessary for every
	   map tile and object representation. */
	X, Y      int
	Char      string
	Name      string
	Color     string
	ColorDark string
}

type TreasureProperties struct {
	Treasure     bool
	TreasureChar string
	TreasureCol  string
}

type VisibilityProperties struct {
	/* VisibilityProperties is simple struct
	   for checking if object is always visible,
	   regardless of player's fov, and what
	   is its layer. */
	Layer         int
	AlwaysVisible bool
}

type CollisionProperties struct {
	/* CollisionProperties is struct filled with
	   boolean values, for checking several
	   collision conditions: if cell is blocked,
	   if it blocks creature sight, etc. */
	Blocked     bool
	BlocksSight bool
}

type FighterProperties struct {
	/* FighterProperties stores information about
	   things that can live and fight (ie fighters);
	   it may be used for destructible environment
	   elements as well.
	   AI types are iota (integers) defined
	   in creatures.go. */
	AIType         int
	AITriggered    int
	HPMax          int
	HPCurrent      int
	Attack         int
	Defense        int
	Speed          int
	Stuck          bool
	SpawnTimer     int
	TurnCounter    int
	RangedMaxAmmo  int
	RangedCurAmmo  int
	ThrowablesMax  int
	ThrowablesCur  int
	ChallengeLevel int
	PatrolPoints   [][]int
	NextPoint      int
	LastPosition   []int
	LastSawX       int
	LastSawY       int
	Hidden         bool
	MaxOutOfFOV    int
	OutOfFOV       int
	LightItem1     bool
	LightItem2     bool
	LightItem3     bool
	MediumItem1    bool
	MediumItem2    bool
	HeavyItem1     bool
	StoleAnything  bool
}

type ObjectProperties struct {
	/* Not every Object can be picked up - like tables;
	   also, not every Object can be equipped - like cheese.
	   It's place for other properties - like slot it will
	   occupy, use cases, etc.
	   Note that currently Equippable can not be Consumable,
	   due to removing from Inventory / Equipment problems. */
	DmgMinimal int
	DmgMaximal int
	Pickable   bool
	Equippable bool
	Consumable bool
	Slot       int
	Use        int
}

type EquipmentComponent struct {
	/* EquipmentComponent helps with inventory management.
	   It's part of Creature.
	   Equipment is list of equipped items. It uses
	   const declared in objects.go.
	   Inventory is list of items in backpack. */
	Equipment Objects
	Inventory Objects
}

type GameData struct {
	/* GameData is aggregator of all kind of game
	   variables, like turn counter, last target, etc.
	   In future, it should help creating morgue files
	   by providing info about monsters killed,
	   favorite weapon, etc. */
	// Read monsters's data for future use.
	MonstersList   []string
	TurnCounter    int
	MonstersKilled []string
	KillPoints     int
	TotalHPLost    int
	TotalDMGDealt  int
	// Points take into account:
	// - turns passed (bonus, how long player lived)
	// - number of waves survived
	// - amount of lost HP (malus)
	// - TODO: bonus for killing
	// - TODO: combo kills
	Points     int
	LastTarget *Creature
	// How often monsters will spawn; in turns.
	SpawnRatio     int
	SpawnAmount    int
	WaveNo         int
	WaveMax        int
	WaveCur        int
	LivingMonsters int
	BreakTime      int
	SmallStolen    int
	MediumStolen   int
	HeavyStolen    int
	SpecialPoints  []int
	Breaks         int
	Score          Score
}

type Score struct {
	PlayerName string
	Points     int
}

type Scores struct {
	Entries []Score
}
