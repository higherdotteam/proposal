package main

import "fmt"
import "github.com/nlopes/slack"
import "os"
import "time"
import "strings"

var UserState map[string]int = make(map[string]int)
var Me string

func handleRtm(rtm *slack.RTM) {

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				fmt.Println(ev.Msg.Type, ev.Msg.Channel, ev.Msg.User, ev.Msg.Text)

				if ev.Msg.User != Me && strings.HasPrefix(ev.Msg.Channel, "D") {
					//name := ev.Msg.Text
					//api := slack.New(os.Getenv("BOTADMIN"))
					//api.CreateGroup(name)
					//api.InviteUserToGroup(name, ev.Msg.User)

					from := ev.Msg.User

					switch UserState[from] {
					case 0:
						m := rtm.NewOutgoingMessage("Hello, let's start your proposal. What is first name?", ev.Msg.Channel)
						rtm.SendMessage(m)
					case 1:
						m := rtm.NewOutgoingMessage("What is last name?", ev.Msg.Channel)
						rtm.SendMessage(m)
					default:
						m := rtm.NewOutgoingMessage("All done", ev.Msg.Channel)
						rtm.SendMessage(m)
					}
					UserState[from] += 1
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
	list, _ := api.GetUsers()
	for _, u := range list {
		if u.Name == "proposal" {
			Me = u.ID
			break
		}
	}

	rtm := api.NewRTM()

	go rtm.ManageConnection()
	go handleRtm(rtm)

	for {
		time.Sleep(time.Second)
	}
}
