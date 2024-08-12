package discord

import (
	"io"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/sullyh7/local-storage/config"
)

var Ses *discordgo.Session

func InitSession(conf config.AppConfig) error {
	session, err := discordgo.New("Bot " + conf.DiscordBotToken)
	if err != nil {
		log.Fatalf("error connecting to discord: %s\n", err.Error())
	}
	Ses = session
	return nil
}

func UploadFile(channelId, name string, file io.Reader) error {
	if err := Ses.Open(); err != nil {
		return err
	}
	defer Ses.Close()
	_, err := Ses.ChannelFileSend(config.Config.DiscordChannelID, name, file)
	if err != nil {
		return err
	}
	return nil
}

func GetFiles(channelId string) ([]*discordgo.Message, error) {
	if err := Ses.Open(); err != nil {
		return nil, err
	}
	defer Ses.Close()
	msgs, err := Ses.ChannelMessages(channelId, 20, "", "", "")
	if err != nil {
		return nil, err
	}
	filtered := make([]*discordgo.Message, 0)
	for _, m := range msgs {
		if len(m.Attachments) > 0 {
			filtered = append(filtered, m)
		}
	}
	return filtered, nil
}

func DeleteFile(channelID, msgID string) error {
	if err := Ses.Open(); err != nil {
		return err
	}
	defer Ses.Close()
	if err := Ses.ChannelMessageDelete(channelID, msgID); err != nil {
		return err
	}
	return nil
}
