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

type Topic struct {
	SelfName string
	Elements []Biome
}

type Biome struct {
	SelfName string
	Elements []BiomeElement
}

type BiomeElement struct {
	SelfName      string
	Chances       int
	Char          []string
	Name          []string
	Color         []string
	ColorDark     string
	Layer         int
	AlwaysVisible bool
	Explored      bool
	Slows         bool
	Blocked       bool
	BlocksSight   bool
}

var Topics = []Topic{Desert, Forest, Mountains}

var (
	Desert = Topic{"Desert", []Biome{DesertErg, DesertErg, DesertOasis}}

	DesertErg = Biome{"DesertErg", []BiomeElement{DesertErgSand}}

	DesertErgSand = BiomeElement{
		"DesertErgSand",
		100,
		[]string{"."},
		[]string{"sand"},
		[]string{"#C2B280"},
		"grey",
		2,
		true,
		true,
		false,
		false,
		false,
	}

	DesertOasis = Biome{"DesertOasis", []BiomeElement{DesertOasisShallowWater, DesertOasisSand,
		DesertOasisTree, DesertOasisGrass, DesertOasisBush}}

	DesertOasisSand = BiomeElement{
		"DesertOasisSand",
		20,
		[]string{"."},
		[]string{"sand"},
		[]string{"#C2B280"},
		"grey",
		2,
		true,
		true,
		false,
		false,
		false,
	}

	DesertOasisShallowWater = BiomeElement{
		"DesertOasisShallowWater",
		50,
		[]string{"~"},
		[]string{"shallow water"},
		[]string{"#ADD8E6"},
		"grey",
		2,
		true,
		true,
		false,
		false,
		false,
	}

	DesertOasisTree = BiomeElement{
		"DesertOasisTree",
		75,
		[]string{"T"},
		[]string{"tree"},
		[]string{"#568203"},
		"grey",
		2,
		true,
		true,
		false,
		true,
		true,
	}

	DesertOasisGrass = BiomeElement{
		"DesertOasisGrass",
		90,
		[]string{","},
		[]string{"grass"},
		[]string{"#A9BA9D"},
		"grey",
		2,
		true,
		true,
		false,
		false,
		false,
	}

	DesertOasisBush = BiomeElement{
		"DesertOasisBush",
		100,
		[]string{":"},
		[]string{"bush"},
		[]string{"#4B6F44"},
		"grey",
		2,
		true,
		true,
		true,
		false,
		false,
	}
)

var (
	Forest = Topic{"Forest", []Biome{ForestWoods, ForestYoungForest, ForestClearing}}

	ForestWoods = Biome{"ForestWoods", []BiomeElement{ForestWoodsTree, ForestWoodsGrass,
		ForestWoodsSand, ForestWoodsBush}}

	ForestWoodsTree = BiomeElement{
		"ForestWoodsTree",
		40,
		[]string{"T"},
		[]string{"tree"},
		[]string{"#4B6F44"},
		"grey",
		2,
		true,
		true,
		false,
		true,
		true,
	}

	ForestWoodsGrass = BiomeElement{
		"ForestWoodsGrass",
		80,
		[]string{","},
		[]string{"grass"},
		[]string{"#4F7942"},
		"grey",
		2,
		true,
		true,
		false,
		false,
		false,
	}

	ForestWoodsSand = BiomeElement{
		"ForestWoodsSand",
		85,
		[]string{"."},
		[]string{"sand"},
		[]string{"#C2B280"},
		"grey",
		2,
		true,
		true,
		false,
		false,
		false,
	}

	ForestWoodsBush = BiomeElement{
		"ForestWoodsBush",
		100,
		[]string{":"},
		[]string{"bush"},
		[]string{"#4B6F44"},
		"grey",
		2,
		true,
		true,
		true,
		false,
		false,
	}

	ForestYoungForest = Biome{"ForestYoungForest", []BiomeElement{ForestYoungForestTree,
		ForestYoungForestGrass, ForestYoungForestSand, ForestYoungForestBush,
		ForestYoungForestYoungTree}}

	ForestYoungForestTree = BiomeElement{
		"ForestYoungForestTree",
		15,
		[]string{"T"},
		[]string{"tree"},
		[]string{"#4B6F44"},
		"grey",
		2,
		true,
		true,
		false,
		true,
		true,
	}

	ForestYoungForestGrass = BiomeElement{
		"ForestYoungForestGrass",
		45,
		[]string{","},
		[]string{"grass"},
		[]string{"#4F7942"},
		"grey",
		2,
		true,
		true,
		false,
		false,
		false,
	}

	ForestYoungForestSand = BiomeElement{
		"ForestYoungForestSand",
		50,
		[]string{"."},
		[]string{"sand"},
		[]string{"#C2B280"},
		"grey",
		2,
		true,
		true,
		false,
		false,
		false,
	}

	ForestYoungForestBush = BiomeElement{
		"ForestYoungForestBush",
		60,
		[]string{":"},
		[]string{"bush"},
		[]string{"#4B6F44"},
		"grey",
		2,
		true,
		true,
		true,
		false,
		false,
	}

	ForestYoungForestYoungTree = BiomeElement{
		"ForestYoungForestYoungTree",
		100,
		[]string{"^"},
		[]string{"young tree"},
		[]string{"#5C8055"},
		"grey",
		2,
		true,
		true,
		true,
		false,
		false,
	}

	ForestClearing = Biome{"ForestClearing", []BiomeElement{ForestClearingTree,
		ForestClearingGrass, ForestClearingSand, ForestClearingBush, ForestClearingYoungTree}}

	ForestClearingTree = BiomeElement{
		"ForestClearingTree",
		5,
		[]string{"T"},
		[]string{"tree"},
		[]string{"#4B6F44"},
		"grey",
		2,
		true,
		true,
		false,
		true,
		true,
	}

	ForestClearingGrass = BiomeElement{
		"ForestClearingGrass",
		80,
		[]string{","},
		[]string{"grass"},
		[]string{"#4F7942"},
		"grey",
		2,
		true,
		true,
		false,
		false,
		false,
	}

	ForestClearingSand = BiomeElement{
		"ForestClearingSand",
		85,
		[]string{"."},
		[]string{"sand"},
		[]string{"#C2B280"},
		"grey",
		2,
		true,
		true,
		false,
		false,
		false,
	}

	ForestClearingBush = BiomeElement{
		"ForestClearingBush",
		93,
		[]string{":"},
		[]string{"bush"},
		[]string{"#4B6F44"},
		"grey",
		2,
		true,
		true,
		true,
		false,
		false,
	}

	ForestClearingYoungTree = BiomeElement{
		"ForestClearingYoungTree",
		100,
		[]string{"^"},
		[]string{"young tree"},
		[]string{"#5C8055"},
		"grey",
		2,
		true,
		true,
		true,
		false,
		false,
	}
)

