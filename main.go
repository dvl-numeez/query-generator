package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ollama/ollama/api"
)

func main() {

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> Enter your query : ")
		reader.Scan()
		query := reader.Text()
		if len(query) == 0 {
			continue
		}
		startTime := time.Now()
		generateQuery(query)
		fmt.Println("..................................................")
		fmt.Println("Time taken to make query: ", time.Since(startTime))
		fmt.Println("..................................................")
	}

}

func generateQuery(statement string) {
	fmt.Println()
	fmt.Println("generating.............")
	fmt.Println()
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Println(err)
	}
	request := &api.GenerateRequest{
		Model:  "codellama",
		Prompt: fmt.Sprintf("%s\n\n%s\n\n%s\n\nUser instruction: %s", instruction, dbInstruction, datastructureInstruction, statement),
		Stream: new(bool),
		System: systemPrompt,
	}
	ctx := context.Background()
	respFunc := func(resp api.GenerateResponse) error {
		fmt.Println(resp.Response)
		return nil
	}

	err = client.Generate(ctx, request, respFunc)
	if err != nil {
		log.Println(err)
	}
}
