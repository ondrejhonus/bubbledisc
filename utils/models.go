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
}

func (t Track) Title() string       { return fmt.Sprintf("Track %02d", t.Index+1) }
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
	return Model{List: l}
}

func (m Model) Init() tea.Cmd {
	return nil
}
