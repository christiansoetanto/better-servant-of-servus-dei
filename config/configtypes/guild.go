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
	GuildID   GuildID
	GuildName GuildName
}
type GuildID string
type GuildName string
