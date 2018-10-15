package golang

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"log"
	"net/http"
	"os"
)

var (
	db = getDynamodbClient()
)

type Game struct {
	Date                   string `json:"Date"`     // Match Date (dd/mm/yy)
	HomeTeam               string `json:"HomeTeam"` // Home Team
	AwayTeam               string `json:"AwayTeam"` // Away Team
	FullTimeHomeTeamGoals  int    `json:"FTHG"`     // Full Time Home Team Goals
	FullTimeAwayTeamGoals  int    `json:"FTAG"`     // Full Time Away Team Goals
	FullTimeResult         string `json:"FTR"`      // Full Time Result (H=Home Win, D=Draw, A=Away Win)
	HalfTimeHomeTeamGoals  int    `json:"HTHG"`     // Half Time Home Team Goals
	HalfTimeAwayTeamGoals  int    `json:"HTAG"`     // Half Time Away Team Goals
	HalfTimeResult         string `json:"HTR"`      // Half Time Result (H=Home Win, D=Draw, A=Away Win)
	Referee                string `json:"Referee"`  // Match Referee
	HomeTeamShots          int    `json:"HS"`       // Home Team Shots
	AwayTeamShots          int    `json:"AS"`       // Away Team Shots
	HomeTeamShotsOnTarget  int    `json:"HST"`      // Home Team Shots on Target
	AwayTeamShotsOnTarget  int    `json:"AST"`      // Away Team Shots on Target
	HomeTeamFoulsCommitted int    `json:"HF"`       // Home Team Fouls Committed
	AwayTeamFoulsCommitted int    `json:"AF"`       // Away Team Fouls Committed
	HomeTeamCorners        int    `json:"HC"`       // Home Team Corners
	AwayTeamCorners        int    `json:"AC"`       // Away Team Corners
	HomeTeamYellowCards    int    `json:"HY"`       // Home Team Yellow Cards
	AwayTeamYellowCards    int    `json:"AY"`       // Away Team Yellow Cards
	HomeTeamRedCards       int    `json:"HR"`       // Home Team Red Cards
	AwayTeamRedCards       int    `json:"AR"`       // Away Team Red Cards
}

func getDynamodbClient() dynamodbiface.DynamoDBAPI {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION"))},
	)
	return dynamodb.New(sess)
}

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	ht, ok := request.QueryStringParameters["HomeTeam"]
	if !ok {
		return events.APIGatewayProxyResponse{
			Body:       "Missing param \"HomeTeam\"",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	at, ok := request.QueryStringParameters["AwayTeam"]
	if !ok {
		return events.APIGatewayProxyResponse{
			Body:       "Missing param \"AwayTeam\"",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE")),
		Key: map[string]*dynamodb.AttributeValue{
			"HomeTeam": {
				S: aws.String(ht),
			},
			"AwayTeam": {
				S: aws.String(at),
			},
		},
	}

	res, err := db.GetItem(input)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Not Found",
			StatusCode: http.StatusNotFound,
		}, nil
	}

	g := new(Game)
	if dynamodbattribute.UnmarshalMap(res.Item, &g) != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	js, err := json.Marshal(g)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(js),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
