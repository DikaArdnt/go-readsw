package commands

import (
	"fmt"
	"hisoka/src/libs"
	"os/exec"
	"runtime"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     `\$`,
		As:       []string{"$"},
		Tags:     "owner",
		IsPrefix: false,
		IsOwner:  true,
		Execute: func(conn *libs.IClient, m *libs.IMessage) bool {
			var cmd *exec.Cmd

			switch runtime.GOOS {
			case "windows":
				cmd = exec.Command("cmd", "/C", m.Text)
			default:
				cmd = exec.Command("sh", "-c", m.Text)
			}

			out, err := cmd.Output()
			if err != nil {
				m.Reply(fmt.Sprintf("%s", err))
				return true
			}
			m.Reply(string(out))

			return true
		},
	})
}
