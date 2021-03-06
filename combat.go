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
	"fmt"
	"strings"
)

const (
	// Attack types vs equipment slots
	MeleeSlot = iota
	RangedSlot
	ThrowableSlot
)

func (c *Creature) AttackTarget(t *Creature, o *Objects, b *Board, cs *Creatures, attack string) {
	/* Method Attack handles damage rolls for combat. Receiver "c" is attacker,
	   argument "t" is target. Including o *Objects is necessary for dropping
	   loot by dead enemies.
	   Critical hit is if attack roll is the same as receiver
	   attack attribute.
	   Result of attack is displayed in combat log, but messages need more polish. */
	playerAttacks := true
	if c != (*cs)[0] {
		playerAttacks = false
	}
	att := RandInt(c.Attack) //basic attack roll
	if playerAttacks == true {
		switch attack {
		case StrMelee:
			att = RandRange(c.Equipment[MeleeSlot].DmgMinimal,
							c.Equipment[MeleeSlot].DmgMaximal)
		case StrRanged:
			att = RandRange(c.Equipment[RangedSlot].DmgMinimal,
							c.Equipment[RangedSlot].DmgMaximal)
		case StrThrowable:
			att = RandRange(c.Equipment[ThrowableSlot].DmgMinimal,
							c.Equipment[ThrowableSlot].DmgMaximal)
		default:
			att = RandInt(c.Attack)
			fmt.Println("Something went wrong. Player's attack was not found in: " +
				"[StrMelee, StrRanged, StrThrowable].")
		}
	}
	att2 := 0                //critical bonus
	def := t.Defense         //opponent's defense
	dmg := 0                 //dmg delivered
	crit := false            //was it critical hit?
	if RandInt(100) >= 90 {  //critical hit!
		crit = true
		att = c.Attack
		att2 = RandInt(c.Attack)
		if playerAttacks == true {
			switch attack {
			case StrMelee:
				att = c.Equipment[MeleeSlot].DmgMaximal
				att2 = RandRange(c.Equipment[MeleeSlot].DmgMinimal,
							c.Equipment[MeleeSlot].DmgMaximal)
			case StrRanged:
				att = c.Equipment[RangedSlot].DmgMaximal
				att2 = RandRange(c.Equipment[RangedSlot].DmgMinimal,
							c.Equipment[RangedSlot].DmgMaximal)
			case StrThrowable:
				att = c.Equipment[ThrowableSlot].DmgMaximal
				att2 = RandRange(c.Equipment[ThrowableSlot].DmgMinimal,
							c.Equipment[ThrowableSlot].DmgMaximal)
			default:
				att = c.Attack
				att2 = RandInt(c.Attack)
				fmt.Println("Something went wrong. Player's attack was not found in: " +
					"[StrMelee, StrRanged, StrThrowable].")
			}
		}
	}
	switch {
	case att < def: // Attack score if lower than target defense.
		if crit == false {
			if playerAttacks {
				AddMessage(strings.Title(t.Name) + " deflects your attack.")
			}
		} else {
			dmg = att2 // Critical hit, but against heavily armored enemy.
		}
	case att == def: // Attack score is equal to target defense.
		if crit == false {
			dmg = 1 // It's just a scratch...
		} else {
			dmg = att
		}
	case att > def: // Attack score is bigger than target defense.
		if crit == false {
			dmg = att
		} else {
			dmg = att + att2 // Critical attack!
			if playerAttacks {
				AddMessage("[color=dark green]You landed a critical attack!")
			} else {
				AddMessage("[color=dark red]You got critically hit by " + c.Name + ".")
			}
		}
	}
	if t == (*cs)[0] {
		// Player is hit
		Game.TotalHPLost += dmg
	} else if c == (*cs)[0] {
		// Player hits
		Game.TotalDMGDealt += dmg
	}
	t.TakeDamage(dmg, o, b, cs)
}

func (c *Creature) TakeDamage(dmg int, o *Objects, b *Board, cs *Creatures) {
	/* Method TakeDamage has *Creature as receiver and takes damage integer
	   as argument. dmg value is deducted from Creature current HP.
	   If HPCurrent is below zero after taking damage, Creature dies.
	   o as map objects is passed with Die to handle dropping loot. */
	c.HPCurrent -= dmg
	if c.HPCurrent <= 0 {
		c.Die(o, b, cs)
		if c != (*cs)[0] {
			AddMessage("[color=dark grey]" + strings.Title(c.Name) + " dies.")
		}
	}
}
