package internal

import "regexp"

type UserCommand struct {
	Command   string
	Arguments []string
}

// [args --text value]
func argumentFilter(commandList []string) UserCommand {
	regexValidator := regexp.MustCompile("(?m)-")

	commandSet := false

	var commandDefinition UserCommand

	for _, argument := range commandList {
		IsMatch := regexValidator.MatchString(argument)
		if !IsMatch && !commandSet {
			commandDefinition.Command = argument
			commandSet = true
	} else if IsMatch {
		commandDefinition.Arguments = append(commandDefinition.Arguments, argument)
	}
	}

	return commandDefinition
}