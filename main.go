package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/ollama/ollama/api"
)

type result struct {
	answers map[string]string
	mux     sync.Mutex
}

var output = result{
	answers: make(map[string]string),
	mux:     sync.Mutex{},
}
const maxConcurrent = 2 
var questions = []string{
	"When did couch first occur in the video?",
	"When did couch last occur in the video?",
	"How many unique objects are there in the video between 00:00:00 to 0:00:47.640?",
	"List all detected objects (including duplicates) in the video between 00:00:00 to 0:00:47.640?",
	"Provide the transcribed text associated with couch shown in the video along with the timestamps.",
	"Which objects were present at the same time when couch was mentioned in the transcription?",
	"List all the unique objects detected in the video.",
	"For objects that appeared more for than 5 seconds in the video, provide the corresponding transcription of those moments.",
	"Provide the list of start times where couch has appeared for more than 3 seconds.",
	"List the objects that appeared for more than 5 seconds in the video.",
	"For how long did chair appear in the video?",
	"When did couch first occur in the video and for how long?",
	"When did couch last occur in the video and for how long?",
	"For how long was chair shown in the video?",
	"How many unique objects are there in the video?",
	"Provide the transcription of the video.",
	"At what timestamp does chair have maximum frame coverage?",
	"At what timestamp does couch have minimum frame coverage?",
	"What is the minimum frame coverage percentage of couch throughout the video?",
	"What is the maximum frame coverage percentage of chair throughout the video?",
	"What is the average frame coverage percentage of couch throughout the video?",
	"Which objects appeared more for than 5 seconds in the video?",
	"Which object appeared for longest time in the video?",
}

func main() {
	wg := sync.WaitGroup{}
	sem := make(chan struct{}, maxConcurrent)
	for _, question := range questions {
		wg.Add(1)
		sem <- struct{}{} // acquire slot

		go func(q string) {
			defer wg.Done()
			generateQuery(q,)
			<-sem // release slot
		}(question)
	}

	wg.Wait()

	for question, answer := range output.answers {
		fmt.Println("Question: ", question)
		fmt.Println("Answer: ", answer)
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println()
	}

}

func generateQuery(statement string) {
	fmt.Println()
	fmt.Println("generating answer for: ", statement)
	fmt.Println()
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Println(err)
	}
	request := &api.GenerateRequest{
		Model:  "codellama",
		Prompt: prompt + "\n" + statement,
		Stream: new(bool),
		System: systemPrompt,
	}
	ctx := context.Background()
	respFunc := func(resp api.GenerateResponse) error {
		fmt.Println("Answer fetched for question : ",statement)
		output.mux.Lock()
		output.answers[statement] = resp.Response
		output.mux.Unlock()
		return nil
	}

	err = client.Generate(ctx, request, respFunc)
	if err != nil {
		log.Println(err)
	}
}
