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

type model struct {
	direction Directions
	counter   int
}

func initModel() model {

	return model{
		counter:   0,
		direction: UP,
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
			case "up", "k":
				m.direction = UP
			case "down", "j":
				m.direction = DOWN
			case "left", "h":
				m.direction = LEFT
			case "right", "l":
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
	s := " "
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
