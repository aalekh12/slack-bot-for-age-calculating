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
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4584469582227-4584699882722-4GJ4QC1fXDLFCJlcVjQojg95")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A04H00BHXDL-4656588089909-2e41ba28169592dbf1477b7fd2594e6482ced21ba2c7b777be82fffdfe42dcf1")

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
			age := 2023 - yob
			r := fmt.Sprintf("Your are  %d year Old", age)
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
