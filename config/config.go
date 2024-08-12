package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Config AppConfig

type AppConfig struct {
	DiscordAppID, DiscordPublicKey, DiscordBotToken, DiscordChannelID string
}

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	Config = AppConfig{
		DiscordAppID:     getVal("DISCORD_APP_ID", ""),
		DiscordPublicKey: getVal("DISCORD_PUBLIC_KEY", ""),
		DiscordBotToken:  getVal("DISCORD_BOT_TOKEN", ""),
		DiscordChannelID: getVal("DISCORD_CHANNEL_ID", ""),
	}
	return nil
}

func getVal(key, alt string) string {
	val := os.Getenv(key)
	if val == "" {
		return alt
	}
	return val
}
