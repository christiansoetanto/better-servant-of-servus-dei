package usecase

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/better-servant-of-servus-dei/config/configtypes"
	"github.com/christiansoetanto/better-servant-of-servus-dei/util"
	"strings"
)

const (
	LimitPerRequest  = 100
	MaxMessageAmount = 3000
)

func (u *usecase) initReactionAddHandler() {
	u.Session.AddHandler(u.religiousQuestionReactionAddHandler)
}

type answersMap map[string][]answer
type answer struct {
	user *discordgo.User
	url  string
}

func (u *usecase) religiousQuestionReactionAddHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	if m.UserID == s.State.User.ID {
		return
	}
	cfg, ok := u.getGuildConfig(m.GuildID)
	if !ok {
		return
	}
	if m.ChannelID != cfg.Channel.ReligiousQuestions {
		return
	}
	if m.Emoji.ID != cfg.Reaction.Upvote {
		return
	}
	if !u.isMod(m.UserID, cfg) {
		return
	}

	questionId := m.MessageID
	message, err := s.ChannelMessage(cfg.Channel.ReligiousQuestions, questionId)
	if err != nil {
		u.errorReporter(err)
		return
	}
	question, questionAskerId := message.Content, message.Author.ID
	answers, err := u.getAllAnswers(m.ChannelID, m.MessageID, cfg)
	if err != nil {
		u.errorReporter(err)
		return
	}
	if answers == nil {
		return
	}
	answers, err = u.generateAnswerUrl(answers, question, cfg.Guild.GuildID)
	if err != nil {
		u.errorReporter(err)
		return
	}
	err = u.archiveQuestion(answers, questionAskerId, questionId, question, cfg)
	if err != nil {
		u.errorReporter(err)
		return
	}
}

func (u *usecase) getAllAnswers(channelId, messageId string, cfg configtypes.GuildConfig) (answersMap, error) {
	rd1, err := u.getAnswersForChannel(channelId, messageId, cfg.Reaction.ReligiousDiscussionOneWhiteCheckmark)
	if err != nil {
		u.errorReporter(err)
		return nil, err
	}
	rd2, err := u.getAnswersForChannel(channelId, messageId, cfg.Reaction.ReligiousDiscussionsTwoBallotBoxWithCheck)
	if err != nil {
		u.errorReporter(err)
		return nil, err
	}
	answersMap := make(answersMap)
	if rd1 != nil {
		answersMap[cfg.Channel.ReligiousDiscussions1] = rd1
	}
	if rd2 != nil {
		answersMap[cfg.Channel.ReligiousDiscussions2] = rd2
	}

	return answersMap, nil
}

func (u *usecase) getAnswersForChannel(channelId, messageId, reactionId string) ([]answer, error) {
	users, err := u.Session.MessageReactions(channelId, messageId, reactionId, 0, "", "")
	if err != nil {
		u.errorReporter(err)
		return nil, err
	}
	answers := answersBuilder(users)
	return answers, err
}

func answersBuilder(users []*discordgo.User) []answer {
	var answers []answer
	for _, user := range users {
		answers = append(answers, answer{
			user: user,
		})
	}
	return answers
}

func (u *usecase) generateAnswerUrl(answerMap answersMap, question, guildId string) (answersMap, error) {
	for channelId := range answerMap {
		answers := answerMap[channelId]
		lastMessageId := ""
		totalAnswerToBeFound := len(answers)
		for i := 0; i < MaxMessageAmount/LimitPerRequest && totalAnswerToBeFound > 0; i++ {
			fmt.Printf("current iter: %d, max iter: %d, answer left: %d\n", i, MaxMessageAmount/LimitPerRequest, totalAnswerToBeFound)
			messages, err := u.Session.ChannelMessages(channelId, LimitPerRequest, lastMessageId, "", "")
			if err != nil {
				u.errorReporter(err)
				return nil, err
			}
			if len(messages) == 0 {
				return answerMap, nil
			}
			lastMessageId = messages[len(messages)-1].ID
			//loop setiap message, cari apakah message itu dimiliki oleh user yang menjawab
			for _, message := range messages {
				for i := range answers {
					if message.Author.ID == answers[i].user.ID {
						//TODO ganti ke levenshtein
						if strings.Contains(util.ToOnlyAlphanum(message.Content), util.ToOnlyAlphanum(question)) {
							answerLink := fmt.Sprintf("https://discord.com/channels/%s/%s/%s", guildId, channelId, message.ID)
							answers[i].url = answerLink
							totalAnswerToBeFound -= 1
						}
					}
				}
				if totalAnswerToBeFound <= 0 {
					break
				}
			}
		}
	}
	return answerMap, nil
}

func (u *usecase) archiveQuestion(answers answersMap, questionAskerId, questionId, questionContent string, cfg configtypes.GuildConfig) error {
	title := "Religious Question Police!"
	description := fmt.Sprintf("\nQuestion by <@%s>:\n\n%s\n", questionAskerId, questionContent)
	fieldValue := ""
	fieldName := "Answer(s):"
	for _, answers := range answers {
		for _, answer := range answers {
			fieldValue += fmt.Sprintf("\n- <@%s>", answer.user.ID)
			if len(answer.url) > 0 {
				fieldValue += fmt.Sprintf(" [jump to answer!](%s)", answer.url)
			}
		}
	}

	var fields []*discordgo.MessageEmbedField
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   fieldName,
		Value:  fieldValue,
		Inline: false,
	})
	embed := util.EmbedBuilder(title, description, fields)

	_, err := u.Session.ChannelMessageSendEmbed(cfg.Channel.AnsweredQuestions, embed)

	if err != nil {
		u.errorReporter(err)
		return err
	}

	if !u.Config.AppConfig.DevMode {
		err = u.Session.ChannelMessageDelete(cfg.Channel.ReligiousQuestions, questionId)
		if err != nil {
			u.errorReporter(err)
			return err
		}
	}

	return nil
}
