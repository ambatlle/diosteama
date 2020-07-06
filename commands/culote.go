package commands

import (
	"fmt"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func culote(update tgbotapi.Update, bot *tgbotapi.BotAPI, argv []string) {
	text := "%s, tienes un culote como para meter %s"

	reply := fmt.Sprintf(text, update.Message.From.FirstName, strings.Join(argv, " "))
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
	bot.Send(msg)
}
