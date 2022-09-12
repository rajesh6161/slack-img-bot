package main

import (
	"context"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	g "github.com/serpapi/google-search-results-golang"
	"github.com/shomali11/slacker"
)

// function to load .env file and return env variables
func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func webScrap(query string) string {

	parameter := map[string]string{
		"q":       query,
		"tbm":     "isch",
		"ijn":     "0",
		"api_key": goDotEnvVariable("WEB_SCRAP_API_KEY"),
	}

	search := g.NewGoogleSearch(parameter, goDotEnvVariable("WEB_SCRAP_API_KEY"))
	results, err := search.GetJSON()
	if err != nil {
		return err.Error()
	}
	images_results := results["images_results"].([]interface{})
	return images_results[0].(map[string]interface{})["original"].(string)
}

func main() {

	bot := slacker.NewClient(goDotEnvVariable("SLACK_BOT_TOKEN"), goDotEnvVariable("SLACK_APP_TOKEN"))

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
