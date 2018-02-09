package main

import (
	"log"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func messageHandler(msg *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	log.Printf("===== I got a Message!! =====")
	log.Printf("[%s] %s", msg.From.UserName, msg.Text)

	reply := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
	reply.ReplyToMessageID = msg.MessageID

	bot.Send(reply)
}

func inlineQueryHandler(inlineQuery *tgbotapi.InlineQuery, bot *tgbotapi.BotAPI) {
	log.Printf("===== I got an Inline Query!! =====")

	reply := tgbotapi.NewInlineQueryResultPhoto(inlineQuery.ID, "http://i1.kym-cdn.com/photos/images/original/001/136/185/604.jpg")
	reply.Description = inlineQuery.Query
	reply.ThumbURL = "http://i1.kym-cdn.com/photos/images/original/001/136/185/604.jpg"
	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: inlineQuery.ID,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       []interface{}{reply},
	}
	if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
		log.Println(err)
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI("475294998:AAFom93-YXbuv8ll1_43EWuYar9cR3Xs4sI")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message != nil {
			messageHandler(update.Message, bot)
		} else if update.InlineQuery != nil {
			inlineQueryHandler(update.InlineQuery, bot)
		}
	}
}
