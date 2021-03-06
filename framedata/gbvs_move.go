package framedata

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/shurcooL/graphql"

	"github.com/toukat/toukabot-v2/config"
	"github.com/toukat/toukabot-v2/util"
)

type GBVSMove struct {
	Image    string `json:"image"`
	Name     string `json:"name"`
	Input    string `json:"input"`
	Damage   string `json:"damage"`
	Guard    string `json:"guard"`
	Startup  string `json:"startup"`
	Active   string `json:"active"`
	Recovery string `json:"recovery"`
	OnBlock  string `json:"onBlock"`
	OnHit    string `json:"onHit"`
}

const GBVSQuery = "gbvsMove"

func GetGBVSMove(character string, input string) (GBVSMove, error) {
	log.Info(fmt.Sprintf("Making API request to get move data for GBVS, character=%s, input=%s", character,
		input))

	c := config.GetConfig()
	cl := graphql.NewClient(c.APIHost + "/graphql", nil)

	var query struct {
		GbvsMove struct {
			Image    string
			Name     string
			Input    string
			Damage   string
			Guard    string
			Startup  string
			Active   string
			Recovery string
			OnBlock  string
			OnHit    string
		} `graphql:"gbvsMove(characterName: $character, input: $input)"`
	}

	v := map[string]interface{}{
		"character": graphql.String(character),
		"input": graphql.String(input),
	}

	err := cl.Query(context.TODO(), &query, v)
	if err != nil {
		log.Error(err)
		return GBVSMove{}, err
	}

	return GBVSMove{
		Image: query.GbvsMove.Image,
		Name: query.GbvsMove.Name,
		Input: query.GbvsMove.Input,
		Damage: query.GbvsMove.Damage,
		Guard: query.GbvsMove.Guard,
		Startup: query.GbvsMove.Startup,
		Active: query.GbvsMove.Active,
		Recovery: query.GbvsMove.Recovery,
		OnBlock: query.GbvsMove.OnBlock,
		OnHit: query.GbvsMove.OnHit,
	}, err
}

func (m GBVSMove) SendAsEmbed(session *discordgo.Session, message *discordgo.MessageCreate) error {
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
	embed.AddField("On Hit", m.OnHit, true)

	_, err := session.ChannelMessageSendEmbed(message.ChannelID, embed.MessageEmbed)
	if err != nil {
		log.Error("Unable to send move information")
		log.Error(err)
		return err
	}

	return nil
}
