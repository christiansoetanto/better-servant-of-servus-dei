package util

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
)

const (
	WelcomeTitle        = "Welcome to Servus Dei!"
	LogoURL             = "https://cdn.discordapp.com/avatars/767426889294938112/0e100e9fec18866892ed0c875b341926.png"
	Author              = "Servant of Servus Dei"
	WelcomeImageURL     = "https://cdn.discordapp.com/attachments/751174152588623912/976921809607880714/You_Doodle_2022-05-19T18_58_15Z.jpg"
	WelcomeImage2URL    = "https://media.discordapp.net/attachments/751174152588623912/975368929008558130/Screenshot_2022-05-11_at_11.42.51_PM.png"
	FooterText          = "2022 | Made for Servus Dei by soetanto™\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000"
	ServusDeiWebsiteURL = "https://www.servusdeicatholic.com/"
	GoldenYellowColor   = 16769280
)

type ImageUrl string

func RandomWelcomeImage() string {
	in := []string{WelcomeImageURL, WelcomeImage2URL}
	return in[rand.Intn(len(in))]
}
func EmbedBuilder(title string, description string, param ...interface{}) *discordgo.MessageEmbed {
	var imageUrl string
	var fields []*discordgo.MessageEmbedField
	for _, p := range param {
		switch v := p.(type) {
		case ImageUrl:
			imageUrl = string(v)
		case []*discordgo.MessageEmbedField:
			fields = v
		}

	}
	return &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       title,
		Description: description,
		Color:       GoldenYellowColor,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: LogoURL,
		},
		Image: &discordgo.MessageEmbedImage{
			URL: imageUrl,
		},
		URL:    ServusDeiWebsiteURL,
		Fields: fields,
	}
}
