package libs

import (
	"regexp"
	"strings"
)

var lists []ICommand

func NewCommands(cmd *ICommand) {
	lists = append(lists, *cmd)
}

func GetList() []ICommand {
	return lists
}

func HasCommand(name string) bool {
	var prefix string
	pattern := regexp.MustCompile(`[?!.#]`)
	for _, f := range pattern.FindAllString(name, -1) {
		prefix = f
	}
	for _, cmd := range lists {
		re := regexp.MustCompile(`^` + cmd.Name + `$`)
		if valid := len(re.FindAllString(strings.ReplaceAll(name, prefix, ""), -1)) > 0; valid {
			return true
		}
	}
	return false
}
