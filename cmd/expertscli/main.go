package main

import (
	"fmt"

	"github.com/fafa1/goCurso/cliUdemy/commands"
	"github.com/fafa1/goCurso/cliUdemy/commands/message"
	"github.com/fafa1/goCurso/cliUdemy/internal"
)




var commandList = []internal.Command{
	new(commands.Start),
	new(message.Message),
}

func main() {
	err := internal.CommandInit("expertscli").Start(commandList)

	if err != nil {
		fmt.Println(err.Error())
	}

}