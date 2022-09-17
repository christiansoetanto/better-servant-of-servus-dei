package usecase

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/better-servant-of-servus-dei/config/configtypes"
	"github.com/christiansoetanto/better-servant-of-servus-dei/util"
	"log"
)

type commandHandler map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error

const (
	PingCommand          = "ping"
	SDVerifyCommand      = "sdverify"
	SDQuestionOneCommand = "sdquestionone"
	CalendarCommand      = "calendar"
	NiceTryBro           = "Nice try, bro! You are not allowed to use this command... <@255514888041005057>"
)

func (u *usecase) registerSlashCommand() {
	log.Println("Adding commands...")
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        PingCommand,
			Description: "Ping",
		},
		{
			Name:        SDVerifyCommand,
			Description: "Command for verifying new peeps and welcoming them",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user-option",
					Description: "User to verify",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "role-option",
					Description: "Religion role to give",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  string(configtypes.LatinCatholic),
							Value: configtypes.LatinCatholic,
						},
						{
							Name:  string(configtypes.EasternCatholic),
							Value: configtypes.EasternCatholic,
						},
						{
							Name:  string(configtypes.OrthodoxChristian),
							Value: configtypes.OrthodoxChristian,
						},
						{
							Name:  string(configtypes.RCIACatechumen),
							Value: configtypes.RCIACatechumen,
						},
						{
							Name:  string(configtypes.Protestant),
							Value: configtypes.Protestant,
						},
						{
							Name:  string(configtypes.NonCatholic),
							Value: configtypes.NonCatholic,
						},
						{
							Name:  string(configtypes.Atheist),
							Value: configtypes.Atheist,
						},
					},
				},
			},
		},
		{
			Name:        SDQuestionOneCommand,
			Description: "Command for alerting peeps that they missed question one code",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user-option",
					Description: "User to alert",
					Required:    true,
				},
			},
		},
		{
			Name:        CalendarCommand,
			Description: "Get today's liturgical calendar",
		},
	}
	for guildId := range u.Config.GuildConfig {
		_, err := u.Session.ApplicationCommandBulkOverwrite(u.Session.State.User.ID, guildId, commands)
		if err != nil {
			log.Fatalf("Cannot create command: %v", err)
		}
	}
}

func (u *usecase) initCommandHandler() {
	commandHandlers := u.commandHandlerBuilder()

	u.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			err := h(s, i)
			if err != nil {
				u.errorReporter(err)
			}
		}
	})

}

func (u *usecase) commandHandlerBuilder() commandHandler {
	return commandHandler{
		PingCommand:          u.pingCommandFunc,
		SDVerifyCommand:      u.sdVerifyCommandFunc,
		SDQuestionOneCommand: u.sdQuestionOneCommandFunc,
		CalendarCommand:      u.calendarCommandFunc,
	}
}

