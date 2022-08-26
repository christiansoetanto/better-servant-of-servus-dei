package configtypes

type Role struct {
	Vetting            RoleId
	VettingQuestioning RoleId
	ApprovedUser       RoleId
	LatinCatholic      RoleId
	EasternCatholic    RoleId
	OrthodoxChristian  RoleId
	RCIACatechumen     RoleId
	Protestant         RoleId
	NonCatholic        RoleId
	Atheist            RoleId
}
type RoleId string
