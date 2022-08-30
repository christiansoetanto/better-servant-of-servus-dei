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
			GeneralDiscussion:            "1013780724345745508",
			ReactionRoles:                "1013780802619854848",
			ServerInformation:            "1013780836203638836",
			ReligiousQuestions:           "1013780754834145333",
			ReligiousDiscussions1:        "1013780733510287472",
			ReligiousDiscussions2:        "1013780741542379620",
			AnsweredQuestions:            "1013780765307310091",
			FAQ:                          "1013780853282844672",
			Responses:                    "1013780662798528592",
			VettingQuestioning:           "1013780704330526834",
			RulesVetting:                 "1013780880591954002",
			LiturgicalCalendarDiscussion: "1013780907192221757",
			BotTesting:                   "1013780949026230292",
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
			Moderator:          "1013781460953616404",
		},
		Moderator: configtypes.Moderator{
			"255514888041005057": "soetanto",
		},
		Reaction: configtypes.Reaction{
			Upvote:                               "1013782200052887683",
			Dab:                                  "<:Upvote:1013782200052887683>",
			ReligiousDiscussionOneWhiteCheckmark: "✅",
			ReligiousDiscussionsTwoBallotBoxWithCheck: "☑️",
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
