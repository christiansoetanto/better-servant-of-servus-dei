package dbot

import (
	"encoding/json"
	"fmt"
	"github.com/christiansoetanto/better-servant-of-servus-dei/util"
	"github.com/robfig/cron/v3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type AllLiturgicalDays struct {
	LiturgicalDaysEn []LiturgicalDay
	LiturgicalDaysLa []LiturgicalDay
}
type LiturgicalDay struct {
	Key                   string        `json:"key"`
	Date                  string        `json:"date"`
	Precedence            string        `json:"precedence"`
	Rank                  string        `json:"rank"`
	IsHolyDayOfObligation bool          `json:"isHolyDayOfObligation"`
	IsOptional            bool          `json:"isOptional"`
	Martyrology           []interface{} `json:"martyrology"`
	Titles                []string      `json:"titles"`
	Calendar              Calendar      `json:"calendar"`
	Cycles                Cycles        `json:"cycles"`
	Name                  string        `json:"name"`
	RankName              string        `json:"rankName"`
	ColorName             []string      `json:"colorName"`
	SeasonNames           []string      `json:"seasonNames"`
}
type Calendar struct {
	WeekOfSeason          int    `json:"weekOfSeason,omitempty"`
	DayOfSeason           int    `json:"dayOfSeason,omitempty"`
	DayOfWeek             int    `json:"dayOfWeek,omitempty"`
	NthDayOfWeekInMonth   int    `json:"nthDayOfWeekInMonth,omitempty"`
	StartOfSeason         string `json:"startOfSeason,omitempty"`
	EndOfSeason           string `json:"endOfSeason,omitempty"`
	StartOfLiturgicalYear string `json:"startOfLiturgicalYear,omitempty"`
	EndOfLiturgicalYear   string `json:"endOfLiturgicalYear,omitempty"`
}
type Cycles struct {
	ProperCycle  string `json:"properCycle"`
	SundayCycle  string `json:"sundayCycle"`
	WeekdayCycle string `json:"weekdayCycle"`
	PsalterWeek  string `json:"psalterWeek"`
}
type Martyrology struct {
	Key               string   `json:"key"`
	CanonizationLevel string   `json:"canonizationLevel"`
	DateOfDeath       int      `json:"dateOfDeath"`
	Titles            []string `json:"titles,omitempty"`
}

type Messages struct {
	Messages []MessageItem `json:"messages"`
}
type MessageItem struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

const DailyCron = "@daily"
const Every5SecondCron = "@every 5s"

func (b *Bot) initCronJob() {
	c := cron.New()
	_, err := c.AddFunc(DailyCron, b.liturgicalCalendarCronJob)
	if err != nil {
		return
	}
	c.Start()
}

func (b *Bot) liturgicalCalendarCronJob() {
	functionsUrl := os.Getenv("ROMCAL_API_FUNCTIONS_URL")
	response, err := http.Get(functionsUrl)
	if err != nil {
		b.errorReporter(err)
		return
	}
	data, _ := ioutil.ReadAll(response.Body)

	var allLiturgicalDays AllLiturgicalDays
	err = json.Unmarshal(data, &allLiturgicalDays)
	if err != nil {
		b.errorReporter(err)
		return
	}

	currentTime := time.Now()
	text := getCelebrations(allLiturgicalDays.LiturgicalDaysEn)
	title := fmt.Sprintf("%s, %d %s %d", currentTime.Weekday(), currentTime.Day(), currentTime.Month(), currentTime.Year())
	for _, config := range b.Cfg {
		_, err = b.Session.ChannelMessageSendEmbed(config.Channel.BotTesting, util.EmbedBuilder(title, text))
		if err != nil {
			b.errorReporter(err)
			return
		}
	}
}

func getCelebrations(liturgicalDays []LiturgicalDay) string {
	text := "The Roman Catholic Church is celebrating:\n\n"
	for _, day := range liturgicalDays {
		text += "• "
		//[day, date] //if memorial/feast/solemnity [rank] [name] in [seasonName] season.
		rank, rankName, isHolyDayOfObligation, name, seasonNames := strings.ToLower(day.Rank), day.RankName,
			day.IsHolyDayOfObligation, day.Name, day.SeasonNames
		if rank == "memorial" || rank == "feast" || rank == "solemnity" {
			text += fmt.Sprintf("%s of %s", cases.Title(language.AmericanEnglish).String(rankName), name)
			if len(seasonNames) > 0 {
				text += fmt.Sprintf(" in the %s", seasonNames[0])
			}
		} else {
			text += name
		}

		if isHolyDayOfObligation {
			text += ". A Holy Day of Obligation"
		}

		text += ".\n"
	}
	return text

}
