package model

type About struct {
	IntroAbout string
	FocusAbout string
}

type AboutPageData struct {
	IntroData *About
	SkillData []Skill
}
