package twitter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/toukat/toukabot-v2/util"
	"github.com/toukat/toukabot-v2/util/logger"
)

var log = logger.GetLogger("TwitterParser")

/*
 * Function: ParseTwitterLink
 * Get GIFs and videos from Twitter links
 *
 * Params:
 * url: Twitter URL to get content from
 *
 * Returns:
 * URL to GIF or video if there is one
 * error if anything went wrong
 */
func ParseTwitterLink(session *discordgo.Session, channelID string, url string) {
	log.Info(fmt.Sprintf("Checking URL to see if it's a valid URL, url=%s", url))
	available := util.URLAvailable(url)

	if !available {
		log.Error(fmt.Sprintf("URL not available, url=%s", url))
		return
	}

	twitterClient := TwitterAuth()

	splitMessage := strings.Split(url, "/")
	var id int64 = -1
	for i, v := range splitMessage {
		if i > 0 && splitMessage[i - 1] == "status" {
			splitId := strings.Split(v, "?")
			temp, err := strconv.ParseInt(splitId[0], 10, 64)
			if err != nil {
				log.Error(fmt.Sprintf("Unable to parse Tweet ID, err=%s", err))
				return
			}

			id = temp
		}
	}

	if id < 0 {
		log.Info(fmt.Sprintf("Link did not contain valid Twitter ID, url=%s", url))
		return
	}

	tweet, _, err := twitterClient.Statuses.Show(id, nil)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to fetch Tweet, id=%d, err=%s", id, err))
		return
	}

	if tweet.ExtendedEntities == nil || len(tweet.ExtendedEntities.Media) < 1 {
		log.Info(fmt.Sprintf("Tweet has no media or animated media, id=%d", id))
		return
	}

	media :=  tweet.ExtendedEntities.Media[0].Type

	if media == "video" {
		log.Info(fmt.Sprintf("Tweet has video, id=%d", id))
		_, err = session.ChannelMessageSend(channelID, tweet.ExtendedEntities.Media[0].VideoInfo.Variants[0].URL)
		if err!= nil {
			log.Error("Unable to send content")
			log.Error(err)
		}
	} else if media == "animated_gif" {
		log.Info(fmt.Sprintf("Tweet has GIF, id=%d", id))
		_, err = session.ChannelMessageSend(channelID, tweet.ExtendedEntities.Media[0].VideoInfo.Variants[0].URL)
		if err != nil {
			log.Error("Unable to send content")
			log.Error(err)
		}
	} else {
		log.Info(fmt.Sprintf("Tweet doesn't have GIF or video, id=%d", id))
	}
}