func (u *usecase) pingCommandFunc(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	guildCfg, ok := u.getGuildConfig(i.GuildID)
	if !ok {
		return nil
	}
	if !u.isMod(i.Member.User.ID, guildCfg) {
		err := u.alertNonMod(i)
		if err != nil {
			return err
		}
		return nil
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *usecase) sdQuestionOneCommandFunc(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	guildCfg, ok := u.getGuildConfig(i.GuildID)
	if !ok {
		return nil
	}
	if !u.isMod(i.Member.User.ID, guildCfg) {
		err := u.alertNonMod(i)
		if err != nil {
			return err
		}
		return nil
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Processing... Please wait...",
		},
	})
	if err != nil {
		return err
	}
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	args := make([]interface{}, 0)

	var user *discordgo.User

	userOpt, userOptOk := optionMap["user-option"]
	if userOptOk {
		user = userOpt.UserValue(s)
		args = append(args, user.ID)
		args = append(args, guildCfg.Channel.RulesVetting)
		err = s.GuildMemberRoleAdd(i.GuildID, user.ID, guildCfg.Role.VettingQuestioning)
		if err != nil {
			return err
		}

		_, err = s.ChannelMessageSendComplex(guildCfg.Channel.VettingQuestioning, &discordgo.MessageSend{
			Content: fmt.Sprintf("<@%s>", user.ID),
			Embed: util.EmbedBuilder("Vetting Police!", fmt.Sprintf(
				guildCfg.Wording.MissedQuestionOneFormat,
				args...,
			)),
		})
		if err != nil {
			return err
		}

		responseEdit := fmt.Sprintf("done please check <#%s>.", guildCfg.Channel.VettingQuestioning)
		_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &responseEdit,
		})

		if err != nil {
			return err
		}
	} else {
		c := "Please select user"
		_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &c,
		})
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}
func (u *usecase) sdVerifyCommandFunc(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	guildCfg, ok := u.getGuildConfig(i.GuildID)
	if !ok {
		return nil
	}
	if !u.isMod(i.Member.User.ID, guildCfg) {
		err := u.alertNonMod(i)
		if err != nil {
			return err
		}
		return nil
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Processing... Please wait...",
		},
	})
	if err != nil {
		return err
	}
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	var user *discordgo.User

	userOpt, userOptOk := optionMap["user-option"]
	roleOpt, roleOptOk := optionMap["role-option"]
	var roleType string
	if userOptOk && roleOptOk {
		user = userOpt.UserValue(s)
		ackMessageArgs := make([]interface{}, 0, len(options))
		ackMessageArgs = append(ackMessageArgs, user.ID)
		welcomeMessageArgs := make([]interface{}, 0)
		welcomeMessageArgs = append(welcomeMessageArgs, user.ID)
		welcomeMessageArgs = append(welcomeMessageArgs, guildCfg.Channel.ReactionRoles)
		welcomeMessageArgs = append(welcomeMessageArgs, guildCfg.Channel.ServerInformation)

		roleType = roleOpt.StringValue()
		roleId := guildCfg.ReligionRoleMappingMap[configtypes.ReligionRoleType(roleType)]
		ackMessageArgs = append(ackMessageArgs, roleId)

		err = s.GuildMemberRoleAdd(i.GuildID, user.ID, string(roleId))
		if err != nil {
			return err
		}
		err = s.GuildMemberRoleAdd(i.GuildID, user.ID, guildCfg.Role.ApprovedUser)
		if err != nil {
			return err
		}
		err = s.GuildMemberRoleRemove(i.GuildID, user.ID, guildCfg.Role.Vetting)
		if err != nil {
			return err
		}
		err = s.GuildMemberRoleRemove(i.GuildID, user.ID, guildCfg.Role.VettingQuestioning)
		if err != nil {
			return err
		}
		mod := i.Member
		content := fmt.Sprintf(guildCfg.Wording.WelcomeMessageFormat, user.Mention(), mod.Mention())
		_, err = s.ChannelMessageSend(guildCfg.Channel.GeneralDiscussion, content)
		if err != nil {
			return err
		}

		_, err = s.ChannelMessageSendEmbed(
			guildCfg.Channel.GeneralDiscussion,
			util.EmbedBuilder(
				guildCfg.Wording.WelcomeTitle,
				fmt.Sprintf(guildCfg.Wording.WelcomeMessageEmbedFormat, welcomeMessageArgs...),
				util.ImageUrl(util.RandomWelcomeImage()),
			),
		)
		if err != nil {
			return err
		}
		emptyString := ""
		_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &emptyString,
			Embeds: &[]*discordgo.MessageEmbed{
				util.EmbedBuilder("Verify Police!", fmt.Sprintf(
					guildCfg.Wording.VerifyAckMessageFormat,
					ackMessageArgs...,
				)),
			},
		})

		if err != nil {
			return err
		}
	} else {
		c := "Please choose user and role."
		_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &c,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *usecase) calendarCommandFunc(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Processing... Please wait...",
		},
	})
	if err != nil {
		return err
	}

	embed, err := u.generateCalendarEmbed()
	if err != nil {
		return err
	}
	emptyString := ""
	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{
			embed,
		},
		Content: &emptyString,
	})

	if err != nil {
		return err
	}

	log.Printf("[%s] : [%s] ", CalendarCommand, i.Interaction.Member.User.Username)
	return nil
}
