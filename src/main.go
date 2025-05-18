package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type track struct {
	title string
	duration string
	index   int
}

func (t track) Title() string       { return fmt.Sprintf("Track %02d", t.index+1) }
func (t track) Description() string { return fmt.Sprintf("%s (%s)", t.title, t.duration) }
func (t track) FilterValue() string { return t.title }

var titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))

type model struct {
	list     list.Model
	selected bool
}

func initialModel() model {
	items := []list.Item{
		track{title: "Fake Song A", duration: "3:45", index: 0},
		track{title: "Fake Song B", duration: "4:20", index: 1},
		track{title: "Fake Song C", duration: "2:58", index: 2},
	}

	l := list.New(items, list.NewDefaultDelegate(), 40, 10)
	l.Title = "📀 Inserted CD — Track List"
	l.SetShowHelp(true)

	return model{
		list:     l,
		selected: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if !m.selected {
				item := m.list.SelectedItem().(track)
				playTrack(item.index + 1)
				m.selected = true
				return m, nil
			}
		}
	}
	newList, cmd := m.list.Update(msg)
	m.list = newList
	return m, cmd
}

func (m model) View() string {
	if m.selected {
		return titleStyle.Render("▶️  Playing track... Press q to quit.") + "\n"
	}
	return m.list.View()
}

func playTrack(trackNum int) {
	cmd := exec.Command("mpv", fmt.Sprintf("cdda:// --cdrom-device=/dev/sr0 --track=%d", trackNum))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Start() // Run in background
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

