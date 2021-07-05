package framedata

import (
	"context"
	"fmt"
	"github.com/toukat/toukabot-v2/config"
	"github.com/shurcooL/graphql"
	// "github.com/toukat/toukabot-v2/graphql"
	"github.com/toukat/toukabot-v2/util"
	"github.com/toukat/toukabot-v2/util/logger"

	"github.com/bwmarrin/discordgo"
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

const FrameDataEndpoint = "/api/v1/moves/%s/%s/%s"
const GBVSQuery = "gbvsMove"
//const GBVSImage = "image"
//const GBVSName = "name"
//const GBVSInput = "input"
//const GBVSDamage = "damage"
//const GBVSGuard = "guard"
//const GBVSStartup = "startup"
//const GBVSActive = "active"
//const GBVSRecovery = "recovery"
//const GBVSOnBlock = "onBlock"
//const GBVS

var log = logger.GetLogger("FrameData")

func GetGBVSMove(character string, input string) (GBVSMove, error) {
	log.Info(fmt.Sprintf("Making API request to get move data for GBVS, character=%s, input=%s", character,
		input))

	//query := fmt.Sprintf(`{%s(%s: "%s", %s: "%s"){`, GBVSQuery, CharacterName, character, Input, input)
	//fields := util.GetStructFields(GBVSMove{})
	//for _, v := range fields {
	//	query += fmt.Sprintf("%s\n", string(unicode.ToLower(rune(v[0]))) + v[1:])
	//}
	//query += "}}"

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
	//resp, err := graphql.Query(query, nil, &m)
	//if err != nil {
	//	log.Error(fmt.Sprintf("Unable to get move data for GBVS, character=%s, input=%s", character, input))
	//	return GBVSMove{}, err
	//}
	//
	//log.Info(resp)

	err := cl.Query(context.Background(), &query, v)
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

	//resp, err := request.GetRequest(fmt.Sprintf(FrameDataEndpoint, game, character, input))
	//if err != nil {
	//	log.Error(fmt.Sprintf("Unable to get move data for GBVS, character=%s, input=%s", character, input))
	//	return GBVSMove{}, err
	//}
	//
	//var move GBVSMove
	//decoder := json.NewDecoder(resp)
	//err = decoder.Decode(&move)
	//if err != nil {
	//	log.Error("Error decoding API response")
	//	return GBVSMove{}, err
	//}
	//
	//return move, nil
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
