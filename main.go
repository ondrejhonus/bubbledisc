package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/ondrejhonus/bubbledisc/share"
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
			exec.Command("pkill", "mpv")
			return m, tea.Quit
		case "enter":
			if !m.Selected {
				SelectedIndex := m.List.Index()
				items := m.List.Items()

				for i := range items {
					t := items[i].(utils.Track)
					t.Playing = (i == SelectedIndex)
					items[i] = t
				}

				m.List.SetItems(items)

				item := items[SelectedIndex].(utils.Track)
				share.PlayTrack(item.Index)
				// m.Selected = true
				return m, nil
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
		b.WriteString(share.HelpBar())
	}

	return b.String()
}

func main() {
	p := tea.NewProgram(localModel{Model: utils.InitialModel()}, tea.WithoutBracketedPaste(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