var (
	Mountains = Topic{"Mountains", []Biome{MountainsHighMountains, MountainsMeadows, MountainsForest}}

	MountainsHighMountains = Biome{"MountainsHighMountains", []BiomeElement{MountainsHighMountainsMountains}}

	MountainsHighMountainsMountains = BiomeElement{
		"MountainsHighMountainsMountains",
		100,
		[]string{"^"},
		[]string{"mountains"},
		[]string{"#DCDCDC"},
		"grey",
		2,
		true,
		true,
		false,
		true,
		true,
	}

	MountainsMeadows = Biome{"MountainsMeadows", []BiomeElement{MountainsMeadowsGrass, MountainsMeadowsBush,
		MountainsMeadowsTree}}

	MountainsMeadowsGrass = BiomeElement{
		"MountainsMeadowsGrass",
		85,
		[]string{","},
		[]string{"grass"},
		[]string{"#8A9A5B"},
		"grey",
		2,
		true,
		true,
		false,
		false,
		false,
	}

	MountainsMeadowsBush = BiomeElement{
		"MountainsMeadowsBush",
		95,
		[]string{":"},
		[]string{"bush"},
		[]string{"#355E3B"},
		"grey",
		2,
		true,
		true,
		true,
		false,
		false,
	}

	MountainsMeadowsTree = BiomeElement{
		"MountainsMeadowsTree",
		100,
		[]string{"T"},
		[]string{"tree"},
		[]string{"#138808"},
		"grey",
		2,
		true,
		true,
		false,
		true,
		true,
	}

	MountainsForest = Biome{"MountainsForest", []BiomeElement{MountainsForestGrass, MountainsForestBush,
		MountainsForestTree}}

	MountainsForestGrass = BiomeElement{
		"MountainsForestGrass",
		45,
		[]string{","},
		[]string{"grass"},
		[]string{"#228B22"},
		"grey",
		2,
		true,
		true,
		false,
		false,
		false,
	}

	MountainsForestBush = BiomeElement{
		"MountainsForestBush",
		65,
		[]string{":"},
		[]string{"bush"},
		[]string{"#00693E"},
		"grey",
		2,
		true,
		true,
		true,
		false,
		false,
	}

	MountainsForestTree = BiomeElement{
		"MountainsForestTree",
		100,
		[]string{"T"},
		[]string{"tree"},
		[]string{"#009000"},
		"grey",
		2,
		true,
		true,
		false,
		true,
		true,
	}
)
