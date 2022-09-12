package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
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
		"api_key": "38f86475574fdc5f19cf74e393dcc24f568bb68ca6a213a711ab63ecac59adab",
	}

	search := g.NewGoogleSearch(parameter, "38f86475574fdc5f19cf74e393dcc24f568bb68ca6a213a711ab63ecac59adab")
	results, err := search.GetJSON()
	if err != nil {
		return err.Error()
	}
	images_results := results["images_results"].([]interface{})
	return images_results[0].(map[string]interface{})["original"].(string)
}

func main() {
	bot := slacker.NewClient("xoxb-4063480188803-4063522353843-XK840kQJ1g8lc51Qi7OBWXCj", "xapp-1-A041VF9U51R-4064536318866-fc0897c14b4616de3e0cecd5668f57745d9876c7222bde5ea49b89e023dd7eee")

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
