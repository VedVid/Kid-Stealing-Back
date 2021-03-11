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

import "sort"
import "time"

const (
	AllTreasuresStolen = iota
	Dieded
	StoleNothing
)

var specialPoints = map[int]int{
	AllTreasuresStolen: 500,
	Dieded: -500,
	StoleNothing: -250,
}


func UpdateScores() {
	var tempScore = Score{time.Now().Format("20060102T1504"), Game.Points}
	HighScores.Entries = append(HighScores.Entries, tempScore)
	sort.Slice(HighScores.Entries, func(i, j int) bool {
		return HighScores.Entries[i].Points > HighScores.Entries[j].Points
	})
}

func (g GameData) CalculatePoints() int {
	points := 0
	points -= g.TurnCounter
	points -= g.TotalHPLost * 10
	points += g.SmallStolen * 100
	points += g.MediumStolen * 250
	points += g.HeavyStolen * 500
	for _, v := range g.SpecialPoints {
		points += specialPoints[v]
	}
	points -= g.Breaks * 100
	return points
}
