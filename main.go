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
	isLost    bool
	isWon     bool
}

func initModel() model {

	terrain := [MAP_LENGTH][MAP_WIDTH]string{}

	apple := randomCoordinates(1, MAP_LENGTH-1)
	snake := [][2]int{{MAP_LENGTH / 2, MAP_WIDTH / 2}}

	for areCoordinatesInSnake(snake, apple) {
		apple = randomCoordinates(1, MAP_LENGTH-1)
	}

	return model{
		counter:   0,
		direction: UP,
		terrain:   renderMap(terrain, snake, apple),
		snake:     snake,
		apple:     apple,
		isLost:    false,
		isWon:     false,
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
				if !(m.direction == DOWN && len(m.snake) > 1) {
					m.direction = UP
				}
			case "down", "s":
				if !(m.direction == UP && len(m.snake) > 1) {
					m.direction = DOWN
				}
			case "left", "a":
				if !(m.direction == RIGHT && len(m.snake) > 1) {
					m.direction = LEFT
				}
			case "right", "d":
				if !(m.direction == LEFT && len(m.snake) > 1) {
					m.direction = RIGHT
				}
			}
			if (m.isLost || m.isWon) && msg.String() == "b" {
				return initModel(), tick()
			}
		}
	case tickMsg:
		{
			newPosition := m.snake[0]
			switch m.direction {
			case UP:
				if m.direction != DOWN {
					newPosition[0]--
				}
			case DOWN:
				if m.direction != UP {
					newPosition[0]++
				}
			case LEFT:
				if m.direction != RIGHT {
					newPosition[1]--
				}
			case RIGHT:
				if m.direction != LEFT {
					newPosition[1]++
				}
			}

			if len(m.snake) == ((MAP_LENGTH-2)*(MAP_WIDTH-2) - 1) {
				println("You won the game")
				m.isWon = true
				return m, nil
			}
			if areCoordinatesInSnake(m.snake[:len(m.snake)-1], newPosition) {
				m.isLost = true
				return m, nil
			}

			isSnakeOutOffBorder := newPosition[0] <= 0 || newPosition[0] >= MAP_LENGTH-1 || newPosition[1] <= 0 || newPosition[1] >= MAP_WIDTH-1

			if isSnakeOutOffBorder {
				m.isLost = true
				return m, nil
			}

			newPositionSlice := [][2]int{{newPosition[0], newPosition[1]}}

			m.snake = append(newPositionSlice, m.snake...)
			isToRemoveValue := true

			isSnakeOnApple := newPosition[0] == m.apple[0] && newPosition[1] == m.apple[1]

			if isSnakeOnApple {
				isToRemoveValue = false
				m.apple = randomCoordinates(1, MAP_LENGTH-1)
				for areCoordinatesInSnake(m.snake, m.apple) {
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

	s := ""
	if m.isLost {
		s += "You lost the game with " + fmt.Sprint(len(m.snake)) + " points\n Press b to play again\n"
		s += "Or press q to exit\n"
		return s
	}

	if m.isWon {
		s += "You won the game with " + fmt.Sprint(len(m.snake)) + " points\n Press b to play again\n"
		s += "Or press q to exit\n"
		return s
	}
	s += "This is the snake game\n"

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
