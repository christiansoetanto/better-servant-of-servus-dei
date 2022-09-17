package configtypes

type Role struct {
	Vetting            string
	VettingQuestioning string
	ApprovedUser       string
	LatinCatholic      string
	EasternCatholic    string
	OrthodoxChristian  string
	RCIACatechumen     string
	Protestant         string
	NonCatholic        string
	Atheist            string
	Moderator          string
}

const (
	LatinCatholic     ReligionRoleType = "Latin Catholic"
	EasternCatholic   ReligionRoleType = "Eastern Catholic"
	OrthodoxChristian ReligionRoleType = "Orthodox Christian"
	RCIACatechumen    ReligionRoleType = "RCIA / Catechumen"
	Protestant        ReligionRoleType = "Protestant"
	NonCatholic       ReligionRoleType = "Non Catholic"
	Atheist           ReligionRoleType = "Atheist"
)

type ReligionRoleType string
type ReligionRoleId string
type ReligionRoleMappingMap map[ReligionRoleType]ReligionRoleId
