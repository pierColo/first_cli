package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Directions int8

type model struct {
	direction Directions
	counter   int
	terrain   [MAP_LENGTH][MAP_WIDTH]string
	snake     [][2]int
	apple     [2]int
}

func initModel() model {

	terrain := [MAP_LENGTH][MAP_WIDTH]string{}

	apple := randomCoordinates(1, MAP_LENGTH-1)
	snake := [][2]int{{MAP_LENGTH / 2, MAP_WIDTH / 2}}

	for isAppleInSnake(snake, apple) {
		apple = randomCoordinates(1, MAP_LENGTH-1)
	}

	return model{
		counter:   0,
		direction: UP,
		terrain:   renderMap(terrain, snake, apple),
		snake:     snake,
		apple:     apple,
	}
}

func (m model) Init() tea.Cmd {
	return tick()

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		{
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "up", "w":
				m.direction = UP
			case "down", "s":
				m.direction = DOWN
			case "left", "a":
				m.direction = LEFT
			case "right", "d":
				m.direction = RIGHT
			}
		}
	case tickMsg:
		{
			newPosition := m.snake[0]
			switch m.direction {
			case UP:
				newPosition[0]--
			case DOWN:
				newPosition[0]++
			case LEFT:
				newPosition[1]--
			case RIGHT:
				newPosition[1]++
			}
			newPositionSlice := [][2]int{{newPosition[0], newPosition[1]}}

			m.snake = append(newPositionSlice, m.snake...)
			isToRemoveValue := true

			if newPosition[0] == m.apple[0] && newPosition[1] == m.apple[1] {
				isToRemoveValue = false
				m.apple = randomCoordinates(1, MAP_LENGTH-1)
				for isAppleInSnake(m.snake, m.apple) {
					m.apple = randomCoordinates(1, MAP_LENGTH-1)
				}
				m.terrain[m.apple[0]][m.apple[1]] = APPLE_CHAR
			}

			if isToRemoveValue {
				removedValue := m.snake[len(m.snake)-1]
				m.snake = m.snake[:len(m.snake)-1]
				m.terrain[removedValue[0]][removedValue[1]] = TERRAIN_CHAR
			}

			for i := 0; i < len(m.snake); i++ {
				m.terrain[m.snake[i][0]][m.snake[i][1]] = USER_CHAR
			}
			m.counter++
			return m, tick()
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "This is the snake game\n"
	s += "Press q to quit.\n"
	for i := 0; i < MAP_LENGTH; i++ {
		for j := 0; j < MAP_WIDTH; j++ {
			s += m.terrain[i][j]
		}
		s += "\n"
	}
	s += "\n Your current points are: " + fmt.Sprint(len(m.snake)) + "\n"

	return s
}

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func main() {
	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
