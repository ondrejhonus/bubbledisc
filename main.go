package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ondrejhonus/bubbledisc/utils"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type track struct {
	title    string
	duration string
	index    int
}

func (t track) Title() string       { return fmt.Sprintf("Track %02d", t.index+1) }
func (t track) Description() string { return fmt.Sprintf("%s (%s)", t.title, t.duration) }
func (t track) FilterValue() string { return t.title }

var titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))

type model struct {
	list     list.Model
	width    int
	height   int
	selected bool
}

func initialModel() model {
	items := []list.Item{
		track{title: "Fake Song A", duration: "3:45", index: 0},
		track{title: "Fake Song B", duration: "4:20", index: 1},
		track{title: "Fake Song C", duration: "2:58", index: 2},
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "📀 CD Contents"
	l.SetShowHelp(false)
	return model{list: l}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if t, ok := m.list.SelectedItem().(track); ok {
				utils.PlayTrack(t.index + 1)
			}
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var b strings.Builder

	if m.selected {
		b.WriteString(titleStyle.Render("▶️  Playing track... Press q to quit.") + "\n")
	} else {
		b.WriteString(m.list.View())
		b.WriteString("\n")
		b.WriteString(utils.HelpBar())
	}

	return b.String()
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithoutBracketedPaste())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
