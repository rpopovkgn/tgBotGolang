package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.Text = "Привет, я бот помощник. Я умею показывать баланс смс и звонков! Команда для смс: balance Команда для звонков: spacetel"
				bot.Send(msg)
			case "/balance":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.Text, _ = GetBalanceQtelecom()
				bot.Send(msg)
			case "/spacetel":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.Text, _ = GetBalanceSpacetel()
				bot.Send(msg)
			}
		}
	}
}
