package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Telegram struct {
		APItoken string `yaml:"api_token"`
		ChatID   int64  `yaml:"chat_id"`
	} `yaml:"telegram"`
	DB struct {
		Type     string `yaml:"type"`
		Path     string `yaml:"path"`
		InitSQL  string `yaml:"init_sql"`
		AlterSQL string `yaml:"alter_sql"`
	} `yaml:"db"`
	Discord struct {
		Token        string
		ARPChannelID string
		PSChannelID  string
		WPSChannelID string
	}
}

func InitConfig(filepath string) *Config {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Could not obtain file from %s. Error: %s", filepath, err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Could not unmarshal %s. Error: %s", filepath, err)
	}
	if cfg.Discord.Token = os.Getenv("DISCORD_BOT"); cfg.Discord.Token == "" {
		log.Fatalf("Could not obtain discord token")
	}
	if cfg.Discord.ARPChannelID = os.Getenv("ARP_CHANNEL_ID"); cfg.Discord.ARPChannelID == "" {
		log.Println("Could not obtain ARP channel ID")
	} else {
		log.Println("Obtained ARP channel ID:", cfg.Discord.ARPChannelID)
	}
	if cfg.Discord.PSChannelID = os.Getenv("PS_CHANNEL_ID"); cfg.Discord.PSChannelID == "" {
		log.Println("Could not obtain PS channel ID")
	} else {
		log.Println("Obtained PS channel ID:", cfg.Discord.ARPChannelID)
	}
	if cfg.Discord.WPSChannelID = os.Getenv("WPS_CHANNEL_ID"); cfg.Discord.WPSChannelID == "" {
		log.Println("Could not obtain WPS channel ID")
	} else {
		log.Println("Obtained WPS channel ID:", cfg.Discord.ARPChannelID)
	}
	return &cfg
}
