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
	"math"
	"strconv"

	blt "bearlibterminal"
)

const (
	// Values that are important for creating and backtracking graph.
	nodeInitialWeight = -1 // Nodes not traversed.
)

type Node struct {
	/* Node is struct that mimics some properties
	   of Tile struct (implemented in map.go).
	   X, Y are coords of Node, and Weight is value important
	   for graph creating, and later - finding shortest path
	   from source (creature) to goal (coords).
	   Weight is supposed to be set to -1 initially - it
	   marks Node as not traversed during
	   graph creation process. */
	X, Y   int
	Weight int
}

func TilesToNodes() [][]*Node {
	/* TilesToNodes is function that takes Board
	   (ie map, or fragment, of level) as argument. It converts
	   Tiles to Nodes, and returns 2d array of *Node to mimic
	   Board behaviour.
	   In future, it may be worth to create new type, ie
	   type Nodes [][]*Node.
	   During initialization, every newly created Node has
	   its Weight set to -1 to mark that it's not traversed. */
	nodes := make([][]*Node, MapSizeX)
	for i := range nodes {
		nodes[i] = make([]*Node, MapSizeY)
	}
	for x := 0; x < MapSizeX; x++ {
		for y := 0; y < MapSizeY; y++ {
			nodes[x][y] = &Node{x, y, nodeInitialWeight}
		}
	}
	return nodes
}

func FindAdjacent(b Board, c *Creature, cs Creatures, nodes [][]*Node, frontiers []*Node,
	start *Node, w int) ([]*Node, bool) {
	/* Function FindAdjacent takes Board, Board-like [][]*Node array,
	   coords of starting point, and current value to attribute Weight field
	   of Node; FindAdjacent returns slice of adjacent tiles and startFound
	   bool flag.
	   At start, empty slice of *Node is created, and boolean flag startFound
	   is set to false; this flag will be set to true, if function will find
	   node that is source of path, and it'll break the loops.
	   Primary for loop uses one of frontiers, and x, y nested loops
	   checks for its adjacent tiles (more details in in-line comments);
	   if tile met conditions, its Weight is set to current w value, then node
	   is added to list of adjacent tiles. */
	var adjacent = []*Node{}
	startFound := false
	for i := 0; i < len(frontiers); i++ {
		for x := frontiers[i].X - 1; x <= frontiers[i].X+1; x++ {
			for y := frontiers[i].Y - 1; y <= frontiers[i].Y+1; y++ {
				if x == start.X && y == start.Y &&
					(x == frontiers[i].X || y == frontiers[i].Y) {
					startFound = true
					nodes[x][y].Weight = w
					goto End
				}
				if x < 0 || x >= MapSizeX || y < 0 || y >= MapSizeY {
					continue //node is out of map bounds
				}
				if nodes[x][y].Weight != nodeInitialWeight {
					continue //node is marked as traversed already
				}
				if x == frontiers[i].X && y == frontiers[i].Y {
					continue //it's the current frontier node
				}
				if b[x][y].Blocked == true {
					continue //tile is blocked
				}
				if c.AITriggered != AITriggered {
					if b[x][y].Char != "." && b[x][y].Char != "+" {
						if b[x][y].Treasure == false {
							continue
						}
					}
				}
				if GetAliveCreatureFromTile(x, y, cs) != nil {
					continue //tile is occupied by other monster
				}
				if x == frontiers[i].X || y == frontiers[i].Y {
					nodes[x][y].Weight = w
					adjacent = append(adjacent, nodes[x][y])
				}
			}
		}
	}
End:
	return adjacent, startFound
}

func (c *Creature) MoveTowardsPath(b Board, cs Creatures, tx, ty int) error {
	/* MoveTowardsPath is one of main pathfinding methods. It takes
	   Board and ints tx, ty (ie target coords) as arguments.
	   MoveTowardsPath uses weighted graph to find shortest path
	   from goal (tx, ty - it's more universal than passing Node or
	   Creature) to source (creature, ie receiver).
	   At first, it creates simple graph with all nodes' Weight set to
	   -1 as not-yet-traversed. Later, it starts potentially infinite loop
	   that breaks if starting position is found by FindAdjacent function,
	   or when FindAdjacent won't find any proper tiles that are
	   adjacent to previously found ones (ie frontiers).
	   After every iteration, local variable "w" used to attribute
	   node Weight increases by one, to mark that it's another step away
	   from goal position; it makes backtracking easy - Creature position
	   is end of path / graph, so Creature has only find node with
	   Weight set to lesser value that node occupied by Creature.
	   Effect may be a bit strange as it takes first node that met
	   conditions, but works rather well with basic MoveTowards method. */
	nodes := TilesToNodes()
	start := nodes[c.X][c.Y]
	startFound := false
	goal := nodes[tx][ty]
	goal.Weight = 0
	var frontiers = []*Node{goal}
	w := 0
	for {
		w++
		if len(frontiers) == 0 || startFound == true {
			break
		}
		frontiers, startFound = FindAdjacent(b, c, cs, nodes, frontiers, start, w)
	}
	//RenderWeights was here
	if ShowMapValidationProcess != ShowMapValidationProcessOff {
		blt.Clear()
		CastRays(b, c.X, c.Y)
		PrintBoard(b, Creatures{c})
		PrintCreatures(b, Creatures{c})
		blt.Refresh()
	}
	if ShowMapValidationProcess == ShowMapValidationProcessOnSlow {
		_ = ReadInput()
	}
	dx, dy, err := BacktrackPath(nodes, start)
	c.Move(dx, dy, b, cs)
	return err
}

