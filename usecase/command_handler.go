package usecase

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

type commandHandler map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error

const (
	PingCommand          = "pingv2"
	SDVerifyCommand      = "sdverifyv2"
	SDQuestionOneCommand = "sdquestiononev2"
	CalendarCommand      = "calendarv2"
	NiceTryBro           = "Nice try, bro! You are not allowed to use this command... <@255514888041005057>"
)

func (u *usecase) registerSlashCommand() {
	log.Println("Adding commands...")
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        PingCommand,
			Description: "Ping",
		},
		//{
		//	Name:        SDVerifyCommand,
		//	Description: "Command for verifying new peeps and welcoming them",
		//	Options: []*discordgo.ApplicationCommandOption{
		//		{
		//			Type:        discordgo.ApplicationCommandOptionUser,
		//			Name:        "user-option",
		//			Description: "User to verify",
		//			Required:    true,
		//		},
		//		{
		//			Type:        discordgo.ApplicationCommandOptionString,
		//			Name:        "role-option",
		//			Description: "Religion role to give",
		//			Required:    true,
		//			Choices:     buildReligionRoleOptionChoices(),
		//		},
		//	},
		//},
		//{
		//	Name:        SDQuestionOneCommand,
		//	Description: "Command for alerting peeps that they missed question one code",
		//	Options: []*discordgo.ApplicationCommandOption{
		//		{
		//			Type:        discordgo.ApplicationCommandOptionUser,
		//			Name:        "user-option",
		//			Description: "User to alert",
		//			Required:    true,
		//		},
		//	},
		//},
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
		PingCommand: u.pingCommandFunc,
		//SDVerifyCommand:      u.sdVerifyCommandFunc,
		//SDQuestionOneCommand: u.sdQuestionOneCommandFunc,
		CalendarCommand: u.calendarCommandFunc,
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

