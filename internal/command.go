package internal

import (
	"errors"
	"flag"
	"os"
)

type Command interface {
	Name() string
	Example() string
	Help() string
	LongHelp() string
	Register(*flag.FlagSet)
	Run()
}

type CommandRoot struct {
	Name string
	Commands []Command
}

func CommandInit(name string) *CommandRoot {
	return &CommandRoot {
		Name: name,
	}
}

func (cr *CommandRoot) Start(commandList []Command) error {
	if len(commandList) == 0 {
		return errors.New("command line initialization requires one or more commands")
	}

	cr.Commands = commandList

	if len(os.Args) < 2 {
		return errors.New("please pass some command")
	}

	userPassedArguments := os.Args[1:]
}