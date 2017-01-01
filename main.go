package main

import "fmt"
import "github.com/nlopes/slack"
import "os"

func handleRtm(rtm *slack.RTM) {

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				fmt.Println(ev, ev.Msg)
				//name := ev.Msg.Channel

				//h["text"] = ev.Msg.Text
				//h["time"] = ev.Msg.Timestamp
				//h["who"] = ev.Msg.User

			}
		}
	}

}

func main() {
	fmt.Println("vim-go")
	api := slack.New(os.Getenv("BOT"))
	rtm := api.NewRTM()

	go rtm.ManageConnection()
	go handleRtm(rtm)
}
