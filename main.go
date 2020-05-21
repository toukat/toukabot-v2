package main

import (
	"flag"
	"fmt"
	"github.com/toukat/toukabot-v2/config"
	"github.com/toukat/toukabot-v2/util/logger"

	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	d *discordgo.Session
	uid string
	log logger.Logger
)

func onReady(s *discordgo.Session, e *discordgo.Ready) {
	log.Info("ToukaBot started")
	err := s.UpdateStatus(0, "test")

	if err != nil {
		log.Error(err)
		log.Error("Unable to change status")
	}

	uid = e.User.ID
}

func main() {
	log = logger.GetLogger("ToukaBot V2 Main")

	configLocation := flag.String("config", "./config.yml", "Configuration file")
	flag.Parse()

	configFile, err := os.Open(*configLocation)
	if err != nil {
		log.Fatal(err)
		log.Fatal(fmt.Sprintf("Unable to open config file at %s", *configLocation))
		os.Exit(-1)
	}

	c, err := config.CreateConfig(configFile)
	if err != nil {
		log.Fatal("Error parsing config file")
		os.Exit(-1)
	}

	log.Info("Starting Discord session...")
	d, err = discordgo.New(c.BotToken)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Failed to create Discord session")
		os.Exit(-1)
	}

	d.AddHandler(onReady)

	err = d.Open()
	if err != nil {
		log.Fatal(err)
		log.Fatal("Failed to create Discord websocket connection")
		os.Exit(-1)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, os.Kill)
	<-s

	return
}