package main

import (
	"math/rand/v2"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}
func randomCoordinates(from int, to int) [2]int {
	return [2]int{randRange(from, to), randRange(from, to)}
}
func isAppleInSnake(snake [][2]int, apple [2]int) bool {
	for i := 0; i < len(snake); i++ {
		if snake[i] == apple {
			return true
		}
	}
	return false
}

func renderMap(terrain [MAP_WIDTH][MAP_LENGTH]string, snake [][2]int, apple [2]int) [MAP_WIDTH][MAP_LENGTH]string {
	for i := 0; i < MAP_WIDTH; i++ {
		for j := 0; j < MAP_LENGTH; j++ {
			if i == 0 && j == 0 {
				terrain[i][j] = "┌"
				continue
			}
			if i == 0 && j == MAP_WIDTH-1 {
				terrain[i][j] = "┐"
				continue
			}
			if i == MAP_LENGTH-1 && j == 0 {
				terrain[i][j] = "└"
				continue
			}
			if i == MAP_LENGTH-1 && j == MAP_WIDTH-1 {
				terrain[i][j] = "┘"
				continue
			}
			if i == 0 || i == MAP_WIDTH-1 {
				terrain[i][j] = "─"
				continue
			}
			if j == 0 || j == MAP_LENGTH-1 {
				terrain[i][j] = "│"
				continue
			}
			terrain[i][j] = TERRAIN_CHAR
		}
	}

	for i := 0; i < len(snake); i++ {
		terrain[snake[i][0]][snake[i][1]] = USER_CHAR
	}

	terrain[apple[0]][apple[1]] = APPLE_CHAR
	return terrain
}
