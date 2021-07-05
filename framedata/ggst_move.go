package framedata

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/shurcooL/graphql"

	"github.com/toukat/toukabot-v2/config"
	"github.com/toukat/toukabot-v2/util"
)

type GGSTMove struct {
	Image    string `json:"image"`
	Name     string `json:"name"`
	Input    string `json:"input"`
	Damage   string `json:"damage"`
	Guard    string `json:"guard"`
	Startup  string `json:"startup"`
	Active   string `json:"active"`
	Recovery string `json:"recovery"`
	OnBlock  string `json:"onBlock"`
	Notes    string `json:"notes"`
}

const GGSTQuery = "ggstMove"

func GetGGSTMove(character string, input string) (GGSTMove, error) {
	log.Info(fmt.Sprintf("Making API request to get move data for GGST, character=%s, input=%s", character,
		input))

	c := config.GetConfig()
	cl := graphql.NewClient(c.APIHost + "/graphql", nil)

	var query struct{
		GgstMove struct {
			Image    string
			Name     string
			Input    string
			Damage   string
			Guard    string
			Startup  string
			Active   string
			Recovery string
			OnBlock  string
			Notes    string
		} `graphql:"ggstMove(characterName: $character, input: $input)"`
	}

	v := map[string]interface{}{
		"character": graphql.String(character),
		"input": graphql.String(input),
	}

	err := cl.Query(context.TODO(), &query, v)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to get move information, game=ggst, character=%s, input=%s", character,
			input))
		log.Error(err)
		return GGSTMove{}, err
	}

	return GGSTMove{
		Image: query.GgstMove.Image,
		Name: query.GgstMove.Name,
		Input: query.GgstMove.Input,
		Damage: query.GgstMove.Damage,
		Guard: query.GgstMove.Guard,
		Startup: query.GgstMove.Startup,
		Active: query.GgstMove.Active,
		Recovery: query.GgstMove.Recovery,
		OnBlock: query.GgstMove.OnBlock,
		Notes: query.GgstMove.Notes,
	}, err
}

func (m GGSTMove) SendAsEmbed(session *discordgo.Session, message *discordgo.MessageCreate) error {
	embed := util.NewEmbed()
	embed.SetImage(m.Image)
	embed.SetTitle(m.Name)
	embed.AddField("Input", m.Input, true)
	embed.AddBlankField(true)
	embed.AddBlankField(true)
	embed.AddField("Damage", m.Damage, true)
	embed.AddField("Guard", m.Guard, true)
	embed.AddBlankField(true)
	embed.AddField("Startup", m.Startup, true)
	embed.AddField("Active", m.Active, true)
	embed.AddField("Recovery", m.Recovery, true)
	embed.AddField("On Block", m.OnBlock, true)
	embed.AddField("Notes", m.Notes, true)

	_, err := session.ChannelMessageSendEmbed(message.ChannelID, embed.MessageEmbed)
	if err != nil {
		log.Error("Unable to send move information")
		log.Error(err)
		return err
	}

	return nil
}
