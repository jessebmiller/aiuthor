package main

import (
	"context"
	"fmt"
	"log"
	"io/ioutil"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

// getPrompt opens the prompt file and returns the contents
func getPrompt(promptfile string) string {
	prompt, err := ioutil.ReadFile(promptfile)
	if err != nil {
		log.Fatal(err)
	}
	return string(prompt)
}

func main() {
	fmt.Println("Testing Ollama")

	llm, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		log.Fatal(err)
	}

	completion, err := llm.Call(context.Background(), "Who is the president of the United States?",
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Println(string(chunk))
			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(completion)
}
