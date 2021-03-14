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


var (
	TileChairI          = 0x11A00
	TileChairS          = "U+11A00"
	TileCrumblingWallI  = 0x11A01
	TileCrumblingWallS  = "U+11A01"
	TileDoorsI          = 0x11A02
	TileDoorsS          = "U+11A02"
	TileFloorI          = 0x11A03
	TileFloorS          = "U+11A03"
	TileHatchI          = 0x11A04
	TileHatchS          = "U+11A04"
	TilePillarI         = 0x11A05
	TilePillarS         = "U+11A05"
	TilePlayerI         = 0x11A06
	TilePlayerS         = "U+11A06"
	TileTableI          = 0x11A07
	TileTableS          = "U+11A07"
	TileTreasureHeavyI  = 0x11A08
	TileTreasureHeavyS  = "U+11A08"
	TileTreasureLightI  = 0x11A09
	TileTreasureLightS  = "U+11A09"
	TileTreasureMediumI = 0x11A0A
	TileTreasureMediumS = "U+11A0A"
	TileVikingI         = 0x11A0B
	TileVikingS         = "U+11A0B"
	TileWallI           = 0x11A0C
	TileWallS           = "U+11A0C"
)

func SetTiles() {
	_ = SetGlyph("./data/tiles/chair.png", TileChairS)
	_ = SetGlyph("./data/tiles/crumblingwall.png", TileCrumblingWallS)
	_ = SetGlyph("./data/tiles/doors.png", TileDoorsS)
	_ = SetGlyph("./data/tiles/hatch.png", TileHatchS)
	_ = SetGlyph("./data/tiles/pillar.png", TilePillarS)
	_ = SetGlyph("./data/tiles/player.png", TilePlayerS)
	_ = SetGlyph("./data/tiles/table.png", TileTableS)
	_ = SetGlyph("./data/tiles/treasureheavy.png", TileTreasureHeavyS)
	_ = SetGlyph("./data/tiles/treasurelight.png", TileTreasureLightS)
	_ = SetGlyph("./data/tiles/treasuremedium.png", TileTreasureMediumS)
	_ = SetGlyph("./data/tiles/viking.png", TileVikingS)
	_ = SetGlyph("./data/tiles/wall.png", TileWallS)
}
