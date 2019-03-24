package main

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"runtime/debug"
	"generator-super-power-bot/config"
	"generator-super-power-bot/consts"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			debug.PrintStack()
			log.Println(r)
		}
	}()

	cfg, err := config.NewConfig(consts.CONFIG_PATH)
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// init chanel for input updates from API

	updCfg := tgbotapi.NewUpdate(0)
	updCfg.Timeout = 200
	updChan, err := bot.GetUpdatesChan(updCfg)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updChan {
		userName := update.Message.From.UserName
		chatID := update.Message.Chat.ID
		command := update.Message.Command()
		text := update.Message.Text
		log.Printf(
			"UserName: %s, chatID: %s, command: %s, text: %s",
			userName,
			chatID,
			command,
			text,
		)

		repl := fmt.Sprintf("command:%s, text: %s", command, text)

		msg := tgbotapi.NewMessage(chatID, repl)

		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Not sended: %v", err)
		}

	}
}
