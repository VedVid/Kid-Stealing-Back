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
	TileChairI          = 0xE001
	TileChairS          = "U+E001"
	TileCrumblingWallI  = 0xE002
	TileCrumblingWallS  = "U+E002"
	TileDoorsI          = 0xE003
	TileDoorsS          = "U+E003"
	TileFloorI          = 0xE004
	TileFloorS          = "U+E004"
	TileHatchI          = 0xE005
	TileHatchS          = "U+E005"
	TilePillarI         = 0xE006
	TilePillarS         = "U+E006"
	TilePlayerI         = 0xE007
	TilePlayerS         = "U+E007"
	TileTableI          = 0xE008
	TileTableS          = "U+E008"
	TileTreasureHeavyI  = 0xE009
	TileTreasureHeavyS  = "U+E009"
	TileTreasureLightI  = 0xE00A
	TileTreasureLightS  = "U+E00A"
	TileTreasureMediumI = 0xE00B
	TileTreasureMediumS = "U+E00B"
	TileVikingI         = 0xE00C
	TileVikingS         = "U+E00C"
	TileVikingAngryI    = 0xE00D
	TileVikingAngryS    = "U+E00D"
	TileWallI           = 0xE00E
	TileWallS           = "U+E00E"
	TileFootstepsI      = 0xE00F
	TileFootstepsS      = "U+E00F"
	TilePebbleI         = 0xE010
	TilePebbleS         = "U+E010"
)

func SetTiles() {
	_ = SetGlyph("./data/tiles/chair.png", TileChairS)
	_ = SetGlyph("./data/tiles/crumblingwall.png", TileCrumblingWallS)
	_ = SetGlyph("./data/tiles/doors.png", TileDoorsS)
	_ = SetGlyph("./data/tiles/floor.png", TileFloorS)
	_ = SetGlyph("./data/tiles/hatch.png", TileHatchS)
	_ = SetGlyph("./data/tiles/pillar.png", TilePillarS)
	_ = SetGlyph("./data/tiles/player.png", TilePlayerS)
	_ = SetGlyph("./data/tiles/table.png", TileTableS)
	_ = SetGlyph("./data/tiles/treasureheavy.png", TileTreasureHeavyS)
	_ = SetGlyph("./data/tiles/treasurelight.png", TileTreasureLightS)
	_ = SetGlyph("./data/tiles/treasuremedium.png", TileTreasureMediumS)
	_ = SetGlyph("./data/tiles/viking.png", TileVikingS)
	_ = SetGlyph("./data/tiles/vikingangry.png", TileVikingAngryS)
	_ = SetGlyph("./data/tiles/wall.png", TileWallS)
	_ = SetGlyph("./data/tiles/footsteps.png", TileFootstepsS)
	_ = SetGlyph("./data/tiles/pebbles.png", TilePebbleS)
}
