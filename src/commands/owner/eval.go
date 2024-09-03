package commands

import (
	"encoding/json"
	"hisoka/src/libs"

	"github.com/robertkrimen/otto"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     `c`,
		As:       []string{"c"},
		Tags:     "owner",
		IsPrefix: false,
		IsOwner:  true,
		Execute: func(conn *libs.IClient, m *libs.IMessage) bool {
			vm := otto.New()
			vm.Set("M", m)
			vm.Set("Conn", conn)

			h, err := vm.Run(m.Text)
			if err != nil {
				m.Reply(err.Error())
				return false
			}

			if h.IsObject() {
				var data interface{}
				h, _ := vm.Run("JSON.stringify(" + m.Text + ")")
				json.Unmarshal([]byte(h.String()), &data)
				pe, _ := json.MarshalIndent(data, "", "  ")
				m.Reply(string(pe))
			} else {
				m.Reply(h.String())
			}

			return true
		},
	})
}
