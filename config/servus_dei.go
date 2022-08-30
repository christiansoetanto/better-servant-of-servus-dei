package config

import (
	"github.com/christiansoetanto/better-servant-of-servus-dei/config/configtypes"
)

func getServusDeiGuildConfig() configtypes.GuildConfig {
	return configtypes.GuildConfig{
		Guild: configtypes.Guild{
			GuildID:   "751139261515825162",
			GuildName: "Servus Dei",
		},
		Channel: configtypes.Channel{
			GeneralDiscussion:            "751174152588623912",
			ReactionRoles:                "767452241321000970",
			ServerInformation:            "973586981789499452",
			ReligiousQuestions:           "751174501307383908",
			ReligiousDiscussions1:        "751174442217898065",
			ReligiousDiscussions2:        "771836244879605811",
			AnsweredQuestions:            "821657995129126942",
			FAQ:                          "806007417321291776",
			Responses:                    "751151421231202363",
			VettingQuestioning:           "914987511481249792",
			RulesVetting:                 "775654889934159893",
			LiturgicalCalendarDiscussion: "915621423270211594",
			BotTesting:                   "929367689057693757",
		},
		Role: configtypes.Role{
			Vetting:            "751145124834312342",
			VettingQuestioning: "914986915030241301",
			ApprovedUser:       "751144797938384979",
			LatinCatholic:      "751145824532168775",
			EasternCatholic:    "751148911267414067",
			OrthodoxChristian:  "751148354716565656",
			RCIACatechumen:     "751196794771472395",
			Protestant:         "751145951137103872",
			NonCatholic:        "751146099351224382",
			Atheist:            "751148904938209351",
			Moderator:          "751144316843327599",
		},
		Moderator: configtypes.Moderator{
			"255514888041005057": "soetanto",
			"385901039171272726": "cathmeme",
			"505100307051708416": "potato",
			"339808311153000460": "shadowfax",
			"328369198696890378": "hick",
			"201126729564028928": "gio",
			"650493923357229091": "zech",
			"633204791610179584": "chaos",
			"469970745586483210": "hermano",
			"761486036987281438": "trex",
			"302301261103890452": "braydog",
			"401949214373838848": "buggy",
			"274724929478459392": "carlos",
			"501368090504986627": "athanasius",
		},
		Reaction: configtypes.Reaction{
			Upvote: "762045856592822342",
		},
		Wording: configtypes.Wording{
			AcknowledgementMessageFormat: "Verification of user <@%s> with role <@&%s> is successful.\nThank you for using my service. Beep. Boop.\n",
			WelcomeMessageEmbedFormat:    "Welcome to Servus Dei, <@%s>! We are happy to have you! Make sure you check out <#%s> to gain access to the various channels we offer and please do visit <#%s> so you can understand our server better and take use of everything we have to offer. God Bless!",
			MissedQuestionOneFormat:      "Hey <@%s>! It looks like you missed question 1. Please re-read the <#%s> again, we assure you that the code is in there. Thank you for your understanding.",
			WelcomeTitle:                 "Welcome to Servus Dei!",
			WelcomeMessageFormat:         "Hey %s! %s just approved your vetting response. Welcome to the server. Feel free to tag us should you have further questions. Enjoy!",
		},
	}
}
