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

const (
	// Types of AI.
	NoAI = iota
	PlayerAI
	MeleeDumbAI
	MeleePatherAI
	RangedDumbAI
	RangedPatherAI
	PatrollingAI
)

const (
	MinDistanceBetweenPatrolPoints = 7
	MaxDistanceBetweenPatrolPoints = 12
)

const (
	AINotTriggered = iota
	AITriggeringPhase1
	AITriggeringPhase2
	AITriggered
)

const (
	// Probability of triggering AI
	AITrigger = 85
	OutOfFOVToForgetChance = 60
)

const (
	FootstepsChances = 20
)

func CreaturesTakeTurn(b Board, c *Creatures, o Objects, playerSpeed int) {
	/* Function CreaturesTakeTurn is supposed to handle all enemy creatures
	   actions: movement, attacking, etc.
	   It takes Board and Creatures as arguments.
	   Iterates through all Creatures slice, and calls HandleAI function with
	   specific parameters.
	   It skips NoAI and PlayerAI. */
	var ai int
	for _, v := range *c {
		ai = v.AIType
		if ai == NoAI || ai == PlayerAI {
			continue
		}
/*
		if Game.TurnCounter%SpeedValues[v.Speed] == 0 {
			if v.Speed < SpeedNormal {
				continue
			} else if v.Speed > SpeedNormal {
				HandleAI(b, c, o, v)
				TriggerAI(b, (*c)[0], v)
			}
		}
*/
		HandleAI(b, c, o, v)
		TriggerAI(b, (*c)[0], v)
		v.TurnCounter++
	}
}

func TriggerAI(b Board, p, c *Creature) {
	/* TriggerAI is function that takes Board and two Creatures as arguments.
	   First Creature is supposed to be player, second one - enemy.
	   Enemy with AITriggered set to false will ignore the player existence.
	   AITrigger is probability to notice (and, therefore, switch AITriggered)
	   player if is in monster's FOV. */
	fov := FOVLength
	if b[p.X][p.Y].Hides == true {
		fov = FOVLengthShort
	}
	if IsInFOV(b, c.X, c.Y, p.X, p.Y, FOVLength) == true && RandInt(100) <= AITrigger {
		if b[p.X][p.Y].Hides == false {
			c.LastSawX = p.X
			c.LastSawY = p.Y
			switch c.AITriggered {
			case AINotTriggered:
				c.AITriggered = AITriggeringPhase1
			case AITriggeringPhase1:
				c.AITriggered = AITriggeringPhase2
			case AITriggeringPhase2:
				c.AITriggered = AITriggered
			}
		}
	}
	if c.AITriggered != AITriggered {
		c.Color = "#33a2ac"
		c.ColorDark = "#33a2ac"
		c.Char = "☺"
	} else {
		c.Color = "#d17519"
		c.ColorDark = "#d17519"
		c.Char = "☹"
	}
	if IsInFOV(b, p.X, p.Y, c.X, c.Y, fov) == false &&
		p.DistanceTo(c.X, c.Y) < 10 &&
		RandInt(100) <= FootstepsChances {
			c.Color = "dark red"
			c.ColorDark = "dark red"
			c.Char = "‼"
		}
}

