package configtypes

type GuildConfig struct {
	Guild                  Guild
	Channel                Channel
	Role                   Role
	Moderator              Moderator
	Reaction               Reaction
	Wording                Wording
	ReligionRoleMappingMap ReligionRoleMappingMap
}

type Guild struct {
	GuildID   string
	GuildName string
}
