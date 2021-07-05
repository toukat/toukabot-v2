package util

import (
	"fmt"
	"reflect"

	"github.com/bwmarrin/discordgo"
)

type Embed struct {
	*discordgo.MessageEmbed
}

func NewEmbed() *Embed {
	return &Embed{&discordgo.MessageEmbed{}}
}

func (e *Embed) SetImage(image string) *Embed {
	e.Image = &discordgo.MessageEmbedImage{URL: image}
	return e
}

func (e *Embed) SetTitle(title string) *Embed {
	e.Title = title
	return e
}

func (e *Embed) AddField(name string, value interface{}, inline bool) *Embed {
	if value == reflect.Zero(reflect.TypeOf(value)).Interface() {
		return e
	}

	e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
		Name: name,
		Value: fmt.Sprintf("%v", value),
		Inline: inline,
	})

	return e
}

func (e *Embed) AddBlankField(inline bool) *Embed {
	e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
		Name: "​​​\u200B",
		Value: "\u200B",
		Inline: inline,
	})

	return e
}