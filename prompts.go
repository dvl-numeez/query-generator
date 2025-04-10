package main

import "fmt"

const instruction = `
You are a PostgreSQL expert. Given a user instruction, your job is to generate a single valid SQL query.

⚠️ RULES:
1. DO NOT include any explanation or commentary.
2. DO NOT format with comments or newlines.
3. DO NOT include schema or connection setup — assume context is already established.
4. Return ONLY the SQL query (like: SELECT * FROM table WHERE ...).
5. Use correct types (strings in single quotes, numbers as-is, and proper SQL syntax).
6. Make sure the query is valid and executable.
`

const tableName = "object_detection"

var dbInstruction = fmt.Sprintf(`The table  name is "%s".`, tableName)

const systemPrompt = `
You are an AI trained to generate strict, valid SQL queries in response to user instructions.

Rules:
1. Respond ONLY with a valid PostgreSQL SQL query — no explanations.
2. Use correct syntax for JSONB columns, including functions like jsonb_array_elements() to extract array elements.
3. Cast data to the correct types (e.g., numeric, timestamp) when needed.
4. The table structure includes JSONB fields — treat them as embedded arrays or objects.

Your output must work in PostgreSQL and assume the table already exists and contains data.
`

const tableSchema = `
Table name is object_detection

table is made by using following query:
CREATE TABLE object_detection ( id UUID PRIMARY KEY,
reference_id UUID,
gcs_source_url TEXT, request_gcs_url TEXT, request_time TIMESTAMP, status VARCHAR(50), update_time TIMESTAMP, processed_video_path TEXT, items JSONB, transcription JSONB );

-- Notes:
-- 'items' and 'transcription' are JSONB columns.
-- 'transcription' contains a JSON array of objects with keys: startTime (float), endTime (float), and transcribedText (string).
-- To access 'transcription', use: jsonb_array_elements(transcription)
-- 'items' contains nested structure: response → detectedObjects (array), each object has timeFrame (array), averageCoverage, etc.
You can consider the following example:
SELECT * FROM object_detection, jsonb_array_elements(transcription) AS t WHERE t->>'transcribedText' ILIKE '%couch%'
`


const jsonBlob =`The items and transcription field looks like as follow :
{
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

`

// const itemsStructure = `
// The item field structure of the table "object_detection" is as follows :
// {
//   "response": {
//     "detectedObjects": [
//       {
//         "object": "couch",
//         "timeFrame": [
//           {
//             "startTime": "00:00:00",
//             "endTime": "00:00:51"
//           }
//         ],
//         "minimumCoverage": [
//           {
//             "percentage": "1.81",
//             "timestamp": "0:00:48.760"
//           }
//         ],
//         "maximumCoverage": [
//           {
//             "percentage": "25.33",
//             "timestamp": "0:00:47.600"
//           }
//         ],
//         "averageCoverage": "10.08"
//       },
//       {
//         "object": "chair",
//         "timeFrame": [
//           {
//             "startTime": "00:00:00",
//             "endTime": "00:00:02"
//           },
//           {
//             "startTime": "00:00:06",
//             "endTime": "00:00:51"
//           }
//         ],
//         "minimumCoverage": [
//           {
//             "percentage": "0.33",
//             "timestamp": "0:00:48.560"
//           }
//         ],
//         "maximumCoverage": [
//           {
//             "percentage": "11.71",
//             "timestamp": "0:00:00.160"
//           }
//         ],
//         "averageCoverage": "5.12"
//       },
//       {
//         "object": "tv",
//         "timeFrame": [
//           {
//             "startTime": "00:00:00",
//             "endTime": "00:00:39"
//           },
//           {
//             "startTime": "00:00:49",
//             "endTime": "00:00:50"
//           }
//         ],
//         "minimumCoverage": [
//           {
//             "percentage": "0.37",
//             "timestamp": "0:00:01.160"
//           }
//         ],
//         "maximumCoverage": [
//           {
//             "percentage": "4.62",
//             "timestamp": "0:00:10.200"
//           }
//         ],
//         "averageCoverage": "2.94"
//       },
//       {
//         "object": "potted plant",
//         "timeFrame": [
//           {
//             "startTime": "00:00:00",
//             "endTime": "00:00:20"
//           },
//           {
//             "startTime": "00:00:22",
//             "endTime": "00:00:27"
//           },
//           {
//             "startTime": "00:00:29",
//             "endTime": "00:00:51"
//           }
//         ],
//         "minimumCoverage": [
//           {
//             "percentage": "0.18",
//             "timestamp": "0:00:29.680"
//           }
//         ],
//         "maximumCoverage": [
//           {
//             "percentage": "5.84",
//             "timestamp": "0:00:47.640"
//           }
//         ],
//         "averageCoverage": "1.95"
//       },
//       {
//         "object": "book",
//         "timeFrame": [
//           {
//             "startTime": "00:00:00",
//             "endTime": "00:00:06"
//           },
//           {
//             "startTime": "00:00:08",
//             "endTime": "00:00:08"
//           },
//           {
//             "startTime": "00:00:19",
//             "endTime": "00:00:19"
//           }
//         ],
//         "minimumCoverage": [
//           {
//             "percentage": "0.19",
//             "timestamp": "0:00:19.840"
//           }
//         ],
//         "maximumCoverage": [
//           {
//             "percentage": "0.30",
//             "timestamp": "0:00:03.560"
//           }
//         ],
//         "averageCoverage": "0.23"
//       }
//     ]
//   }
// }
// `


