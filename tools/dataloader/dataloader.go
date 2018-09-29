package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

func getGames() (games []Game) {
	raw, err := ioutil.ReadFile("./season-1718.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &games)
	return games
}
func main() {

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	if err != nil {
		fmt.Println("Error creating session:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	games := getGames()
	for _, game := range games {
		av, err := dynamodbattribute.MarshalMap(game)

		if err != nil {
			fmt.Println("Got error marshalling map:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String("Games"),
		}
		_, err = svc.PutItem(input)
		if err != nil {
			fmt.Println("Got error calling PutItem:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Successfully added '", game.HomeTeam, "' vs '", game.AwayTeam, "' to Games table")
	}
}
