package format

import (
	"fmt"
	"html"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pandacrew-net/diosteama/database"
	"github.com/pandacrew-net/diosteama/quotes"
)

func parseTime(t string) time.Time {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		fmt.Println(err)
	}
	i, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		i = 1
	}
	tm := time.Unix(i, 0).In(loc)
	return tm
}

// Quote formats a quote to be delivered to the chat
func Quote(quote quotes.Quote) string {
	nick := strings.SplitN(quote.Author, "!", 2)[0]
	//💩🔞🔪💥
	formatted := fmt.Sprintf("<pre>%s</pre>\n\n<em>🚽 Quote %d by %s on %s</em>",
		html.EscapeString(quote.Text), quote.Recnum, html.EscapeString(nick), parseTime(quote.Date))
	return formatted
}

// RawQuote creates a string out from a list of raw quotes
func RawQuote(msgs []*tgbotapi.Message) string {
	var result string
	for i := range msgs {
		result = result + RawQuoteMessage(msgs[i])
	}
	return result
}

// RawQuoteMessage creates author: text from a raw message
func RawQuoteMessage(msg *tgbotapi.Message) string {
	var user *tgbotapi.User
	var name, text string
	if msg.ReplyToMessage != nil {
		user = msg.ReplyToMessage.From
		text = msg.ReplyToMessage.Text
	} else {
		user = msg.ForwardFrom
		text = msg.Text
	}

	name, err := database.NickFromTGUser(user)
	if err != nil {
		name = user.String()
	}

	return fmt.Sprintf("%s: %s\n", name, text)
}
