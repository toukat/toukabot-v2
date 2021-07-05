package commands

import (
	"github.com/toukat/toukabot-v2/framedata"
	"github.com/toukat/toukabot-v2/util/logger"

	"strings"

	"github.com/bwmarrin/discordgo"
)

var log = logger.GetLogger("Frame Data")

func FrameData(session *discordgo.Session, message *discordgo.MessageCreate) {
	splitMessage := strings.SplitN(message.Content, " ", 4)

	if strings.ToLower(splitMessage[1]) != "gbvs" {
		log.Error("Invalid game")
		return
	}

	move, err := framedata.GetGBVSMove(splitMessage[2], splitMessage[3])
	if err != nil {
		log.Error("Error getting move")
		log.Error(err)

		_, err = session.ChannelMessageSend(message.ChannelID, "Unable to fetch move information.")
		if err != nil {
			log.Error("Unable to send error message")
			log.Error(err)
			return
		}

		return
	}

	err = move.SendAsEmbed(session, message)
	if err != nil {
		log.Error("Unable to send message")
		return
	}
}