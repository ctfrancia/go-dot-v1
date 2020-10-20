package prompts

import (
	"fmt"
	"io"
)

type PromptValues struct {
	profileManager string
}

func NewPrompt(rd *io.Reader) {
	fmt.Println("hi")
	return
}

func (r *Reader) Prompt(txt string) error {
	return
}
