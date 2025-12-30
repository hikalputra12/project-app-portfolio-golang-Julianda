package repository

import "project-portofolio/database"

type Repository struct {
	Biodata    BiodataRepository
	About      AboutRepository
	Resume     ResumeRepository
	Portofolio PortofolioRepository
	Message    MessageRepository
	Skill      SkillRepository
}

func NewRepository(db database.PgxIface) Repository {
	return Repository{
		Biodata:    NewBiodataRepository(db),
		About:      NewAboutRepository(db),
		Resume:     NewResumeRepository(db),
		Portofolio: NewPortofolioRepository(db),
		Message:    NewMessageRepository(db),
		Skill:      NewSkillRepository(db),
	}
}
