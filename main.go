package main

import (
	_ "context"
	"fmt"
	"github.com/shomali11/slacker"
	_ "log"
	"os"
)

//Test test
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4471312034385-4458934173107-D4WqQyuRKdAYU3NdWTFWrBEf")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A04DVHWENKT-4452376892918-144ecfe011b3a8eefd0089639c41b9b37181e2249877e724cefeeb317d402684")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
}
