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
	TileChairI               = 0xE001
	TileChairS               = "U+E001"
	TileCrumblingWallI       = 0xE002
	TileCrumblingWallS       = "U+E002"
	TileDoorsI               = 0xE003
	TileDoorsS               = "U+E003"
	TileFloorI               = 0xE004
	TileFloorS               = "U+E004"
	TileHatchI               = 0xE005
	TileHatchS               = "U+E005"
	TilePillarI              = 0xE006
	TilePillarS              = "U+E006"
	TilePlayerI              = 0xE007
	TilePlayerS              = "U+E007"
	TileTableI               = 0xE008
	TileTableS               = "U+E008"
	TileTreasureHeavyI       = 0xE009
	TileTreasureHeavyS       = "U+E009"
	TileTreasureLightI       = 0xE00A
	TileTreasureLightS       = "U+E00A"
	TileTreasureMediumI      = 0xE00B
	TileTreasureMediumS      = "U+E00B"
	TileVikingI              = 0xE00C
	TileVikingS              = "U+E00C"
	TileVikingAngryI         = 0xE00D
	TileVikingAngryS         = "U+E00D"
	TileWallI                = 0xE00E
	TileWallS                = "U+E00E"
	TileFootstepsI           = 0xE00F
	TileFootstepsS           = "U+E00F"
	TilePebbleI              = 0xE010
	TilePebbleS              = "U+E010"
	TileLookingI             = 0xE011
	TileLookingS             = "U+E011"
	TileCrossI               = 0xE012
	TileCrossS               = "U+E012"
	TileHeartI               = 0xE013
	TileHeartS               = "U+E013"
	TileHeartEmptyI          = 0xE014
	TileHeartEmptyS          = "U+E014"
	TileStoneI               = 0xE015
	TileStoneS               = "U+E015"
	TileStoneEmptyI          = 0xE016
	TileStoneEmptyS          = "U+E016"
	TileTreasureHeavyEmptyI  = 0xE017
	TileTreasureHeavyEmptyS  = "U+E017"
	TileTreasureLightEmptyI  = 0xE018
	TileTreasureLightEmptyS  = "U+E018"
	TileTreasureMediumEmptyI = 0xE019
	TileTreasureMediumEmptyS = "U+E019"
	TileEncNoneI             = 0x1A00
	TileEncNoneS             = "U+1A00"
	TileEncLightI            = 0x1A01
	TileEncLightS            = "U+1A01"
	TileEncMediumI           = 0x1A02
	TileEncMediumS           = "U+1A02"
	TileEncHeavyI            = 0x1A03
	TileEncHeavyS            = "U+1A03"
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
	_ = SetGlyph("./data/tiles/looking.png", TileLookingS)
	_ = SetGlyph("./data/tiles/cross.png", TileCrossS)
	_ = SetGlyph("./data/tiles/heart.png", TileHeartS)
	_ = SetGlyph("./data/tiles/heartempty.png", TileHeartEmptyS)
	_ = SetGlyph("./data/tiles/stone.png", TileStoneS)
	_ = SetGlyph("./data/tiles/stoneempty.png", TileStoneEmptyS)
	_ = SetGlyph("./data/tiles/treasureheavyempty.png", TileTreasureHeavyEmptyS)
	_ = SetGlyph("./data/tiles/treasuremediumempty.png", TileTreasureMediumEmptyS)
	_ = SetGlyph("./data/tiles/treasurelightempty.png", TileTreasureLightEmptyS)
	_ = SetGlyph("./data/tiles/encn.png", TileEncNoneS)
	_ = SetGlyph("./data/tiles/encl.png", TileEncLightS)
	_ = SetGlyph("./data/tiles/encm.png", TileEncMediumS)
	_ = SetGlyph("./data/tiles/ench.png", TileEncHeavyS)
}
