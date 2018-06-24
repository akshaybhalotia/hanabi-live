package main

import (
	"time"
)

var (
	startingTime = 2 * time.Minute  // The amount of time that each player starts with
	timePerTurn  = 20 * time.Second // The amount of extra time a player gets after making a move

	//idleTimeout = time.Minute * 30 // The amount of time that a game is inactive before it is killed by the server
	idleTimeout = time.Minute * 900

	suits = []string{
		"Blue",
		"Green",
		"Yellow",
		"Red",
		"Purple",
		"Black",
		"Rainbow",
		"White",
		"Orange",
	}
	dcSuits = []string{
		"Green",
		"Purple",
		"Navy",
		"Orange",
		"Tan",
		"Burgundy",
	}
	dcClues = []string{
		"Blue",
		"Yellow",
		"Red",
		"Black",
	}
	dcrSuits = []string{
		"Teal",
		"Lime",
		"Orange",
		"Cardinal",
		"Indigo",
		"Rainbow",
	}
	crazySuits = []string{
		"Green",
		"Purple",
		"Orange",
		"White",
		"Rainbow",
		"Black",
	}
	ambiguousSuits = []string{
		"Sky",
		"Navy",
		"Lime",
		"Forest",
		"Tomato",
		"Burgundy",
	}
	ambiguousClues = []string{
		"Blue",
		"Green",
		"Red",
	}
	blueRedSuits = []string{
		"Sky",
		"Berry",
		"Navy",
		"Tomato",
		"Ruby",
		"Mahogany",
	}
	blueRedClues = []string{
		"Blue",
		"Red",
	}
	variants = []string{
		"No Variant",            // 0
		"Orange",                // 1
		"Black (1oE)",           // 2
		"Rainbow",               // 3
		"Dual-color",            // 4
		"Dual & Rainbow",        // 5
		"White & Rainbow",       // 6
		"Wild & Crazy",          // 7
		"Ambiguous",             // 8
		"Blue & Red",            // 9
		"Acid Trip",             // 10
		"Rainbow (1oE)",         // 11
		"Black & Rainbow (1oE)", // 12
	}
)
