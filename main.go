package main

import (
	"fmt"
	"generator-super-power-bot/config"
	"generator-super-power-bot/consts"
	"generator-super-power-bot/power"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"math/rand"
	"runtime/debug"
	"time"
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

	powerCache, err := power.NewPowersCache()
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	go powerCache.Update()

	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// init chanel for input updates from API

	updCfg := tgbotapi.NewUpdate(0)
	updCfg.Timeout = 60
	updChan, err := bot.GetUpdatesChan(updCfg)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updChan {
		repl := ""
		userName := update.Message.From.UserName
		chatID := update.Message.Chat.ID
		command := update.Message.Command()
		//text := update.Message.Text
		log.Println(command)
		switch command {
		case "generate":
			randPower := powerCache.GetRandomPower()
			repl = fmt.Sprintf("%s your power - %s\n%s", userName, randPower.PowerName, randPower.Description)
		default:
			repl = "Please use command /generate for generate your super power"
		}

		msg := tgbotapi.NewMessage(chatID, repl)

		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Not sended: %v", err)
		}

	}
}
