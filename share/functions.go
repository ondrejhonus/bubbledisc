package share

import (
	"os/exec"

	"github.com/charmbracelet/lipgloss"
)

func HelpBar() string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render("j/k: up/down • Enter: Play • q: Quit")
}

func PlayTrack(trackNum int) {
	// cmd := exec.Command("mpv", fmt.Sprintf("cdda:// --cdrom-device=/dev/sr0 --track=%d", trackNum))
	// cmd := exec.Command("mpv", "--no-video", "https://www.youtube.com/watch?v=HzZGLOfIxkM")
	var cmd *exec.Cmd
	switch trackNum {
	case 0:
		cmd = exec.Command("mpv", "--no-video", "https://www.youtube.com/watch?v=E0ftGt9srSk")
	case 1:
		cmd = exec.Command("mpv", "--no-video", "https://www.youtube.com/watch?v=uIJdDAFTkqA")
	case 2:
		cmd = exec.Command("mpv", "--no-video", "https://www.youtube.com/watch?v=Qv0XH2qgGyY")
	}
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	_ = cmd.Start()

}
