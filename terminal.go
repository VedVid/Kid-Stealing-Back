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
	"runtime"
	"strconv"

	blt "bearlibterminal"
)

const (
	// Setting BearLibTerminal window.
	// What's significance of "6"?
	MapSizeX     = 25
	MapSizeY     = 25
	MapPositionX = 0
	MapPositionY = 0
	WindowSizeX  = (MapSizeX * 6) + (20 * 3)
	WindowSizeY  = (MapSizeY + 5 + 1) * 6
	UIPosX       = MapSizeX * 6
	UIPosY       = 0
	UISizeX      = WindowSizeX - (MapSizeX * 6)
	UISizeY      = WindowSizeY
	LogSizeX     = WindowSizeX
	LogSizeY     = (WindowSizeY / 6) - MapSizeY
	LogPosX      = 0
	LogPosY      = MapSizeY * 6

	GameTitle   = "About The Kid Who Stole The Relics Back"
	GameVersion = "(0.1 20210314)"
)

var (
	ReferenceFontName = "Deferral-Square.ttf"
	ReferenceFontSize = 2
	UIFontName        = "NovaMono-Regular.ttf"
	UIFontSize        = 12
	UIFontSpacingX    = UIFontSize / ReferenceFontSize / 2
	UIFontSpacingY    = UIFontSize / ReferenceFontSize
	UIFontSpacing     = strconv.Itoa(UIFontSpacingX) + "x" +
		strconv.Itoa(UIFontSpacingY)
	GameFontName      = "Deferral-Square.ttf"
	GameFontSize      = 12
	GameFontSpacingX  = UIFontSize / ReferenceFontSize
	GameFontSpacingY  = UIFontSize / ReferenceFontSize
	GameFontSpacing   = strconv.Itoa(GameFontSpacingX) + "x" +
		strconv.Itoa(GameFontSpacingY)
	TitleFontName     = "NovaMono-Regular.ttf"
	TitleFontSize     = 36
	TitleFontSpacingX = TitleFontSize / ReferenceFontSize / 2
	TitleFontSpacingY = TitleFontSize / ReferenceFontSize
	TitleFontSpacing  = strconv.Itoa(TitleFontSpacingX) + "x" +
		strconv.Itoa(TitleFontSpacingY)
	TileSizeX         = 18
	TileSizeY         = 18
)

func constrainThreads() {
	/* Constraining processor and threads is necessary,
	   because BearLibTerminal often crashes otherwise. */
	runtime.GOMAXPROCS(1)
	runtime.LockOSThread()
}

func InitializeBLT() {
	/* Constraining threads and setting BearLibTerminal window. */
	constrainThreads()
	blt.Open()
	sizeX, sizeY := strconv.Itoa(WindowSizeX), strconv.Itoa(WindowSizeY)
	window := "window: size=" + sizeX + "x" + sizeY
	currentVersion := GameVersion
	if Tiles {
		currentVersion += " [[Tiles]]"
	}
	blt.Set(window + ", title=' " + GameTitle + " " + GameVersion +
		"'; font: " + ReferenceFontName + ", size=" +
		strconv.Itoa(ReferenceFontSize))
	uiFont := "ui font: " + UIFontName + ", size=" +
		strconv.Itoa(UIFontSize) + ", spacing=" + UIFontSpacing
	blt.Set(uiFont)
	gameFont := "game font: " + GameFontName + ", size=" +
		strconv.Itoa(GameFontSize) + ", spacing=" + GameFontSpacing
	blt.Set(gameFont)
	titleFont := "title font: " + TitleFontName + ", size=" +
		strconv.Itoa(TitleFontSize) + ", spacing=" + TitleFontSpacing
	blt.Set(titleFont)
	blt.Clear()
	blt.Refresh()
}
