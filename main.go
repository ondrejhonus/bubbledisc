package main

import (
	"log"
	"strings"

	"github.com/ondrejhonus/bubbledisc/utils"

	tea "github.com/charmbracelet/bubbletea"
)

type localModel struct {
	utils.Model
}

func (m localModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.List.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if t, ok := m.List.SelectedItem().(utils.Track); ok {
				utils.PlayTrack(t.Index + 1)
			}
		}
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m localModel) View() string {
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
	p := tea.NewProgram(localModel{Model: utils.InitialModel()}, tea.WithAltScreen(), tea.WithoutBracketedPaste())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