func HandleAI(b Board, cs *Creatures, o Objects, c *Creature) {
	/* HandleAI is robust function that takes Board, Creatures, Objects,
	   and specific Creature as arguments. The most notable argument is
	   the last one - behavior of this entity will be decided in function body.
	   Its behavior will be decided regarding to AIType.
	   This function is very big and *wet*, but it is here to stay, for a while,
	   at least. I thought about code duplication removal by introducing one
	   generic function that would take Creature as argument, and - after
	   AIType check - would use proper HandleMeleeDumbAI (etc.) functions; or
	   would start with available weapons check. (One may want to peek at
	   issue #98 in repo - https://github.com/VedVid/RAWIG/issues/98 ).
	   But, on the other hand, ai has so many variations and edge cases that
	   unifying monster's behavior would result in smaller flexibility. */
	ai := c.AIType
	p := (*cs)[0]
	fov := FOVLength
	if b[p.X][p.Y].Hides == true {
		fov = FOVLengthShort
	}
	switch ai {
	case PatrollingAI:
		switch {
		// 1A - IF NOT AI TRIGGERED
		case c.AITriggered != AITriggered &&
			 p.Hidden == true:
			 // player hidden, continue patrolling
			if c.DistanceTo(c.PatrolPoints[c.NextPoint][0], c.PatrolPoints[c.NextPoint][1]) > 1 {
				c.MoveTowards(b, *cs, c.PatrolPoints[c.NextPoint][0], c.PatrolPoints[c.NextPoint][1], ai)
			} else {
				c.MoveTowards(b, *cs, c.PatrolPoints[c.NextPoint][0], c.PatrolPoints[c.NextPoint][1], ai)
				c.NextPoint++
				if c.NextPoint >= len(c.PatrolPoints) {
					c.NextPoint = 0
				}
			}
		case c.AITriggered != AITriggered &&
			 p.Hidden == false &&
			 IsInFOV(b, c.X, c.Y, p.X, p.Y, fov) == false:
			 // player not in FOV, continue patrolling
			if c.DistanceTo(c.PatrolPoints[c.NextPoint][0], c.PatrolPoints[c.NextPoint][1]) > 1 {
				c.MoveTowards(b, *cs, c.PatrolPoints[c.NextPoint][0], c.PatrolPoints[c.NextPoint][1], ai)
			} else {
				c.MoveTowards(b, *cs, c.PatrolPoints[c.NextPoint][0], c.PatrolPoints[c.NextPoint][1], ai)
				c.NextPoint++
				if c.NextPoint >= len(c.PatrolPoints) {
					c.NextPoint = 0
				}
			}
		case c.AITriggered != AITriggered &&
			 p.Hidden == false &&
			 IsInFOV(b, c.X, c.Y, p.X, p.Y, fov) == true:
			 // enemy didn't notice player yet; continue patrolling
			if c.DistanceTo(c.PatrolPoints[c.NextPoint][0], c.PatrolPoints[c.NextPoint][1]) > 1 {
				c.MoveTowards(b, *cs, c.PatrolPoints[c.NextPoint][0], c.PatrolPoints[c.NextPoint][1], ai)
			} else {
				c.MoveTowards(b, *cs, c.PatrolPoints[c.NextPoint][0], c.PatrolPoints[c.NextPoint][1], ai)
				c.NextPoint++
				if c.NextPoint >= len(c.PatrolPoints) {
					c.NextPoint = 0
				}
			}
		// 1B - IF AI TRIGGERED
		// PLAYER HIDDEN
		case c.AITriggered == AITriggered &&
			 p.Hidden == true &&
			 c.DistanceTo(p.X, p.Y) <= 1:
			 // enemy is standing next to the player
			 // even in player is, in theory, hidden
			 // it's too close to, being alerted,
			 // not notice the player;
			 // therefore: attack!
			if c.X == (*cs)[0].X || c.Y == (*cs)[0].Y {
				c.AttackTarget((*cs)[0], &o, &b, cs, "")
			} else {
				c.MoveTowards(b, *cs, p.X, p.Y, ai)
			}
			c.LastSawX = p.X
			c.LastSawY = p.Y
			c.OutOfFOV = 0
		case c.AITriggered == AITriggered &&
			 p.Hidden == true &&
			 c.DistanceTo(p.X, p.Y) > 1 &&
			 IsInFOV(b, c.X, c.Y, p.X, p.Y, fov) == true:
			 // enemy is alerted, close (but not too close)
			 // to the hidden player;
			 // maybe in the same room?
			c.OutOfFOV = 0
			if c.DistanceTo(p.X, p.Y) < 4 {
				c.LastSawX = c.X
				c.LastSawY = c.Y
				if RandInt(100) < 50 {
					c.AITriggered = AINotTriggered
					// ZERO LAST SAW
				}
			} else {
				c.MoveTowards(b, *cs, p.X, p.Y, ai)
				c.LastSawX = p.X
				c.LastSawY = p.Y
			}
		case c.AITriggered == AITriggered &&
			 p.Hidden == true &&
			 c.DistanceTo(p.X, p.Y) > 1 &&
			 IsInFOV(b, c.X, c.Y, p.X, p.Y, fov) == false &&
			 c.OutOfFOV < c.MaxOutOfFOV:
			 // enemy is alerted,
			 // but far from the hidden player
			 // enemy is still actively searching
			c.OutOfFOV++
			c.MoveTowards(b, *cs, p.X, p.Y, ai)
			c.LastSawX = p.X
			c.LastSawY = p.Y
		case c.AITriggered == AITriggered &&
			 p.Hidden == true &&
			 c.DistanceTo(p.X, p.Y) > 1 &&
			 IsInFOV(b, c.X, c.Y, p.X, p.Y, fov) == false &&
			 c.OutOfFOV >= c.MaxOutOfFOV:
			 // enemy is alerted,
			 // but far from the hidden player
			 // enemy just checks the last place
			 // where the player was seen
			 // then gives up
			c.OutOfFOV++
			if c.X == c.LastSawX && c.Y == c.LastSawY {
				if RandInt(100) < 50 {
					c.AITriggered = AINotTriggered
				}
			} else {
				c.MoveTowards(b, *cs, c.LastSawX, c.LastSawY, ai)
			}
		// PLAYER NOT HIDDEN
		case c.AITriggered == AITriggered &&
			 p.Hidden == false &&
			 c.DistanceTo(p.X, p.Y) <= 1:
			 // player is not hidden, next to the enemy;
			 // therefore, enemy attacks!
			if c.X == (*cs)[0].X || c.Y == (*cs)[0].Y {
				c.AttackTarget((*cs)[0], &o, &b, cs, "")
			} else {
				c.MoveTowards(b, *cs, p.X, p.Y, ai)
			}
			c.LastSawX = p.X
			c.LastSawY = p.Y
			c.OutOfFOV = 0
		case c.AITriggered == AITriggered &&
			 p.Hidden == false &&
			 c.DistanceTo(p.X, p.Y) > 1 &&
			 IsInFOV(b, c.X, c.Y, p.X, p.Y, fov) == true:
			 // enemy is alerted, not too close to the player,
			 // but actually see the player
			c.OutOfFOV = 0
			c.MoveTowards(b, *cs, p.X, p.Y, ai)
			c.LastSawX = p.X
			c.LastSawY = p.Y
		case c.AITriggered == AITriggered &&
			 p.Hidden == false &&
			 c.DistanceTo(p.X, p.Y) > 1 &&
			 IsInFOV(b, c.X, c.Y, p.X, p.Y, fov) == false &&
			 c.OutOfFOV < c.MaxOutOfFOV:
			 // enemy is alerted, player is not hidden, but
			 // not close (not in fov) to the enemy
			 // who is actively searching
			c.OutOfFOV++
			c.MoveTowards(b, *cs, p.X, p.Y, ai)
			c.LastSawX = p.X
			c.LastSawY = p.Y
		case c.AITriggered == AITriggered &&
			 p.Hidden == false &&
			 c.DistanceTo(p.X, p.Y) > 1 &&
			 IsInFOV(b, c.X, c.Y, p.X, p.Y, fov) == false &&
			 c.OutOfFOV >= c.MaxOutOfFOV:
			c.OutOfFOV++
			if c.X == c.LastSawX && c.Y == c.LastSawY {
				if RandInt(100) < 50 {
					c.AITriggered = AINotTriggered
				}
			} else {
				c.MoveTowards(b, *cs, c.LastSawX, c.LastSawY, ai)
			}
		default:
			fmt.Println("AI corner case that has not been considered; continuing patrolling.")
			if c.DistanceTo(c.PatrolPoints[c.NextPoint][0], c.PatrolPoints[c.NextPoint][1]) > 1 {
				c.MoveTowards(b, *cs, c.PatrolPoints[c.NextPoint][0], c.PatrolPoints[c.NextPoint][1], ai)
			} else {
				c.MoveTowards(b, *cs, c.PatrolPoints[c.NextPoint][0], c.PatrolPoints[c.NextPoint][1], ai)
				c.NextPoint++
				if c.NextPoint >= len(c.PatrolPoints) {
					c.NextPoint = 0
				}
			}
		}
	}
}
