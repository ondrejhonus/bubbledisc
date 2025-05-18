package utils

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Track struct {
	Title    string
	Duration string
	Index    int
}

func (t Track) Name() string        { return fmt.Sprintf("Track %02d", t.Index+1) }
func (t Track) Description() string { return fmt.Sprintf("%s (%s)", t.Title, t.Duration) }
func (t Track) FilterValue() string { return t.Title }

var TitleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))

type Model struct {
	List     list.Model
	Width    int
	Height   int
	Selected bool
}

func InitialModel() Model {
	items := []list.Item{
		Track{Title: "Fake Song A", Duration: "3:45", Index: 0},
		Track{Title: "Fake Song B", Duration: "4:20", Index: 1},
		Track{Title: "Fake Song C", Duration: "2:58", Index: 2},
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "📀 CD Contents"
	l.SetShowHelp(false)
	return Model{List: l}
}

func (m Model) Init() tea.Cmd {
	return nil
}
