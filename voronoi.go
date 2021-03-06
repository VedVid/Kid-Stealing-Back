/*
Copyright (c) 2021, Tomasz "VedVid" Nowakowski
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

type VoronoiTile struct {
	X   int
	Y   int
	Val int
}

type VoronoiMap []*VoronoiTile

func NewVoronoi(numCells int, topic Topic) VoronoiMap {
	var vm = VoronoiMap{}
	var nx = []int{}
	var ny = []int{}
	var nv = []int{}
	for i := 0; i < (numCells); i++ {
		nx = append(nx, rand.Intn(MapSizeX)-1)
		ny = append(ny, rand.Intn(MapSizeY)-1)
		nv = append(nv, rand.Intn(len(topic.Elements)))
	}
	for y := 0; y < MapSizeY; y++ {
		for x := 0; x < MapSizeX; x++ {
			//dMin := EuclideanDistance(MapSizeX-1, MapSizeY-1)
			dMin := ManhattanDistance(MapSizeX-1, MapSizeY-1)
			j := -1
			for i := 0; i < (numCells); i++ {
				//d := EuclideanDistance(nx[i]-x, ny[i]-y)
				d := ManhattanDistance(nx[i]-x, ny[i]-y)
				if d < dMin {
					dMin = d
					j = i
				}
			}
			vt := VoronoiTile{x, y, nv[j]}
			vm = append(vm, &vt)
		}
	}
	return vm
}
