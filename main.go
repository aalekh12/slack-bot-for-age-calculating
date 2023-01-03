package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func PrintCommandEvent(anyalaticschannel <-chan *slacker.CommandEvent) {
	for event := range anyalaticschannel {
		fmt.Println("Running Commans Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Event.ChannelName)
		fmt.Println(event.Command)
		fmt.Println(event.Event)
		fmt.Println(event.Parameters)
		fmt.Println()
	}
}
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4584469582227-4584699882722-doJhAud4WIqXFKGWIRb89VOl")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A04H00BHXDL-4578037636534-cd6078a784608e48bd781b04a0a06d219f9a2207b9278ba850bf09e71c3baaf9")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob Calculator",
		Examples:    []string{"my yob is 2020"},
		Handler: func(botsctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("error")
			}
			age := 2022 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	go PrintCommandEvent(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {

		log.Fatal(err)
	}
}
