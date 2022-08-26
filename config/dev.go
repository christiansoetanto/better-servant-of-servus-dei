package config

import (
	"github.com/christiansoetanto/better-servant-of-servus-dei/config/configtypes"
)

func getDevConfig() configtypes.GuildConfig {
	return configtypes.GuildConfig{
		Guild: configtypes.Guild{
			GuildName: "Local Server",
			GuildID:   "813302330782253066",
		},
		Channel: configtypes.Channel{
			GeneralDiscussion:            "813302330782253069",
			ReactionRoles:                "941213323444244501",
			ServerInformation:            "848858055944306698",
			ReligiousQuestions:           "",
			ReligiousDiscussions1:        "",
			ReligiousDiscussions2:        "",
			AnsweredQuestions:            "",
			FAQ:                          "",
			Responses:                    "",
			VettingQuestioning:           "",
			RulesVetting:                 "",
			LiturgicalCalendarDiscussion: "813302330782253069",
		},
		Role: configtypes.Role{
			Vetting:            "974632148952809482",
			VettingQuestioning: "974632188823863296",
			ApprovedUser:       "974632216304943155",
			LatinCatholic:      "974630535395680337",
			EasternCatholic:    "974667212587671613",
			OrthodoxChristian:  "974667248826449950",
			RCIACatechumen:     "974667251498225704",
			Protestant:         "974667253045919784",
			NonCatholic:        "974667254627201084",
			Atheist:            "974667257122795570",
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
			Upvote: "",
		},
		Wording: configtypes.Wording{
			AcknowledgementMessageFormat: "Verification of user <@%s> with role <@&%s> is successful.\nThank you for using my service. Beep. Boop.\n",
			WelcomeMessageEmbedFormat:    "Welcome to Servus Dei, <@%s>! We are happy to have you! Make sure you check out <#%s> to gain access to the various channels we offer and please do visit <#%s> so you can understand our server better and take use of everything we have to offer. God Bless!",
			WelcomeMessageFormat:         "Hey <@%s>! It looks like you missed question 1. Please re-read the <#%s> again, we assure you that the code is in there. Thank you for your understanding.\\nPS: if you are sure you got it right, please ignore this message.",
			MissedQuestionOneFormat:      "Welcome to Servus Dei!",
			WelcomeTitle:                 "Hey %s! %s just approved your vetting response. Welcome to the server. Feel free to tag us should you have further questions. Enjoy!",
		},
	}

}
