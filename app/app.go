package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/covrom/herthbot/store"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct {
	bot         *tgbotapi.BotAPI
	currentList *store.DayList
}

func New() *App {
	tz := os.Getenv("TZ")
	if tz == "" {
		tz = "Asia/Yerevan"
	}
	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatal(err)
	}
	time.Local = loc // -> this is setting the global timezone

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	a := &App{
		bot: bot,
	}
	return a
}

func (a *App) Serve(ctx context.Context) {
	// a.LoadDayList()

	tck := time.NewTicker(time.Second)
	defer tck.Stop()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := a.bot.GetUpdatesChan(u)
	for {
		select {
		case <-ctx.Done():
			return
		case <-tck.C:
			if a.currentList == nil || time.Now().After(a.currentList.StopAt) {
				a.NewDayList()
			}
		case update := <-updates:
			msg := update.Message
			if msg != nil {
				if msg.IsCommand() {
					switch msg.Command() {
					case "about":
						a.ReplyToMsg(msg, `This bot supports people queues.
Send your full name to queue up, or send /list command to see full queue for today.`)
					case "list":
						if a.currentList == nil {
							a.SendToChat(msg.Chat.ID, "Queue is not open yet.")
						} else if time.Now().Before(a.currentList.StartedAt) {
							a.SendToChat(msg.Chat.ID, "Queue is not open yet. Wait until "+a.currentList.StartedAt.Format("2 January 15:04 MST"))
						} else {
							a.SendToChat(msg.Chat.ID, a.currentList.String())
						}
					}
				} else {
					if a.currentList == nil {
						a.SendToChat(msg.Chat.ID, "Queue is not open yet.")
					} else if time.Now().Before(a.currentList.StartedAt) {
						a.SendToChat(msg.Chat.ID, "Queue is not open yet. Wait until "+a.currentList.StartedAt.Format("2 January 15:04 MST"))
					} else {
						if err := a.currentList.TailAddPeople(msg.Text); err != nil {
							a.SendToChat(msg.Chat.ID, fmt.Sprintf("Error: %s. Try again.", err))
						} else {
							a.SendToChat(msg.Chat.ID, "Successfully added to the queue! You can view the queue with the /list command.")
						}
					}
				}
			}
		}
	}
}
