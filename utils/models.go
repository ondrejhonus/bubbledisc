package utils

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Track struct {
	Name     string
	Duration string
	Index    int
	Playing  bool
}

type Model struct {
	List         list.Model
	Width        int
	Height       int
	Selected     bool
	PlayingIndex int
}

func (t Track) Title() string {
	label := fmt.Sprintf("Track %02d", t.Index+1)
	if t.Playing {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("42")). // green
			Bold(true).
			Render("▶ " + label)
	}
	return label
}
func (t Track) Description() string { return fmt.Sprintf("%s (%s)", t.Name, t.Duration) }
func (t Track) FilterValue() string { return t.Name }

var TitleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))

func InitialModel() Model {
	items := []list.Item{
		Track{Name: "Citizens Of Earth", Duration: "2:40", Index: 0},
		Track{Name: "Threat Level Midnight", Duration: "2:46", Index: 1},
		Track{Name: "Can't Kick Up The Roots", Duration: "2:49", Index: 2},
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "📀 CD Contents"
	l.SetShowHelp(false)
	return Model{
		List:         l,
		PlayingIndex: -1,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
