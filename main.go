package main

import (
	"log"
	"strings"
	"gopkg.in/telegram-bot-api.v4"
	"github.com/peterhellberg/flip"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("314464484:AAE9UtsTGalSgaTa6rVJ51dJmXaToaUBLdk")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		reply := ""

		if strings.HasPrefix(update.Message.Text, "/about ") {
			reply = "This is a weird bot. Very Alay"
		} else if strings.HasPrefix(update.Message.Text, "/alayed ") {
			reply = update.Message.Text[8:]
			reply = strings.Replace(reply,"a","@",-1)
			reply = strings.Replace(reply,"i","!",-1)
			reply = strings.Replace(reply,"e","3",-1)
			reply = strings.Replace(reply,"o","0",-1)
		}  else if strings.HasPrefix(update.Message.Text, "/pusing ") {
			reply = flip.Table(update.Message.Text[8:])
			reply = reverse(reply)
		} else if strings.HasPrefix(update.Message.Text, "/balik2 ") {
			reply = flip.Table(update.Message.Text[8:])
		} else if strings.HasPrefix(update.Message.Text, "/balik ") {
			reply = reverse(update.Message.Text[7:])
		} else {
			reply = "You can type \n/alayed [text]\n/balik [text]\n/balik2 [text]\n/pusing [text]\n/about"
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