func BacktrackPath(nodes [][]*Node, start *Node) (int, int, error) {
	/* Function BacktrackPath takes 2d array of *Node, and
	   starting *Node as arguments; it returns two ints, that serves
	   as coords.
	   BacktrackPath is used in pathfinding. It uses weighted graph
	   that has some sort of path already created (more in comments for
	   MoveTowardsPath and FindAdjacent). Instead of creating
	   proper path, or using search algorithm, structure of graph
	   allows to use just node with smaller Weight than start node.
	   It returns error if can't find proper tile.
	   Note: returning three values at once is ugly. */
	direction := *start
	for x := start.X - 1; x <= start.X+1; x++ {
		for y := start.Y - 1; y <= start.Y+1; y++ {
			if x < 0 || x >= MapSizeX || y < 0 || y >= MapSizeY {
				continue // Node is out of map bounds.
			}
			if x == start.X && y == start.Y {
				continue // This node is the current node.
			}
			if nodes[x][y].Weight == nodeInitialWeight {
				continue // Node is not part of path.
			}
			if nodes[x][y].Weight == direction.Weight-1 {
				direction = *nodes[x][y] // Node is closer to goal than current node.
				goto EndLoop
			}
			continue
		}
	}
EndLoop:
	var err error
	if direction == *start {
		// This error doesn't need helper function from err_.go.
		err = errors.New("Warning: function BacktrackPath could not find direction that met all requirements." +
			"\n    Returned coords are coords of starting position.")
	}
	dx := direction.X - start.X
	dy := direction.Y - start.Y
	return dx, dy, err
}

func RenderWeights(nodes [][]*Node) {
	/* RenderWeights is created for debugging purposes.
	   Clears whole map, and prints Weights of all nodes
	   of graph, then waits for user input to continue
	   game loop.
	   It's supposed to be called near the end of
	   MoveTowardsPath method. */
	blt.Clear()
	for x := 0; x < MapSizeX; x++ {
		for y := 0; y < MapSizeY; y++ {
			glyph := strconv.Itoa(nodes[x][y].Weight)
			if nodes[x][y].Weight == nodeInitialWeight {
				glyph = "-"
			} else if nodes[x][y].Weight > 9 {
				glyph = "+"
			}
			blt.Print(x, y, glyph)
		}
	}
	blt.Refresh()
	blt.Read()
}

func (c *Creature) MoveTowards(b Board, cs Creatures, tx, ty int, ai int) {
	/* MoveTowards is *the* main method for pathfinding.
	   Has *Creature as receiver, and takes Board (ie map of level),
	   ints tx and ty (ie coords of Node - in that case, it's more
	   universal than passing whole Node or Creature), and ai - it's
	   style of ai (these style markers are enums declared in ai.go)
	   as arguments.
	   The default implementation calls MoveTowardsPath for "smart"
	   monsters and uses surprisingly complex algorithm for the
	   "dumb" ones.
	   In this project, however, I assume that every monster is kind
	   of "smart" so right now it will just utilize MoveTowardsPath
	   for all kind of monsters. */
	err := c.MoveTowardsPath(b, cs, tx, ty)
	if err != nil {
		if RandInt(100) < 75 {
			x, y := 0, 0
			if RandInt(100) < 50 {
				if RandInt(100) < 50 {
					y = -1
				} else {
					y = 1
				}
			} else {
				if RandInt(100) < 50 {
					x = -1
				} else {
					x = 1
				}
			}
			c.Move(x, y, b, cs)
		}
	}
}

func (c *Creature) DistanceTo(tx, ty int) int {
	/* DistanceTo is Creature method. It takes target x and target y as args.
	   Computes, then returns, distance from receiver to target. */
	dx := float64(tx - c.X)
	dy := float64(ty - c.Y)
	return RoundFloatToInt(math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2)))
}
