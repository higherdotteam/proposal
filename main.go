package main

import "fmt"
import "github.com/nlopes/slack"
import "os"
import "time"
import "strings"

func handleRtm(rtm *slack.RTM) {

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				//fmt.Println(ev.Msg.Type, ev.Msg.Channel, ev.Msg.User, ev.Msg.Username, ev.Msg.BotID)

				if strings.HasPrefix(ev.Msg.Channel, "D") {
					name := ev.Msg.Text
					api := slack.New(os.Getenv("BOT"))
					api.CreateGroup(name)
					api.InviteUserToGroup(name, ev.Msg.User)
					m := rtm.NewOutgoingMessage("Okay making a new private channel for: "+name, ev.Msg.Channel)
					rtm.SendMessage(m)
				}

				//name := ev.Msg.Channel

				//h["text"] = ev.Msg.Text
				//h["time"] = ev.Msg.Timestamp
				//h["who"] = ev.Msg.User

			}
		}
	}

}

func main() {
	fmt.Println("listening for proposals...")
	api := slack.New(os.Getenv("BOT"))

	rtm := api.NewRTM()

	go rtm.ManageConnection()
	go handleRtm(rtm)

	for {
		time.Sleep(time.Second)
	}
}
