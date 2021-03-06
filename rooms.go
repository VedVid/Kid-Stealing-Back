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

/*
LEGEND:
    . - floor
    o - pillar
    h - chair
    T - table
*/

var (
	roomFurniture1 = [][]string{
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
	}

	roomFurniture2 = [][]string{
		[]string{".", ".", ".", ".", "."},
		[]string{".", "o", ".", "o", "."},
		[]string{".", ".", ".", ".", "."},
		[]string{".", "o", ".", "o", "."},
		[]string{".", ".", ".", ".", "."},
	}

	roomFurniture3 = [][]string{
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", "h", ".", "."},
		[]string{".", "h", "T", "h", "."},
		[]string{".", ".", "h", ".", "."},
		[]string{".", ".", ".", ".", "."},
	}

	roomFurniture4 = [][]string{
		[]string{".", ".", ".", ".", "."},
		[]string{".", "h", "T", "h", "."},
		[]string{".", "h", "T", "h", "."},
		[]string{".", "h", "T", "h", "."},
		[]string{".", ".", ".", ".", "."},
	}

	roomFurniture5 = [][]string{
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", "o", ".", "."},
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
	}

	roomFurniture6 = [][]string{
		[]string{"h", ".", ".", ".", "h"},
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
		[]string{"h", ".", ".", ".", "h"},
	}

	roomFurniture7 = [][]string{
		[]string{"o", ".", ".", ".", "o"},
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
		[]string{"o", ".", ".", ".", "o"},
	}

	roomFurniture8 = [][]string{
		[]string{"o", ".", ".", ".", "o"},
		[]string{".", ".", "h", ".", "."},
		[]string{".", "h", "T", "h", "."},
		[]string{".", ".", "h", ".", "."},
		[]string{"o", ".", ".", ".", "o"},
	}
)

var HardcodedRooms = [][][]string{
	roomFurniture1, roomFurniture2, roomFurniture3, roomFurniture4,
	roomFurniture5, roomFurniture6, roomFurniture7, roomFurniture8,
}
