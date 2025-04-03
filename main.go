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

// const instruction = "Make a mongodb query which can be run in the mongo shell and the query should be not be in any kind of language it should be purely a mongodb query for the following : "
const instruction = "Provide only the MongoDB query, without any explanations or additional text, that can be run in the mongo shell for the following: "
const collection = "userManagement"

const datastructureInstruction = `This is the structure of the mongodb document in the collection
{
  "referenceID": "aa4f920c-c9f2-4ad9-8fed-4d5ffd1d5093",
  "items": {
    "response": {
      "detectedObjects": [
        {
          "object": "couch",
          "timeFrame": [
            {
              "startTime": "00:00:00",
              "endTime": "00:00:51"
            }
          ],
          "minimumCoverage": [
            {
              "percentage": "1.81",
              "timestamp": "0:00:48.760"
            }
          ],
          "maximumCoverage": [
            {
              "percentage": "25.33",
              "timestamp": "0:00:47.600"
            }
          ],
          "averageCoverage": "10.08"
        },
        {
          "object": "chair",
          "timeFrame": [
            {
              "startTime": "00:00:00",
              "endTime": "00:00:02"
            },
            {
              "startTime": "00:00:06",
              "endTime": "00:00:51"
            }
          ],
          "minimumCoverage": [
            {
              "percentage": "0.33",
              "timestamp": "0:00:48.560"
            }
          ],
          "maximumCoverage": [
            {
              "percentage": "11.71",
              "timestamp": "0:00:00.160"
            }
          ],
          "averageCoverage": "5.12"
        },
        {
          "object": "tv",
          "timeFrame": [
            {
              "startTime": "00:00:00",
              "endTime": "00:00:39"
            },
            {
              "startTime": "00:00:49",
              "endTime": "00:00:50"
            }
          ],
          "minimumCoverage": [
            {
              "percentage": "0.37",
              "timestamp": "0:00:01.160"
            }
          ],
          "maximumCoverage": [
            {
              "percentage": "4.62",
              "timestamp": "0:00:10.200"
            }
          ],
          "averageCoverage": "2.94"
        },
        {
          "object": "potted plant",
          "timeFrame": [
            {
              "startTime": "00:00:00",
              "endTime": "00:00:20"
            },
            {
              "startTime": "00:00:22",
              "endTime": "00:00:27"
            },
            {
              "startTime": "00:00:29",
              "endTime": "00:00:51"
            }
          ],
          "minimumCoverage": [
            {
              "percentage": "0.18",
              "timestamp": "0:00:29.680"
            }
          ],
          "maximumCoverage": [
            {
              "percentage": "5.84",
              "timestamp": "0:00:47.640"
            }
          ],
          "averageCoverage": "1.95"
        },
        {
          "object": "book",
          "timeFrame": [
            {
              "startTime": "00:00:00",
              "endTime": "00:00:06"
            },
            {
              "startTime": "00:00:08",
              "endTime": "00:00:08"
            },
            {
              "startTime": "00:00:19",
              "endTime": "00:00:19"
            }
          ],
          "minimumCoverage": [
            {
              "percentage": "0.19",
              "timestamp": "0:00:19.840"
            }
          ],
          "maximumCoverage": [
            {
              "percentage": "0.30",
              "timestamp": "0:00:03.560"
            }
          ],
          "averageCoverage": "0.23"
        }
      ]
    }
  },
  "transcription": [
    {
      "startTime": 0,
      "endTime": 4.54,
      "transcribedText": "If you are renovating your drawing room and you are searching for modern design, then"
    },
    {
      "startTime": 4.54,
      "endTime": 5.62,
      "transcribedText": "your search can end here."
    },
    {
      "startTime": 5.84,
      "endTime": 10.5,
      "transcribedText": "So relax and watch each drawing room designs carefully and get some great and innovative"
    },
    {
      "startTime": 10.5,
      "endTime": 15.54,
      "transcribedText": "ideas from these designs that are recommended by top interior designers all over the world."
    },
    {
      "startTime": 15.84,
      "endTime": 21.16,
      "transcribedText": "The space matters a lot while decorating any room like bedroom, drawing room etc."
    },
    {
      "startTime": 22.24,
      "endTime": 27.84,
      "transcribedText": "Before decorating and furnishing your drawing room, decide how much time your family members"
    },
    {
      "startTime": 27.84,
      "endTime": 28.72,
      "transcribedText": "spend there."
    },
    {
      "startTime": 29.52,
      "endTime": 29.98,
      "transcribedText": "If you have any questions or suggestions, please feel free to contact us."
    },
    {
      "startTime": 30,
      "endTime": 36.66,
      "transcribedText": "less space, you need more care to furnish it. If you have more space, you can use the"
    },
    {
      "startTime": 36.66,
      "endTime": 43.86,
      "transcribedText": "space more efficiently. Ornamental pieces such as sculptures, wall hangings, flowers"
    }
  ]
}

Use this to make the queries 

`

var dbInstruction = fmt.Sprintf("Use the following collection to make the query , the collection name is :%s", collection)

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
		Model:  "mistral",
		Prompt: datastructureInstruction+"\n"+instruction + statement,
		Stream: new(bool),
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
