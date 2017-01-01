package main

import "fmt"
import "github.com/nlopes/slack"
import "os"
import "time"

func handleRtm(rtm *slack.RTM) {

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				/*
					&{{message D3LAES7C5 U035LF6C1 wefwef 1483234485.000004 false [] [] <nil>  false     <nil>      [] <nil> false <nil>  0 T035N23CL []} <nil>} {message D3LAES7C5 U035LF6C1 wefwef 1483234485.000004 false [] [] <nil>  false     <nil>      [] <nil> false <nil>  0 T035N23CL []}

				*/
				//fmt.Println(ev, ev.Msg)
				name := ev.Msg.Text
				m := rtm.NewOutgoingMessage("Okay making a new private channel for: "+name, ev.Msg.Channel)
				rtm.SendMessage(m)
				//api := slack.New(os.Getenv("BOT"))
				//g, _ := api.CreateGroup("name")
				//_, _, _ = api.InviteUserToGroup("name", "user")

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

	for {
		time.Sleep(time.Second)
	}
}
