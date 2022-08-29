package configtypes

type GuildConfig struct {
	Guild     Guild
	Channel   Channel
	Role      Role
	Moderator Moderator
	Reaction  Reaction
	Wording   Wording
}

type Guild struct {
	GuildID   string
	GuildName string
}
