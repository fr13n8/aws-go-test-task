package main

import (
	"bytes"
	"encoding/json"
	"github.com/fr13n8/myservice/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/twpayne/go-geom"
)

type Response events.APIGatewayProxyResponse

func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer
	var points []models.Point
	err := json.Unmarshal([]byte(request.Body), &points)
	if err != nil {
		return Response{StatusCode: 400}, err
	}

	var polygon []geom.Coord
	for _, point := range points {
		polygon = append(polygon, geom.Coord{point.X, point.Y, point.Z})
	}

	area := geom.NewPolygon(geom.XYZ).MustSetCoords([][]geom.Coord{
		polygon,
	}).Area()

	body, err := json.Marshal(map[string]interface{}{
		"area": area,
	})

	if err != nil {
		return Response{
			StatusCode:      404,
			Headers: map[string]string{
				"Content-Type":           "application/json",
				"Access-Control-Allow-Origin": "*",
			},
		}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
