package main

import (
	"flag"
	"fmt"
	config "github.com/toukat/toukabot-v2/config"
	"github.com/toukat/toukabot-v2/util/logger"

	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
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
		log.Fatal(fmt.Sprint("Unable to open config file at %s", *configLocation))
		os.Exit(-1)
	}

	c := config.Config{}
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Unable to parse config file")
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