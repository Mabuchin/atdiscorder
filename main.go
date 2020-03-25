package main

import (
	"contest-daily-bot/pkg/collector"
	"contest-daily-bot/pkg/model"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"strings"
	"time"
)

var (
	Token   = fmt.Sprintf("Bot %v", os.Getenv("DISCORD_BOT_TOKEN"))
	BotName = os.Getenv("DISCORD_BOT_NAME")
	RoomId  = os.Getenv("DISCORD_SEND_ROOM_ID")
)

func main() {
	db := model.InitDB()
	defer db.Close()

	// Initialize database data
	data := collector.CollectProblems()
	model.AddProblemList(data)

	// start discord
	discord, err := discordgo.New()
	discord.Token = Token
	if err != nil {
		log.Fatal(err)
	}
	discord.AddHandler(onMessageCreate)
	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening...")
	problem := model.GetRandomProblemData()
	sendText := fmt.Sprintf("今日やるべき問題はこれだ！！\n :ballot_box_with_check: %s \n :link: %s\n", problem.Title, problem.Url)
	_, err = discord.ChannelMessageSend(RoomId, sendText)
	if err != nil{
		fmt.Printf(err.Error())
		log.Fatal(err)
	}
	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}
	log.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)

	switch {
	case strings.HasPrefix(m.Content, fmt.Sprintf("%s %s", BotName, "hi")):
		guildChannels, _ := s.GuildChannels(c.GuildID)
		var sendText string
		for _, a := range guildChannels {
			sendText += fmt.Sprintf("%vチャンネルの%v(IDは%v)\n", a.Type, a.Name, a.ID)
		}
		sendMessage(s, c, sendText)
	}
}

func sendMessage(s *discordgo.Session, c *discordgo.Channel, msg string) {
	_, err := s.ChannelMessageSend(c.ID, msg)
	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}
