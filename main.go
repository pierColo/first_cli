package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Directions int8

const (
	UP    Directions = 1
	DOWN  Directions = 2
	LEFT  Directions = 3
	RIGHT Directions = 4
)

// mapLength:= 10
// width := 20
type model struct {
	direction Directions
	counter   int
	terrain   [11][21]string
	snake     [][2]int
}

func initModel() model {

	terrain := [11][21]string{}

	snake := [][2]int{{5, 5}}
	for i := 0; i < 11; i++ {
		for j := 0; j < 21; j++ {
			if i == 0 && j == 0 {
				terrain[i][j] = "┌"
				continue
			}
			if i == 0 && j == 19 {
				terrain[i][j] = "┐"
				continue
			}
			if i == 9 && j == 0 {
				terrain[i][j] = "└"
				continue
			}
			if i == 9 && j == 19 {
				terrain[i][j] = "┘"
				continue
			}
			if i == 0 || i == 9 {
				terrain[i][j] = "─"
				continue
			}
			if j == 0 || j == 19 {
				terrain[i][j] = "│"
				continue
			}
			terrain[i][j] = " "

		}
	}

	for i := 0; i < len(snake); i++ {
		terrain[snake[i][0]][snake[i][1]] = "#"
	}

	return model{
		counter:   0,
		direction: UP,
		terrain:   terrain,
		snake:     snake,
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
			m.counter++
			return m, tick()
		}
	}

	return m, nil
}

func (m model) View() string {
	s := ""

	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			s += m.terrain[i][j]
		}
		s += "\n"
	}
	s += "\n This is the direction: " + fmt.Sprint(m.direction) + "\n"
	s += "\n This is the counter: " + fmt.Sprint(m.counter) + "\n"
	s += "\nPress q to quit.\n"

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
