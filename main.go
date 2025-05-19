package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/ondrejhonus/bubbledisc/utils"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	List     list.Model
	Width    int
	Height   int
	Selected bool
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.List.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			exec.Command("killall", "mpv")
			return m, tea.Quit
		case "enter":
			if t, ok := m.List.SelectedItem().(utils.Track); ok {
				utils.PlayTrack(&m, t.Index)
			}
		}
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	var b strings.Builder

	if m.Selected {
		b.WriteString(utils.TitleStyle.Render("▶️  Playing track... Press q to quit.") + "\n")
	} else {
		b.WriteString(m.List.View())
		b.WriteString("\n")
		b.WriteString(utils.HelpBar())
	}

	return b.String()
}

func main() {
	p := tea.NewProgram(utils.InitialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
