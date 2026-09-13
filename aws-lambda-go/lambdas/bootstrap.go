package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type App struct {
	id string
}

func newApp(id string) *App {
	return &App{
		id: id,
	}
}

func (app *App) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Content-Type":                 "application/json",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "Content-Type,Authorization",
	}

	if request.HTTPMethod == http.MethodGet && (request.Resource == "/hi" || request.Path == "/hi") {
		respBody := map[string]string{
			"message": "Hi you have hit some route",
		}
		respJSON, err := json.Marshal(respBody)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers:    headers,
			}, err
		}
		return events.APIGatewayProxyResponse{
			Body:       string(respJSON),
			StatusCode: http.StatusOK,
			Headers:    headers,
		}, nil
	}

	notFoundBody, _ := json.Marshal(map[string]string{
		"message": "Not Found",
	})
	return events.APIGatewayProxyResponse{
		Body:       string(notFoundBody),
		StatusCode: http.StatusNotFound,
		Headers:    headers,
	}, nil
}

func main() {
	id := "someString"

	app := newApp(id)

	lambda.Start(app.Handler)
}
