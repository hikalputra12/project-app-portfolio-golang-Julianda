package model

type Education struct {
	Institution string
	Degree      string
	Major       string
	Start_year  int
	End_year    int
	Description string
}

type Experience struct {
	Title       string
	Company     string
	Type        string
	Start_year  int
	End_year    int
	Description string
}

type Skill struct {
	Name     string
	Category string
	Level    string
	Icon     string
}
