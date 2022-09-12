package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	g "github.com/serpapi/google-search-results-golang"
	"github.com/shomali11/slacker"
)

// function to load .env file and return env variables
// func goDotEnvVariable(key string) string {
// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	return os.Getenv(key)
// }

type apiConfigData struct {
	WEB_SCRAP_API_KEY string `json:"WEB_SCRAP_API_KEY"`
	SLACK_BOT_TOKEN   string `json:"SLACK_BOT_TOKEN"`
	SLACK_APP_TOKEN   string `json:"SLACK_APP_TOKEN"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	var config apiConfigData
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return apiConfigData{}, err
	}
	return config, nil
}
func webScrap(query string) string {
	apiConfig, x := loadApiConfig(".apiConfig")
	if x != nil {
		log.Fatal(x)
		return ""
	}
	parameter := map[string]string{
		"q":       query,
		"tbm":     "isch",
		"ijn":     "0",
		"api_key": apiConfig.WEB_SCRAP_API_KEY,
	}

	search := g.NewGoogleSearch(parameter, apiConfig.WEB_SCRAP_API_KEY)
	results, err := search.GetJSON()
	if err != nil {
		return err.Error()
	}
	images_results := results["images_results"].([]interface{})
	return images_results[0].(map[string]interface{})["original"].(string)
}

func main() {
	apiConfig, x := loadApiConfig(".apiConfig")
	if x != nil {
		log.Fatal(x)
		return
	}

	bot := slacker.NewClient(apiConfig.SLACK_BOT_TOKEN, apiConfig.SLACK_APP_TOKEN)

	definition := &slacker.CommandDefinition{
		Description: "Enter a query to search for!",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("query")
			url := webScrap(query)
			response.Reply(url)
		},
	}

	bot.Command("search for <query>", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
