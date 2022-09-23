package usecase

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/better-servant-of-servus-dei/util"
	"log"
	"regexp"
	"strings"
)

const INRI = "inri"
const ANDGIVEUSTHECODE = "andgiveusthecode"
const WHATCODE = "whatcode"

func (u *usecase) initMessageHandler() {
	u.Session.AddHandler(u.invalidVettingResponseHandler)
	u.Session.AddHandler(u.vettingQuestioningResponseHandler)
}

func detectVettingResponse(input string) bool {
	reg, err := regexp.Compile(".*1.*2.*3.*4.*5.*6.*")
	if err != nil {
		log.Println(err)
		return false
	}
	return strings.Contains(input, ANDGIVEUSTHECODE) || strings.Contains(input, WHATCODE) || reg.MatchString(input)
}

func isValidVettingResponse(input string) bool {
	input = strings.ReplaceAll(util.ToOnlyAlphanum(input), "latinrite", "")
	if detectVettingResponse(input) && !strings.Contains(input, INRI) {
		return false
	}
	return true
}

func (u *usecase) invalidVettingResponseHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	cfg, ok := u.getGuildConfig(m.GuildID)
	if !ok {
		return
	}
	if m.ChannelID != cfg.Channel.Responses {
		return
	}

	fmt.Println(m.Content, m.GuildID, m.ChannelID)
	if !isValidVettingResponse(m.Content) {
		content := fmt.Sprintf("Hey <@%s>! It looks like you missed question 1. Please re-read the <#%s> again, we assure you that the code is in there. Thank you for your understanding.", m.Author.ID, cfg.Channel.RulesVetting)
		_, err := s.ChannelMessageSendEmbedReply(cfg.Channel.Responses, util.EmbedBuilder(fmt.Sprintf("%s Vetting %s", cfg.Reaction.Dab, cfg.Reaction.Dab), content), m.Reference())
		if err != nil {
			u.errorReporter(err)
			return
		}
	}
}
func (u *usecase) vettingQuestioningResponseHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	cfg, ok := u.getGuildConfig(m.GuildID)
	if !ok {
		return
	}
	if m.ChannelID != cfg.Channel.VettingQuestioning {
		return
	}
	if strings.Contains(util.OnlyAlphanumAndSpace(m.Content), INRI) {
		title := fmt.Sprintf("%s Vetting %s", cfg.Reaction.Dab, cfg.Reaction.Dab)
		description := fmt.Sprintf("\nNice job, <@%s>! <@&%s> give this man a cookie.\n\n\nPS: this is a joke. please wait for our human (or are they) mods to verify you.", m.Author.ID, cfg.Role.Moderator)
		embed := util.EmbedBuilder(title, description)
		content := fmt.Sprintf("Come here you <@&%s>. Look at this dude <@%s>", cfg.Role.Moderator, m.Author.ID)
		_, err := s.ChannelMessageSendComplex(cfg.Channel.VettingQuestioning, &discordgo.MessageSend{
			Content:   content,
			Embed:     embed,
			Reference: m.Reference(),
		})
		if err != nil {
			u.errorReporter(err)
			return
		}
		err = s.ChannelMessageDelete(m.ChannelID, m.ID)
		if err != nil {
			u.errorReporter(err)
			return
		}
	}
}
