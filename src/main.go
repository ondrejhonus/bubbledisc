package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type track struct {
	title    string
	duration string
	index    int
}

func (t track) Title() string       { return fmt.Sprintf("Track %02d", t.index+1) }
func (t track) Description() string { return fmt.Sprintf("%s (%s)", t.title, t.duration) }
func (t track) FilterValue() string { return t.title }

type model struct {
	list   list.Model
	width  int
	height int
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
				playTrack(t.index + 1)
			}
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.list.View()
}

func playTrack(trackNum int) {
	cmd := exec.Command("mpv", fmt.Sprintf("cdda:// --cdrom-device=/dev/sr0 --track=%d", trackNum))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Start() // Run without waiting
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithoutBracketedPaste())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