//func (u *usecase) sdVerifyCommandFunc(s *discordgo.Session, i *discordgo.InteractionCreate) error {
//	guildCfg, ok := u.getGuildConfig(i.GuildID)
//
//	if !ok {
//		return nil
//	}
//
//	if !u.isMod(i.Member.User.ID, guildCfg) {
//		err := u.alertNonMod(i)
//		if err != nil {
//			return err
//		}
//		return nil
//
//	}
//	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
//		Type: discordgo.InteractionResponseChannelMessageWithSource,
//		Data: &discordgo.InteractionResponseData{
//			Content: "Processing... please wait...",
//		},
//	})
//	if err != nil {
//		return err
//	}
//	// Access options in the order provided by the user.
//	options := i.ApplicationCommandData().Options
//	guildId := i.GuildID
//
//	// Or convert the slice into a map
//	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
//	for _, opt := range options {
//		optionMap[opt.Name] = opt
//	}
//
//	acknowledgementMessageArgs := make([]interface{}, 0, len(options))
//	acknowledgementMessageFormat := guildCfg.Wording.AcknowledgementMessageFormat
//	welcomeMessageArgs := make([]interface{}, 0, 1)
//
//	var user *discordgo.User
//
//	userOpt, userOptOk := optionMap["user-option"]
//	roleOpt, roleOptOk := optionMap["role-option"]
//	var roleType string
//	if userOptOk && roleOptOk {
//		user = userOpt.UserValue(s)
//		acknowledgementMessageArgs = append(acknowledgementMessageArgs, user.ID)
//		welcomeMessageArgs = append(welcomeMessageArgs, user.ID)
//		welcomeMessageArgs = append(welcomeMessageArgs, guildCfg.Channel.ReactionRoles)
//		welcomeMessageArgs = append(welcomeMessageArgs, guildCfg.Channel.ServerInformation)
//
//		//actually i dont need to put this in here, because user is required anyway. but just to be safe haha
//		roleType = roleOpt.StringValue()
//		roleId := guildCfg.ReligionRoleMappingMap[config.ReligionRoleType(roleType)]
//		acknowledgementMessageArgs = append(acknowledgementMessageArgs, roleId)
//
//		err := s.GuildMemberRoleAdd(guildId, user.ID, string(roleId))
//		if err != nil {
//			fmt.Println(err)
//			return err
//		}
//
//		err = s.GuildMemberRoleAdd(guildId, user.ID, guildCfg.Role.ApprovedUser)
//		if err != nil {
//			return err
//		}
//		err = s.GuildMemberRoleRemove(guildId, user.ID, guildCfg.Role.Vetting)
//		if err != nil {
//			return err
//		}
//		err = s.GuildMemberRoleRemove(guildId, user.ID, guildCfg.Role.VettingQuestioning)
//		if err != nil {
//			return err
//		}
//
//	} else {
//		_, err := s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
//			Content: "Please choose user and role.",
//		})
//		if err != nil {
//			return err
//		}
//		return nil
//	}
//
//	mod := i.Member
//	content := fmt.Sprintf(guildCfg.Wording.WelcomeMessageFormat, user.Mention(), mod.Mention())
//	_, err = s.ChannelMessageSend(guildCfg.Channel.GeneralDiscussion, content)
//	if err != nil {
//		return err
//	}
//
//	_, err = s.ChannelMessageSendEmbed(guildCfg.Channel.GeneralDiscussion, util.EmbedBuilder(guildCfg.Wording.WelcomeTitle, fmt.Sprintf(guildCfg.Wording.WelcomeMessageEmbedFormat, welcomeMessageArgs...), util.RandomWelcomeImage()))
//	if err != nil {
//		return err
//	}
//	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
//		Content: fmt.Sprintf(
//			acknowledgementMessageFormat,
//			acknowledgementMessageArgs...,
//		),
//	})
//
//	if err != nil {
//		return err
//	}
//
//	log.Printf("[%s] : [%s] | [%s] | [%s]", SDVerifyCommand, mod.User.Username, user.Username, roleType)
//	return nil
//}
//
//func (u *usecase) sdQuestionOneCommandFunc(s *discordgo.Session, i *discordgo.InteractionCreate) error {
//
//	guildConfig, ok := u.getGuildConfig(i.GuildID)
//	if !ok {
//		return nil
//	}
//	if !u.isMod(i.Member.User.ID, guildConfig) {
//		err := u.alertNonMod(i)
//		if err != nil {
//			return err
//		}
//		return nil
//	}
//
//	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
//		Type: discordgo.InteractionResponseChannelMessageWithSource,
//		Data: &discordgo.InteractionResponseData{
//			Content: "Processing... please wait...",
//		},
//	})
//	if err != nil {
//		return err
//	}
//	// Access options in the order provided by the user.
//	options := i.ApplicationCommandData().Options
//	guildId := i.GuildID
//
//	// Or convert the slice into a map
//	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
//	for _, opt := range options {
//		optionMap[opt.Name] = opt
//	}
//
//	missedQuestionOneMessageFormatArgs := make([]interface{}, 0)
//	missedQuestionOneMessageFormat := guildConfig.Wording.MissedQuestionOneFormatNoPS
//
//	var user *discordgo.User
//
//	userOpt, userOptOk := optionMap["user-option"]
//	if userOptOk {
//		user = userOpt.UserValue(s)
//		missedQuestionOneMessageFormatArgs = append(missedQuestionOneMessageFormatArgs, user.ID)
//		missedQuestionOneMessageFormatArgs = append(missedQuestionOneMessageFormatArgs, guildConfig.Channel.RulesVetting)
//		err := s.GuildMemberRoleAdd(guildId, user.ID, guildConfig.Role.VettingQuestioning)
//		if err != nil {
//			fmt.Println(err)
//			return err
//		}
//	} else {
//		_, err := s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
//			Content: "Please choose user.",
//		})
//		if err != nil {
//			return err
//		}
//		return nil
//	}
//
//	mod := i.Member
//
//	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
//		Content: fmt.Sprintf(
//			missedQuestionOneMessageFormat,
//			missedQuestionOneMessageFormatArgs...,
//		),
//	})
//
//	if err != nil {
//		return err
//	}
//
//	log.Printf("[%s] : [%s] | [%s]", SDQuestionOneCommand, mod.User.Username, user.Username)
//	return nil
//}

func (u *usecase) calendarCommandFunc(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Processing... please wait...",
		},
	})
	if err != nil {
		return err
	}

	embed, err := u.generateCalendarEmbed()
	if err != nil {
		u.errorReporter(err)
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
		u.errorReporter(err)
		return err
	}

	log.Printf("[%s] : [%s] ", CalendarCommand, i.Interaction.Member.User.Username)
	return nil
}
