package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/toukat/toukabot-v2/commands"
	"github.com/toukat/toukabot-v2/config"
	"github.com/toukat/toukabot-v2/twitter"
	"github.com/toukat/toukabot-v2/util"
)

func ParseMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == uid {
		return
	}

	c := config.GetConfig()
	splitMessage := strings.Split(message.Content, " ")

	// Check for Twitter links
	for _, m := range splitMessage {
		if util.URLValid(m) && util.URLAvailable(m) {
			twitter.ParseTwitterLink(session, message.ChannelID, m)
		}
	}

	if len(message.Content) < 1 || message.Content[0:1] != c.CommandPrefix {
		return
	}

	command := splitMessage[0][1:]

	if command == "frame" {
		commands.FrameData(session, message)
	}
}