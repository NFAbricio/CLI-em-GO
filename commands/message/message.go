package message

import (
	"flag"
	"fmt"
)

type Message struct {
	Text string
	Helpf   bool
}

const helpTextStart = "Show a simple message"
const exampleTextStart = `
	expertscli message --text [print text]
	print Text: string
	expertscli message --text Hello World

`

func (cmd *Message) Name() string {
	return "message"
}

func (cmd *Message) Example() string {
	return exampleTextStart
}

func (cmd *Message) Help() string {
	return helpTextStart
}

func (cmd *Message) Register(fs *flag.FlagSet) {
	fs.StringVar(&cmd.Text, "text", "", "printed text")
	fs.BoolVar(&cmd.Helpf, "help", false, "Help for start command")
}

func (cmd *Message) Run() {
	if cmd.Helpf {
		fmt.Println(cmd.Help())
		return
	}

	if cmd.Text == "" {
		fmt.Println("Please pass the text flag")
		return
	}

	fmt.Printf("The text is %s", cmd.Text)
}
