package commands

import (
	"flag"
	"fmt"
	"net/http"
)

type Start struct {
	Port    int
	Version string
	Helpf   bool
}

const helpTextStart = "The start command starts the application server"
const exampleTextStart = `
	expertscli start --port 3030 --version 1.0.0
	expertscli start --version -v 1.0.0
`

func (cmd *Start) Name() string {
	return "start"
}

func (cmd *Start) Example() string {
	return exampleTextStart
}

func (cmd *Start) Help() string {
	return helpTextStart
}

func (cmd *Start) Register(fs *flag.FlagSet) {
	fs.IntVar(&cmd.Port, "port", 8080, "Port to start the application server")
	fs.StringVar(&cmd.Version, "version", "", "Version of the application")
	fs.BoolVar(&cmd.Helpf, "help", false, "Help for start command")
}

func (cmd *Start) Run() {
	if cmd.Helpf {
		fmt.Println(cmd.Help())
		return
	}

	if cmd.Version == "" {
		fmt.Println("Please pass the version")
		return
	}

	fmt.Printf("Servidor %vs rodando na porta %v", cmd.Version, cmd.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", cmd.Port), nil)
}
