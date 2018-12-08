package main

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

//
type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

//
type UnityCloudBuildWebHookData struct {
  BuildNumber int `json:"buildNumber"`
  BuildStatus string `json:"buildStatus"`
  ProjectGuid string `json:"projectGuid"`
}

//
func Handler(ctx context.Context, request Request) (Response, error) {
  fmt.Println("Received body: ", request.Body)

  jsonStr := request.Body
  ucbWebHookData := UnityCloudBuildWebHookData{}

  json.Unmarshal([]byte(jsonStr), &ucbWebHookData)

  var resultMessage string

  switch ucbWebHookData.BuildStatus {
  case "success":
    //
    resultMessage = "Deploy to Hockeyapp"
  default:
    resultMessage = "No Build"
  }

  resp := Response{
    StatusCode:      200,
    IsBase64Encoded: false,
    Body:            resultMessage,
    Headers: map[string]string{
      "Content-Type":           "application/json",
      "X-MyCompany-Func-Reply": "ucb-handler",
    },
  }

  return resp, nil
}

func main() {
  lambda.Start(Handler)
}
