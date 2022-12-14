package filesystem

import (
	"strconv"
	"strings"
)

type command struct {
	cmd    string
	arg    string
	output []string
}

func Create(input string) Dir {
	root := Dir{
		name: "/",
	}
	pwd := &root

	for i, c := range readCommandsWithOutput(input) {
		if i == 0 && (c.cmd != "cd" || c.arg != "/") {
			panic("first command must be cd /")
		}

		switch c.cmd {
		case "cd":
			switch c.arg {
			case "/":
				break
			case "..":
				pwd = pwd.Parent()
			default:
				pwd = pwd.findChildDirByName(c.arg)
			}
		case "ls":
			for _, content := range c.output {
				if isDir(content) {
					pwd.dirs = append(pwd.dirs, &Dir{
						name:   strings.Split(content, " ")[1],
						parent: pwd,
					})
				} else {
					parts := strings.Split(content, " ")
					size, _ := strconv.Atoi(parts[0])
					pwd.files = append(pwd.files, file{
						name: parts[1],
						size: int64(size),
					})
				}
			}
		}
	}

	return root
}

func readCommandsWithOutput(input string) []command {
	var commands []command
	for _, line := range strings.Split(input, "\n") {
		if isCommand(line) {
			parts := strings.Split(line, " ")
			c := command{
				cmd: parts[1],
			}
			if len(parts) > 2 {
				c.arg = parts[2]
			}
			commands = append(commands, c)
		} else {
			commands[len(commands)-1].output = append(commands[len(commands)-1].output, line)
		}
	}

	return commands
}

func isCommand(str string) bool {
	return strings.HasPrefix(str, "$")
}

func isDir(str string) bool {
	return strings.HasPrefix(str, "dir ")
}
