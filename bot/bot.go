package bot

import (
	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func CreateBot(token string) {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	kbot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal("Bot is not created")
		return
	}

	fmt.Println("bot has been created")

	kbot.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello")
	})

	kbot.Handle(tele.OnText, func(m tele.Context) error {

		log.Print(m.Message().Payload, m.Text())
		payload := m.Message().Payload

		switch payload {
		case "hello":
			err = m.Send(fmt.Sprintf("Hello I'm kbot %s!"))
		}

		return err

	})

	kbot.Start()

}
