package configtypes

type Channel struct {
	GeneralDiscussion            ChannelID
	ReactionRoles                ChannelID
	ServerInformation            ChannelID
	ReligiousQuestions           ChannelID
	ReligiousDiscussions1        ChannelID
	ReligiousDiscussions2        ChannelID
	AnsweredQuestions            ChannelID
	FAQ                          ChannelID
	Responses                    ChannelID
	VettingQuestioning           ChannelID
	RulesVetting                 ChannelID
	LiturgicalCalendarDiscussion ChannelID
}

type ChannelID string
