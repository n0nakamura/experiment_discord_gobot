package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

func main() {
	// config.json の読み取り
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	// fileの値を構造体Config型cfgに格納
	var cfg Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 新しいDiscordGoセッションの作製
	dg, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Fatal(err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

// メッセージ受信時のイベントハンドラ
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// メッセージが自分自身のものでない場合のみ処理を実行
	if m.Author.ID != s.State.User.ID {
		// メッセージ送信
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hello World")
	}
}