// const transcriptionStructure = `
// The transcription field structure of the table "object_detection" is as follows :
// [
//   {
//     "startTime": 0,
//     "endTime": 4.54,
//     "transcribedText": "If you are renovating your drawing room and you are searching for modern design, then"
//   },
//   {
//     "startTime": 4.54,
//     "endTime": 5.62,
//     "transcribedText": "your search can end here."
//   },
//   {
//     "startTime": 5.84,
//     "endTime": 10.5,
//     "transcribedText": "So relax and watch each drawing room designs carefully and get some great and innovative"
//   },
//   {
//     "startTime": 10.5,
//     "endTime": 15.54,
//     "transcribedText": "ideas from these designs that are recommended by top interior designers all over the world."
//   },
//   {
//     "startTime": 15.84,
//     "endTime": 21.16,
//     "transcribedText": "The space matters a lot while decorating any room like bedroom, drawing room etc."
//   },
//   {
//     "startTime": 22.24,
//     "endTime": 27.84,
//     "transcribedText": "Before decorating and furnishing your drawing room, decide how much time your family members"
//   },
//   {
//     "startTime": 27.84,
//     "endTime": 28.72,
//     "transcribedText": "spend there."
//   },
//   {
//     "startTime": 29.52,
//     "endTime": 29.98,
//     "transcribedText": "If you have any questions or suggestions, please feel free to contact us."
//   },
//   {
//     "startTime": 30,
//     "endTime": 36.66,
//     "transcribedText": "less space, you need more care to furnish it. If you have more space, you can use the"
//   },
//   {
//     "startTime": 36.66,
//     "endTime": 43.86,
//     "transcribedText": "space more efficiently. Ornamental pieces such as sculptures, wall hangings, flowers"
//   }
// ]
// `

var prompt = fmt.Sprintf(`
%s

%s

%s



Below is the structure of the JSONB column 'items' and 'transcription' stored in the '%s' table:
%s

Based on the above information, generate a strict and valid SQL query in response to the user instruction below.
REMEMBER: 
- Only one SQL query should be returned.
- DO NOT include explanation, comments, or formatting.
- Assume the table is already created and populated.

USER INSTRUCTION:
`, instruction, dbInstruction, tableSchema, tableName, jsonBlob)
