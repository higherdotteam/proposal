package main

//https://api.slack.com/docs/message-buttons#crafting_your_message

import "fmt"
import "github.com/nlopes/slack"
import "os"
import "time"
import "strings"

var UserState map[string]int = make(map[string]int)
var UserAnswers map[int]string = make(map[int]string)
var Me string
var Questions []string

func handleRtm(rtm *slack.RTM) {

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				fmt.Println(ev.Msg.Type, ev.Msg.Channel, ev.Msg.User, ev.Msg.Text)

				if ev.Msg.User != Me && strings.HasPrefix(ev.Msg.Channel, "D") {
					from := ev.Msg.User

					switch UserState[from] {
					case 0:
						m := rtm.NewOutgoingMessage(fmt.Sprintf("I'm going to ask you %d questions.\n%s",
							len(Questions), Questions[0]), ev.Msg.Channel)
						rtm.SendMessage(m)
					default:
						UserAnswers[UserState[from]] = ev.Msg.Text
						if UserState[from] < len(Questions) {
							m := rtm.NewOutgoingMessage(fmt.Sprintf("%s", Questions[UserState[from]]), ev.Msg.Channel)
							rtm.SendMessage(m)
						} else {
							fmt.Println(UserAnswers)
							api := slack.New(os.Getenv("SLACK_PROPOSAL_ADMIN"))
							name := "p" + fmt.Sprintf("%d", time.Now().Unix())
							g, err := api.CreateGroup(name)
							if err == nil {
								api.InviteUserToGroup(g.ID, ev.Msg.User)
								api.InviteUserToGroup(g.ID, Me)
							}

							m := rtm.NewOutgoingMessage("All done. I have created a new private channel here: #"+name, ev.Msg.Channel)
							rtm.SendMessage(m)

							buffer := ""
							for i, q := range Questions {
								buffer += fmt.Sprintf("%d. %s\n", i+1, q)
								buffer += fmt.Sprintf("%s\n\n", UserAnswers[i+1])
							}
							if err == nil {
								m = rtm.NewOutgoingMessage(buffer, g.ID)
								rtm.SendMessage(m)
							}
						}
					}
					UserState[from] += 1
					if UserState[from] > len(Questions) {
						UserState[from] = 0
					}
				}

			}
		}
	}

}

func main() {
	Questions = []string{"First Name (Homeowner)?",
		"And what's their last name?",
		"What's the property address were we will be installing?",
		"I need as many phone numbers as you have for this customer. It's important for us to contact them. Please provide at least one but 2 or three would be better.",
		"What's the email address. This should be a working email. If you done have one just type in scared to ask for and email. Just kidded type none.",
		"Who is the sales rep here or opener. Please give your first and last name.",
		"If there is another sales rep involved besides yourself please give me their full name.",
		"What utility company does this customer get there power from?",
		"What is the account number off their electric bill?",
		"I need to make sure I get this to you on time. What is the date and time you are going to go back and close this customer. Please give a full date and time.",
		"Please specify what products are going to be used in this proposal.",
		"What loan plan would you like to offer this customer. If you don't know just type I don't know.",
		"How much expected Savings?",
		"If we discover that the home is inefficient how do we pay for it?",
		"If you have a custom package you use often then please tell me the package number and I'll bust it out for you in no time.",
		"What power offset should i use on this proposal.",
		"If this is a PPA please tell me the price per kWh you would like to use and the escalator you would like.",
		"If you can think of anything we didn't cover here or if we need to change any of your answers this is the time to tell me. Any notes and as many notes as you can always help. The more the merrier."}

	fmt.Println("listening for proposals...")
	api := slack.New(os.Getenv("SLACK_PROPOSAL_BOT"))
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
