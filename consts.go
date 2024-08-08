package main

/*Possibile directions*/
const (
	UP    Directions = 1
	DOWN  Directions = 2
	LEFT  Directions = 3
	RIGHT Directions = 4
)

// Map settings
const (
	MAP_LENGTH int = 10
	MAP_WIDTH  int = 10
)

//Map chars
const (
	USER_CHAR    string = "#"
	APPLE_CHAR   string = "♥"
	TERRAIN_CHAR string = " "
)
